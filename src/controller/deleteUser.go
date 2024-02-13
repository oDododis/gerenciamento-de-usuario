package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Recebe o paramertro do id e manda Deleta-lo

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {

	userID := c.Param("userID")
	if err := uc.service.DeleteUser(userID); err != nil {

		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, userID)
}
