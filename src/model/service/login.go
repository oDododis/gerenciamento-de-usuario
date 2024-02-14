package service

import (
	"Teste/src/configuration/rest_error"
	"Teste/src/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//Recebe os campos do controller, valida o email, a senha e retorna as informações

func (ud *userDomainService) LoginServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_error.RestError) {
	userDomain.EncryptPassword()
	db, err := gorm.Open(sqlite.Open("usersFromBreadOfPotato.db"), &gorm.Config{})
	if err != nil {
		return nil, rest_error.NewInternalServerError("Não iniciou o Banco de Dados em service/login")
	}

	if err = db.AutoMigrate(&userDomainService{}); err != nil {
		return nil, rest_error.NewInternalServerError("Não iniciou o Banco de Dados em service/findUser")
	}
	ud = &userDomainService{
		Email:    userDomain.GetEmail(),
		Password: userDomain.GetPassword(),
	}
	if ud.Email == "" {
		return nil, rest_error.NewNotFoundError("Email vazil")
	}
	var lastUd userDomainService
	err = db.First(&lastUd, "email = ?", ud.Email).Error
	if err != nil {
		return nil, rest_error.NewNotFoundError("Email não existe.")
	} else if lastUd.Password != ud.Password {
		return nil, rest_error.NewForbiddenError("Senha incorreta.")
	} else {
		return &lastUd, nil
	}
}
