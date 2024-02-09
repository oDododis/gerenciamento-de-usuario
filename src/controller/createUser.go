package controller

import (
	"Teste/src/configuration/validation"
	"Teste/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restError := validation.ValidateUserError(err)
		c.JSON(restError.Code, restError)
	}
}
