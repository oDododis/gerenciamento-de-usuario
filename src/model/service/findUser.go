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
		return nil, rest_error.NewNotFoundError("ID invalido (menor ou iqual a 0).")
	}
	ud = &userDomainService{}

	err = db.First(&ud, userID).Error
	if err != nil {
		return nil, rest_error.NewNotFoundError("ID não encontrado.")
	}
	return ud, nil
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
	ud = &userDomainService{
		Email: userEmail,
	}
	if userEmail == "" {
		return nil, rest_error.NewNotFoundError("Email vazil")
	} else {
		err = db.First(&ud, "email = ?", userEmail).Error
		if err != nil {
			return nil, rest_error.NewNotFoundError("Email não encontrado.")
		}
		return ud, nil
	}
}

func (ud *userDomainService) HowMuchUsers() (int, *rest_error.RestError) {
	db, err := gorm.Open(sqlite.Open("usersFromBreadOfPotato.db"), &gorm.Config{})
	if err != nil {
		return 0, rest_error.NewInternalServerError("Não iniciou o Banco de Dados em service/findUser")
	}

	if err = db.AutoMigrate(&userDomainService{}); err != nil {
		return 0, rest_error.NewInternalServerError("Não iniciou o Banco de Dados em service/findUser")
	}
	err = db.Last(ud).Error
	if err != nil {
		return 0, rest_error.NewNotFoundError("Não tem Usuarios")
	}
	return int(ud.ID), nil
}

func (ud *userDomainService) ListUserIDServices(userID string) (model.UserDomainInterface, *rest_error.RestError) {

	db, err := gorm.Open(sqlite.Open("usersFromBreadOfPotato.db"), &gorm.Config{})
	if err != nil {
		return nil, rest_error.NewInternalServerError("Não iniciou o Banco de Dados em service/findUser")
	}

	if err = db.AutoMigrate(&userDomainService{}); err != nil {
		return nil, rest_error.NewInternalServerError("Não iniciou o Banco de Dados em service/findUser")
	}
	ud = &userDomainService{}

	db.First(&ud, userID)
	if ud.ID == 0 {
		var empty userDomainService
		userid, _ := strconv.Atoi(userID)
		empty.ID = uint(userid)
		empty.FullName = "nonexistent"
		empty.Email = "nonexistent"
		empty.Username = "nonexistent"
		return &empty, rest_error.NewNotFoundError("Usuario não encontrado.")
	} else {
		return ud, nil
	}

}
