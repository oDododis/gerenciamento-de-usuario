package controller

import (
	"Teste/src/configuration/rest_error"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_error.NewBadRequestError("Errou!")
	c.JSON(err.Code, err)
}
