package service

import (
	"auth-service/auth"
	"auth-service/model"
	"auth-service/repository"
	"errors"
	"fmt"
	"time"
)

type AccountService struct {
	database *repository.Database
}

func NewAccountService(db *repository.Database) *AccountService {
	return &AccountService{
		database: db,
	}
}

func (us *AccountService) Login(user *model.User) (model.UserInfo, error) {
	user.Pass = auth.HashString(user.Pass)
	uByName := model.User{
		UserName: user.UserName,
	}
	err := us.database.GetUserByName(&uByName)
	if err != nil {
		return model.UserInfo{}, err
	}
	if uByName.Try == 0 {
		t, err := time.Parse(time.RFC3339, uByName.UnlockTime)
		if err != nil {
			return model.UserInfo{}, err
		}
		if t.Before(time.Now()) {
			uByName.Try = 5
		} else {
			return model.UserInfo{}, fmt.Errorf("You have been lock for 30 minutes")
		}
	}
	if uByName.Pass != user.Pass {
		uByName.Try--
		uByName.UnlockTime = time.Now().Add(30 * time.Minute).Format(time.RFC3339)
		err := us.database.UpdateUser(&uByName)
		if err != nil {
			return model.UserInfo{}, err
		}
		return model.UserInfo{}, fmt.Errorf("Invalid Password; try: %d", uByName.Try)
	}
	err = us.database.GetUserByUserNameAndPass(user)
	if err != nil {
		return model.UserInfo{}, err
	}
	info := model.UserInfo{
		UserId: user.UserId,
	}
	err = us.database.GetUserInfo(&info)
	if err != nil {
		return model.UserInfo{}, err
	}
	user.Try = 5
	go us.database.UpdateUser(user)
	return info, nil
}

func (us *AccountService) CreateAccount(user *model.User) (model.UserInfo, error) {
	err := us.database.GetUserByName(user)
	if err != nil {
		user.UserId = auth.GenerateUUID()
		user.Pass = auth.HashString(user.Pass)
		user.Try = 5
		err = us.database.CreateUser(user)
		if err != nil {
			return model.UserInfo{}, err
		}
		info := model.UserInfo{
			UserId:   user.UserId,
			UserName: user.UserName,
			Point:    "0",
		}
		err = us.database.CreateUserInfo(&info)
		if err != nil {
			return model.UserInfo{}, err
		}
		return info, nil
	}
	return model.UserInfo{}, errors.New("username already exists")
}
