package controller

// Controlador dos Dominio

import (
	"Teste/src/model/service"
	"github.com/gin-gonic/gin"
)

// Cria um Controle de Interface

func NewUserControllerInterface(serviceInterface service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	FindUserID(c *gin.Context)
	FindUserEmail(c *gin.Context)

	DeleteUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
