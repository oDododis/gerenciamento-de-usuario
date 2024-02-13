package service

import (
	"Teste/src/configuration/rest_error"
	"Teste/src/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strconv"
)

func (ud *userDomainService) UpdateUser(userID string, userDomain model.UserDomainInterface) *rest_error.RestError {
	userDomain.EncryptPassword()
	db, err := gorm.Open(sqlite.Open("usersFromBreadOfPotato.db"), &gorm.Config{})
	if err != nil {
		return rest_error.NewInternalServerError("Não iniciou o Banco de Dados em service/updateUser")
	}
	err = db.AutoMigrate(&userDomainService{})
	if err != nil {
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
