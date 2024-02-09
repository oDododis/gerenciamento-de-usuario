package service

import (
	"Teste/src/configuration/rest_error"
	"Teste/src/model"
	"fmt"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_error.RestError {

	userDomain.EncryptPassword()
	fmt.Println(userDomain.GetPassword())
	return nil
}
