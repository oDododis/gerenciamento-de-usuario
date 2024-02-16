package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Recebe o paramertro do id e manda Deleta-lo

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	autenticationToken := c.Request.Header.Get("Authorization")
	token := strings.Split(autenticationToken, " ")
	if err := uc.serviceToken.TokenAutentication(token[1]); err != nil {

		c.JSON(err.Code, err)
		return
	}
	userID := c.Param("userID")
	if err := uc.service.DeleteUserServices(userID); err != nil {

		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusAccepted, "User excluido")
}
