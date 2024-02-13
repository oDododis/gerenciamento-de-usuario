package service

import (
	"Teste/src/configuration/rest_error"
	"Teste/src/model"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//Recebe os campos do Controller e cria o usuario no Banco de Dados

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_error.RestError {
	userDomain.EncryptPassword()

	db, err := gorm.Open(sqlite.Open("usersFromBreadOfPotato.db"), &gorm.Config{})
	if err != nil {
		return rest_error.NewInternalServerError("Não iniciou o Banco de Dados em service/createUser")
	}
	err = db.AutoMigrate(&userDomainService{})
	if err != nil {
		return nil
	}

	ud = &userDomainService{
		FullName: userDomain.GetFullName(),
		Email:    userDomain.GetEmail(),
		Username: userDomain.GetUsername(),
		Password: userDomain.GetPassword(),
	}

	err = db.First(&ud, "email = ?", ud.Email).Error // db.First(&user, "username = ?", user.Username).Error
	err = db.First(&ud, "username = ?", ud.Username).Error
	if err != nil {
		fmt.Println("PASSOU POR AQUI.")
		db.Create(&ud)
	} else {
		return rest_error.NewBadRequestError("Email ou Username existente.")
	}
	fmt.Println(userDomain.GetPassword())
	return nil
}
