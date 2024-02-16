package controller

import (
	"Teste/src/configuration/validation"
	"Teste/src/controller/model/request"
	"Teste/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Recebe um Login Request e retorna as informações do usuario

func (uc *userControllerInterface) Login(c *gin.Context) {
	var userRequest request.LoginRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restError := validation.ValidateUserError(err)
		c.JSON(restError.Code, restError)
		return
	}

	domain := model.NewUserDomainLogin(
		userRequest.Email,
		userRequest.Password,
	)
	userDomain, err := uc.service.LoginServices(domain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusAccepted, "Token: "+userDomain)
}
