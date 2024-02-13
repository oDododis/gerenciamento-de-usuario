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
	FullName string
	Email    string
	Username string
	Password string
	//birthday time.Time
}

func (ud *userDomainService) GetFullName() string {
	//TODO implement me
	panic("implement me")
}

func (ud *userDomainService) GetEmail() string {
	//TODO implement me
	panic("implement me")
}

func (ud *userDomainService) GetUsername() string {
	//TODO implement me
	panic("implement me")
}

func (ud *userDomainService) GetPassword() string {
	//TODO implement me
	panic("implement me")
}

func (ud *userDomainService) EncryptPassword() {
	//TODO implement me
	panic("implement me")
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_error.RestError
	UpdateUser(string, model.UserDomainInterface) *rest_error.RestError
	FindUserIDServices(string) (model.UserDomainInterface, *rest_error.RestError)
	FindUserEmailServices(string) (model.UserDomainInterface, *rest_error.RestError)
	DeleteUser(string) *rest_error.RestError
}
