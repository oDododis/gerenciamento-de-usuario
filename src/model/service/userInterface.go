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

type tokenDomainService struct {
	Token      string
	UserDomain userDomainService `gorm:"foreignKey:UserID"`
	UserID     uint
}

// Funcões pra coletar as informações individualmente para o view

func (ud *userDomainService) GetFullName() string {
	return ud.FullName
}

func (ud *userDomainService) GetEmail() string {
	return ud.Email
}

func (ud *userDomainService) GetUsername() string {
	return ud.Username
}

func (ud *userDomainService) GetPassword() string {
	return ud.Password
}

func (ud *userDomainService) EncryptPassword() {

}

//Tipos de serviço para utilizarmos

type UserDomainService interface {
	CreateUserServices(model.UserDomainInterface) *rest_error.RestError
	DeleteUserServices(string) *rest_error.RestError
	FindUserIDServices(string) (model.UserDomainInterface, *rest_error.RestError)
	FindUserEmailServices(string) (model.UserDomainInterface, *rest_error.RestError)
	HowMuchUsers() (int, *rest_error.RestError)
	LoginServices(model.UserDomainInterface) (model.UserDomainInterface, *rest_error.RestError)
	UpdateUserServices(string, model.UserDomainInterface) *rest_error.RestError
}
