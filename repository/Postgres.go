package repository

import (
	"auth-service/model"

	"github.com/go-pg/pg/v10"
)

type Database struct {
	db *pg.DB
}

func NewDatabase(db *pg.DB) *Database {
	return &Database{
		db: db,
	}
}

func (r *Database) GetUserByName(user *model.User) error {
	err := r.db.Model(user).Where("user_name = ?", user.UserName).First()
	return err
}

func (r *Database) GetUserByUserNameAndPass(user *model.User) error {
	err := r.db.Model(user).Where("user_name = ?", user.UserName).Where("password = ?", user.Pass).First()
	return err
}

func (r *Database) UpdateUser(user *model.User) error {
	_, err := r.db.Model(user).Where("id = ?", user.UserId).Update()
	return err
}

func (r *Database) GetUserInfo(info *model.UserInfo) error {
	err := r.db.Model(info).Where("user_id = ?", info.UserId).First()
	return err
}

func (r *Database) CreateUser(user *model.User) error {
	_, err := r.db.Model(user).Insert()
	return err
}

func (r *Database) CreateUserInfo(info *model.UserInfo) error {
	_, err := r.db.Model(info).Insert()
	return err
}
