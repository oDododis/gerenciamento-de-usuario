package main

import (
	"Teste/controller"
	"Teste/controller/routes"
	"Teste/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	//Inicialização das dependencias de serviço e controle
	dataBase := service.NewDB()
	err := dataBase.StartConnection()
	if err != nil {
		log.Fatal(err)
	}
	userServices := service.NewUserService(dataBase.GetConnection())
	tokenServices := service.NewTokenService(dataBase.GetConnection())
	
	userController := controller.NewUserController(userServices, tokenServices)

	//Iniciando as Rotas
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
