package service

import (
	"Teste/src/configuration/rest_error"
	"Teste/src/model"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//Recebe os campos do controller, valida o email, a senha e retorna as informações

func (ud *userDomainService) LoginServices(userDomain model.UserDomainInterface) (string, *rest_error.RestError) {
	userDomain.EncryptPassword()
	token := uuid.New().String()

	db, err := gorm.Open(sqlite.Open("usersFromBreadOfPotato.db"), &gorm.Config{})
	if err != nil {
		return "", rest_error.NewInternalServerError("Não iniciou o Banco de Dados em service/login")
	}

	if err = db.AutoMigrate(&userDomainService{}); err != nil {
		return "", rest_error.NewInternalServerError("Não iniciou o Banco de Dados em service/login")
	}

	ud = &userDomainService{
		Email:    userDomain.GetEmail(),
		Password: userDomain.GetPassword(),
	}
	if ud.Email == "" {
		return "", rest_error.NewNotFoundError("Email vazil")
	}
	var lastUd userDomainService
	err = db.First(&lastUd, "email = ?", ud.Email).Error
	if err != nil {
		return "", rest_error.NewNotFoundError("Email não existe.")
	} else if lastUd.Password != ud.Password {
		return "", rest_error.NewForbiddenError("Senha incorreta.")
	} else {
		autentication := &tokenDomainService{
			Token:  token,
			UserID: lastUd.ID,
		}
		if err = db.AutoMigrate(&tokenDomainService{}); err != nil {
			return "", rest_error.NewInternalServerError("Não iniciou o Banco de Dados em service/login")
		}
		db.Create(&autentication)

		return autentication.Token, nil
	}
}
