package main

import (
	"Teste/src/controller/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
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
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func createUser(user Users) {
	db, err := gorm.Open(sqlite.Open("usersFromBreadOfPotato.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Users{})
	if err != nil {
		return
	}

	err = db.First(&user, "email = ?", user.Email).Error // db.First(&user, "username = ?", user.Username).Error
	if err != nil {
		log.Println("Não tem Email.")
		db.Create(&user)
	} else {
		log.Println("Já tem Email.")
	}
}

func updateInfo(ID int, user Users) {
	db, err := gorm.Open(sqlite.Open("usersFromBreadOfPotato.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&Users{})
	if err != nil {
		return
	}
	var lastUser Users
	db.First(&lastUser, ID)
	db.Model(&lastUser).Updates(&user)
	log.Println("Update concluido.")
}

func deleteUser(ID int) {
	db, err := gorm.Open(sqlite.Open("usersFromBreadOfPotato.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&Users{})
	if err != nil {
		return
	}

	var user Users
	db.Delete(&user, ID)
}

// Excluir da lista por completo
func excludeUser(ID int) {
	db, err := gorm.Open(sqlite.Open("usersFromBreadOfPotato.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&Users{})
	if err != nil {
		return
	}
	log.Println(ID)
	db.Exec("DELETE FROM users WHERE ID = ?", ID)
}
