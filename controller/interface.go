package controller

import "github.com/gin-gonic/gin"

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	FindUserID(c *gin.Context)
	FindUserEmail(c *gin.Context)
	Login(c *gin.Context)
	UsersList(c *gin.Context)
	UpdateUser(c *gin.Context)
}
