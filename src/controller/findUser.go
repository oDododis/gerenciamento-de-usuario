package controller

import (
	"Teste/src/view"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//Fas a busca por ID

func (uc *userControllerInterface) FindUserID(c *gin.Context) {

	userID := c.Param("userID")
	userDomain, err := uc.service.FindUserIDServices(userID)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusAccepted, view.ConvertDomainToResponse(userDomain))
}

//Fas a busca por email

func (uc *userControllerInterface) FindUserEmail(c *gin.Context) {
	userEmail := c.Param("userEmail")
	userDomain, err := uc.service.FindUserEmailServices(userEmail)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusAccepted, view.ConvertDomainToResponse(userDomain))
}

// Fas uma listagem dos usuarios

func (uc *userControllerInterface) UsersList(c *gin.Context) {

	userID, err := uc.service.HowMuchUsers()
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	for i := 1; i <= userID; i++ {
		userDomain, err := uc.service.FindUserIDServices(strconv.Itoa(i))
		if err != nil {
			c.JSON(err.Code, err)
			return
		}
		c.JSON(http.StatusAccepted, "=-=-=-=-=-=-=-=-=-=-=-=-=-=")
		c.JSON(http.StatusAccepted, view.ConvertDomainToResponse(userDomain))
	}
	c.JSON(http.StatusAccepted, "=-=-=-=-=-=-=-=-=-=-=-=-=-=")
}
