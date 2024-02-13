package service

import (
	"Teste/src/configuration/rest_error"
	"Teste/src/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strconv"
)

//Procura o usuario baseado no id recebido pelo controller

func (ud *userDomainService) FindUserIDServices(userID string) (model.UserDomainInterface, *rest_error.RestError) {

	db, err := gorm.Open(sqlite.Open("usersFromBreadOfPotato.db"), &gorm.Config{})
	if err != nil {
		return nil, rest_error.NewInternalServerError("Não iniciou o Banco de Dados em service/findUser")
	}

	if err = db.AutoMigrate(&userDomainService{}); err != nil {
		return nil, rest_error.NewInternalServerError("Não iniciou o Banco de Dados em service/findUser")
	}
	userid, _ := strconv.Atoi(userID)
	if userid <= 0 {
		return nil, rest_error.NewBadRequestError("ID invalido (menor ou iqual a 0).")
	}
	var lastUd userDomainService
	db.First(&lastUd, userID)

	return &lastUd, nil
}

//Procura o usuario baseado no email recebido pelo controller

func (ud *userDomainService) FindUserEmailServices(userEmail string) (model.UserDomainInterface, *rest_error.RestError) {
	db, err := gorm.Open(sqlite.Open("usersFromBreadOfPotato.db"), &gorm.Config{})
	if err != nil {
		return nil, rest_error.NewInternalServerError("Não iniciou o Banco de Dados em service/findUser")
	}

	if err = db.AutoMigrate(&userDomainService{}); err != nil {
		return nil, rest_error.NewInternalServerError("Não iniciou o Banco de Dados em service/findUser")
	}

	if userEmail == "" {
		return nil, rest_error.NewBadRequestError("Email vazil")
	}
	var lastUd userDomainService
	db.First(&lastUd, "email = ?", userEmail)

	return &lastUd, nil
}
