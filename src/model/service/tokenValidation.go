package service

import (
	"Teste/src/configuration/rest_error"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (ud *tokenDomainService) TokenValidation(token string) *rest_error.RestError {
	db, err := gorm.Open(sqlite.Open("usersFromBreadOfPotato.db"), &gorm.Config{})
	if err != nil {
		return rest_error.NewInternalServerError("Não iniciou o Banco de Dados em service/findUser")
	}

	if err = db.AutoMigrate(&tokenDomainService{}); err != nil {
		return rest_error.NewInternalServerError("Não iniciou o Banco de Dados em service/findUser")
	}
	err = db.First(&ud, token).Error
	if err != nil {
		return rest_error.NewUnauthorizedRequestError("Token Invalido")
	}
	return nil
}
