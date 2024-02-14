package model

//Criando um Dominio de Usuario para coletar as informações localmente e distribuir individualmente

import (
	"crypto/md5"
	"encoding/hex"
)

//Cria os comandos para coletar os campos enviados

type UserDomainInterface interface {
	GetFullName() string
	GetEmail() string
	GetUsername() string
	GetPassword() string

	EncryptPassword()
}

//Cria o Domionio do Usuario privado

func NewUserDomain(fullName, email, username, password string) UserDomainInterface {
	return &userDomain{fullName, email, username, password}
}
func NewUserDomainLogin(email, password string) UserDomainInterface {
	return &userDomain{email: email, password: password}
}

type userDomain struct {
	fullName string
	email    string
	username string
	password string
	//birthday time.Time
}

//Funções para pegar os dados do usuario de forma privada e apenas aqui

func (ud *userDomain) GetFullName() string {
	return ud.fullName
}
func (ud *userDomain) GetEmail() string {
	return ud.email
}
func (ud *userDomain) GetUsername() string {
	return ud.username
}
func (ud *userDomain) GetPassword() string {
	return ud.password
}

//Emcrypta a Senha com md5 com a senha local.

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}
