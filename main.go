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
	services := service.NewUserService()
	servicesToken := service.NewTokenService()
	userController := controller.NewUserController(services, servicesToken)

	//Iniciando as Rotas
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
