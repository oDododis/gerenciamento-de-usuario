package service

import (
	"Teste/src/configuration/rest_error"
	"Teste/src/model"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (ud *userDomainService) UpdateUser(userID string, userDomain model.UserDomainInterface) *rest_error.RestError {
	userDomain.EncryptPassword()
	db, err := gorm.Open(sqlite.Open("usersFromBreadOfPotato.db"), &gorm.Config{})
	if err != nil {
		return rest_error.NewInternalServerError("Não iniciou o Banco de Dados em service/createUser")
	}
	err = db.AutoMigrate(&userDomainService{})
	if err != nil {
		return nil
	}

	var lastUd userDomainService

	db.First(&lastUd, userID)
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-")
	fmt.Println(&lastUd)
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-")
	db.Model(&lastUd).Updates(&ud)

	return nil
}
