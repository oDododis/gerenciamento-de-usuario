package service

import (
	"Teste/src/configuration/rest_error"
	"Teste/src/model"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_error.RestError {
	userDomain.EncryptPassword()

	type Users struct {
		gorm.Model
		FullName string
		Email    string
		Username string
		Password string
	}

	db, err := gorm.Open(sqlite.Open("usersFromBreadOfPotato.db"), &gorm.Config{})
	if err != nil {
		rest_error.NewInternalServerError("Não abriu o gorm")
	}
	err = db.AutoMigrate(&userDomainService{})
	if err != nil {
		return nil
	}

	err = db.First(&ud, "email = ?", userDomain.GetEmail()).Error // db.First(&user, "username = ?", user.Username).Error
	if err != nil {
		log.Println("Não tem Email.")
		db.Create(&userDomain)
	} else {
		rest_error.NewBadRequestError("Não pode ser criado, email existente.")
	}
	fmt.Println(userDomain.GetPassword())
	return nil
}
