package service

import (
	"20241104/class/2/model"
	"20241104/class/2/repository"
)

type UserService struct {
	User repository.User
}

func InitUserService(repo repository.User) *UserService {
	return &UserService{User: repo}
}

func (repo *UserService) Create(user model.User) *model.Response {
	session, err := repo.User.Create(&user)
	if err != nil {
		return &model.Response{StatusCode: 500, Message: "Server Error", Data: err.Error()}
	}
	return &model.Response{StatusCode: 200, Message: "User successfully registered and login session created", Data: session}
}
