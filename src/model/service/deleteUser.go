package service

import (
	"Teste/src/configuration/rest_error"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strconv"
)

func (ud *userDomainService) DeleteUser(userID string) *rest_error.RestError {
	db, err := gorm.Open(sqlite.Open("usersFromBreadOfPotato.db"), &gorm.Config{})
	if err != nil {
		return rest_error.NewInternalServerError("Não iniciou o Banco de Dados em service/deleteUser")
	}
	err = db.AutoMigrate(&userDomainService{})
	if err != nil {
		return nil
	}
	userid, _ := strconv.Atoi(userID)

	if userid <= 0 {
		return rest_error.NewBadRequestError("ID invalido (menor ou iqual a 0).")
	}
	var lastUd userDomainService
	db.Delete(&lastUd, userID)

	return nil
}
