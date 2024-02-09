package routes

import (
	"Teste/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("getUserID/:userID", userController.FindUserID)
	r.GET("getUserEmail/:userEmail", userController.FindUserEmail)
	r.POST("createUser/", userController.CreateUser)
	r.PUT("updateInfoUser/:userID", userController.UpdateUser)
	r.DELETE("deleteUser/:userID", userController.DeleteUser)
}
