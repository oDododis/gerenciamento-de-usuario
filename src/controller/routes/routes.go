package routes

import (
	"Teste/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("getUserID/:userID", controller.FindUserID)
	r.GET("getUserEmail/:userEmail", controller.FindUserEmail)
	r.POST("createUser/", controller.CreateUser)
	r.PUT("updateInfoUser/:userID", controller.UpdateUser)
	r.DELETE("deleteUser/:userID", controller.DeleteUser)
}
