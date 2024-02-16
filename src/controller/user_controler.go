package controller

// Controlador dos Dominio

import (
	"Teste/src/model/service"
	"github.com/gin-gonic/gin"
)

// Cria os Controles para utilizarmos na criação, atualização, procura e exclusão de usuario

func NewUserControllerInterface(serviceInterface service.UserDomainService, serviceInterfaceToken service.TokenDomainService) UserControllerInterface {
	return &userControllerInterface{
		service:      serviceInterface,
		serviceToken: serviceInterfaceToken,
	}
}

//lista os comando criados acima

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	FindUserID(c *gin.Context)
	FindUserEmail(c *gin.Context)
	UsersList(c *gin.Context)
	Login(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type userControllerInterface struct {
	service      service.UserDomainService
	serviceToken service.TokenDomainService
}
