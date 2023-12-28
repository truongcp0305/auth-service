package incoming

import "auth-service/model"

type LoginParam struct {
	UserName string `json:"userName" form:"userName"`
	Password string `json:"password" form:"password"`
}

func (ic *LoginParam) GetModel() *model.User {
	return &model.User{
		UserName: ic.UserName,
		Pass:     ic.Password,
	}
}

type CreateUserParam struct {
	UserName string `json:"userName" form:"userName"`
	Password string `json:"password" form:"password"`
}

func (ic *CreateUserParam) GetModel() *model.User {
	return &model.User{
		UserName: ic.UserName,
		Pass:     ic.Password,
	}
}
