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
	servicesToken := service.NewTokenDomainService()
	userController := controller.NewUserControllerInterface(services, servicesToken)

	//Iniciando as Rotas
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
