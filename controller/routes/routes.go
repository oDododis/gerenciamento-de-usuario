package routes

import (
	"Teste/controller"
	"github.com/gin-gonic/gin"
)

// Iniciando as Rotas

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("getUserID/:userID", userController.FindUserID)
	r.GET("getUserEmail/:userEmail", userController.FindUserEmail)
	r.GET("getUserList/", userController.UsersList)
	r.POST("login/", userController.Login)
	r.POST("createUser/", userController.CreateUser)
	r.PUT("updateUser/:userID", userController.UpdateUser)
	r.DELETE("deleteUser/:userID", userController.DeleteUser)
}
