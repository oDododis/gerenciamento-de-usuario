package main

import (
	"Teste/src/controller"
	"Teste/src/controller/routes"
	"Teste/src/model/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	//Inicialização das dependencias de serviço e controle
	services := service.NewUserDomainServece()
	userController := controller.NewUserControllerInterface(services)
	tokenServices := service.NewTokenDomainServece()
	userControllerToken := controller.NewUserControllerToken(tokenServices)

	//Iniciando as Rotas
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController, userControllerToken)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
