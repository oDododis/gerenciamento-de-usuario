package main

import (
	"Teste/src/controller"
	"Teste/src/controller/routes"
	"Teste/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
)

type Users struct {
	gorm.Model
	FullName string
	Email    string
	Username string
	Password string
}

func main() {

	//Iniciando o godotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Inicialização das dependencias
	services := service.NewUserDomainServece()
	userController := controller.NewUserControllerInterface(services)

	//Iniciando as Rotas
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

//// Excluir da lista por completo
//func excludeUser(ID int) {
//	db, err := gorm.Open(sqlite.Open("usersFromBreadOfPotato.db"), &gorm.Config{})
//	if err != nil {
//		panic("failed to connect database")
//	}
//	err = db.AutoMigrate(&Users{})
//	if err != nil {
//		return
//	}
//	TABLE := "user_domain_services"
//	log.Println(ID)
//	db.Exec("DELETE FROM ? WHERE ID = ?", TABLE, ID)
//}
