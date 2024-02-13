package service

import (
	"Teste/src/configuration/rest_error"
	"Teste/src/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (ud *userDomainService) FindUserIDServices(userID string) (model.UserDomainInterface, *rest_error.RestError) {

	db, err := gorm.Open(sqlite.Open("usersFromBreadOfPotato.db"), &gorm.Config{})
	if err != nil {
		return nil, rest_error.NewInternalServerError("Não iniciou o Banco de Dados em service/findUser")
	}

	if err = db.AutoMigrate(&userDomainService{}); err != nil {
		return nil, rest_error.NewInternalServerError("Não iniciou o Banco de Dados em service/findUser")
	}
	db.First(&ud, userID)

	return ud, nil
}

func (ud *userDomainService) FindUserEmailServices(userEmail string) (model.UserDomainInterface, *rest_error.RestError) {
	db, err := gorm.Open(sqlite.Open("usersFromBreadOfPotato.db"), &gorm.Config{})
	if err != nil {
		return nil, rest_error.NewInternalServerError("Não iniciou o Banco de Dados em service/findUser")
	}

	if err = db.AutoMigrate(&userDomainService{}); err != nil {
		return nil, rest_error.NewInternalServerError("Não iniciou o Banco de Dados em service/findUser")
	}
	db.First(&ud, userEmail)

	return ud, nil
}
