package service

import (
	"Teste/src/configuration/rest_error"
	"Teste/src/model"
	"gorm.io/gorm"
)

//Cria um UserDomain privado apartir do Publico para manter os dado apenas com este dominio e o banco

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

// Não sei pra que serve mas se tirar da erro.

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

//Tipos de serviço para utilizarmos

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_error.RestError
	UpdateUser(string, model.UserDomainInterface) *rest_error.RestError
	FindUserIDServices(string) (model.UserDomainInterface, *rest_error.RestError)
	FindUserEmailServices(string) (model.UserDomainInterface, *rest_error.RestError)
	DeleteUser(string) *rest_error.RestError
}
