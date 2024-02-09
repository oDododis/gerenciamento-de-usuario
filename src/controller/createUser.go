package controller

import (
	"Teste/src/configuration/validation"
	"Teste/src/controller/model/request"
	"Teste/src/model"
	"Teste/src/model/service"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restError := validation.ValidateUserError(err)
		c.JSON(restError.Code, restError)
	}

	domain := model.NewUserDomain(
		userRequest.FullName,
		userRequest.Email,
		userRequest.Username,
		userRequest.Password,
	)
	service := service.NewUserDomainServece()
	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

}
