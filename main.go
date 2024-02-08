package main

import (
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
	user1 := Users{
		FullName: "Douuglas Barbosa",
		Email:    "douglas@barbosa.com",
		Username: "dododis",
		Password: "12345",
	}
	user2 := Users{
		FullName: "Douuglas Monteiro",
		Email:    "douglas@moteiro.com",
		Username: "dodosubi",
		Password: "54321",
	}

	user3 := Users{
		FullName: "Pão de Batata",
		Email:    "pao@batata.com",
		Username: "paodebatata",
		Password: "p1203798!@$#)(*",
	}
	user4 := Users{
		FullName: "APSOJDPOASD",
		Email:    "pao@batata.com",
		Username: "APKSDJLKASJD",
		Password: "ASD{PLA`PSDKÀS$#)(*",
	}
	user5 := Users{}
	user6 := Users{Email: "ASPIDJ"}
	user7 := Users{Email: "asd"}
	createUser(user1)
	createUser(user2)
	createUser(user3)
	createUser(user4)
	createUser(user5)
	createUser(user6)
	createUser(user7)

	updateInfo(4, user4)

	deleteUser(5)
	excludeUser(6)
}

func createUser(user Users) {
	db, err := gorm.Open(sqlite.Open("usersFromBreadOfPotato.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Users{})

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
	db.AutoMigrate(&Users{})
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
	db.AutoMigrate(&Users{})

	var user Users
	db.Delete(&user, ID)
}

// Excluir da lista por completo
func excludeUser(ID int) {
	db, err := gorm.Open(sqlite.Open("usersFromBreadOfPotato.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Users{})
	log.Println(ID)
	db.Exec("delete from users where ID = ?", ID)
}
