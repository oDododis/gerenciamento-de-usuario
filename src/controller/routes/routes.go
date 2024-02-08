package routes

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.RouterGroup) {
	r.GET("getUserID/:userID")
	r.GET("getUserEmail/:userEmail")
	r.POST("createUser/")
	r.PUT("updateInfoUser/:userID")
	r.DELETE("deleteUser/:userID")
}
