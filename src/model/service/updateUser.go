package service

import (
	"Teste/src/configuration/rest_error"
	"Teste/src/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strconv"
)

//Recebe os campos do controller, o id do usuario e atualiza os campos modificados no envio

func (ud *userDomainService) UpdateUserServices(userID string, userDomain model.UserDomainInterface) *rest_error.RestError {
	userDomain.EncryptPassword()

	db, err := gorm.Open(sqlite.Open("usersFromBreadOfPotato.db"), &gorm.Config{})
	if err != nil {
		return rest_error.NewInternalServerError("Não iniciou o Banco de Dados em service/updateUser")
	}

	if err = db.AutoMigrate(&userDomainService{}); err != nil {
		return nil
	}
	userid, _ := strconv.Atoi(userID)

	if userid <= 0 {
		return rest_error.NewBadRequestError("ID invalido (menor ou iqual a 0).")
	}
	ud = &userDomainService{
		FullName: userDomain.GetFullName(),
		Email:    userDomain.GetEmail(),
		Username: userDomain.GetUsername(),
		Password: userDomain.GetPassword(),
	}
	var lastUd userDomainService

	db.First(&lastUd, userID)
	db.Model(&lastUd).Updates(&ud)

	return nil
}
