package service

import (
	"Teste/src/configuration/rest_error"
	"Teste/src/model"
	"gorm.io/gorm"
)

func NewUserDomainServece() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
	gorm.Model
	fullName string
	email    string
	username string
	password string
	//birthday time.Time

}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_error.RestError
	UpdateUser(string, model.UserDomainInterface) *rest_error.RestError
	FindUser(string) (*model.UserDomainInterface, *rest_error.RestError)
	DeleteUser(string) *rest_error.RestError
}
