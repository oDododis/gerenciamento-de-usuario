package controller

import (
	"Teste/src/view"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

//Fas a busca por ID

func (uc *userControllerInterface) FindUserID(c *gin.Context) {
	autenticationToken := c.Request.Header.Get("Authorization")
	token := strings.Split(autenticationToken, " ")
	if err := uc.serviceToken.TokenAutentication(token[1]); err != nil {

		c.JSON(err.Code, err)
		return
	}
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
	autenticationToken := c.Request.Header.Get("Authorization")
	token := strings.Split(autenticationToken, " ")
	if err := uc.serviceToken.TokenAutentication(token[1]); err != nil {

		c.JSON(err.Code, err)
		return
	}
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
	autenticationToken := c.Request.Header.Get("Authorization")
	token := strings.Split(autenticationToken, " ")
	if err := uc.serviceToken.TokenAutentication(token[1]); err != nil {

		c.JSON(err.Code, err)
		return
	}
	userID, err := uc.service.HowMuchUsers()
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	for i := 1; i <= userID; i++ {
		userDomain, err := uc.service.ListUserIDServices(strconv.Itoa(i))
		c.JSON(http.StatusAccepted, "=-=-=-=-=-=-=-=-=-=-=-=-=-=")
		c.JSON(http.StatusAccepted, view.ConvertDomainToResponse(userDomain))
		if err != nil {
			c.JSON(err.Code, err)
		}
	}
	c.JSON(http.StatusAccepted, "=-=-=-=-=-=-=-=-=-=-=-=-=-=")
}
