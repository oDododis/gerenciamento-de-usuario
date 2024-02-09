package service

import (
	"Teste/src/configuration/rest_error"
	"Teste/src/model"
)

func NewUserDomainServece() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_error.RestError
	UpdateUser(string, model.UserDomainInterface) *rest_error.RestError
	FindUser(string) (*model.UserDomainInterface, *rest_error.RestError)
	DeleteUser(string) *rest_error.RestError
}
