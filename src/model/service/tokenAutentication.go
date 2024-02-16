package service

import (
	"Teste/src/configuration/rest_error"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (ud *tokenDomainService) TokenAutentication(autenticationToken string) *rest_error.RestError {
	db, err := gorm.Open(sqlite.Open("usersFromBreadOfPotato.db"), &gorm.Config{})
	if err != nil {
		return rest_error.NewInternalServerError("Não iniciou o Banco de Dados em service/tokenAutentication")
	}

	if err = db.AutoMigrate(&tokenDomainService{}); err != nil {
		return rest_error.NewInternalServerError("Não iniciou o Banco de Dados em service/tokenAutentication")
	}
	ud = &tokenDomainService{
		Token: autenticationToken,
		UID:   0,
	}
	if err = db.First(&ud, "token = ?", ud.Token).Error; err != nil {
		return rest_error.NewUnauthorizedRequestError("Token errado.")
	} else {
		return nil
	}
}
