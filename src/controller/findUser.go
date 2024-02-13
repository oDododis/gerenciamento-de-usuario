package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//Fas a busca por ID

func (uc *userControllerInterface) FindUserID(c *gin.Context) {

	userID := c.Param("userID")
	userDomain, err := uc.service.FindUserIDServices(userID)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, userDomain)
}

//Fas a busca por email

func (uc *userControllerInterface) FindUserEmail(c *gin.Context) {
	userEmail := c.Param("userEmail")
	userDomain, err := uc.service.FindUserEmailServices(userEmail)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, userDomain)
}
