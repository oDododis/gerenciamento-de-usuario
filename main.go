package main

import (
	"Teste/controller"
	"Teste/controller/routes"
	"Teste/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	//Inicialização das dependencias de serviço, controle e banco de dados
	dataBase := service.NewDB()
	if err := dataBase.StartConnection(); err != nil {
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
