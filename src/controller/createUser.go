package controller

import (
	"Teste/src/configuration/validation"
	"Teste/src/controller/model/request"
	"Teste/src/model"
	"Teste/src/view"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Recebe o Request e manda Criar o usuario

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restError := validation.ValidateUserError(err)
		c.JSON(restError.Code, restError)
		return
	}

	domain := model.NewUserDomain(
		userRequest.FullName,
		userRequest.Email,
		userRequest.Username,
		userRequest.Password,
	)

	if err := uc.service.CreateUserServices(domain); err != nil {

		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusCreated, view.ConvertDomainToResponse(domain))
}
