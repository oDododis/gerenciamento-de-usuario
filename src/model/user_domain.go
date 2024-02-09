package model

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomainInterface interface {
	GetFullName() string
	GetEmail() string
	GetUsername() string
	GetPassword() string

	EncryptPassword()
}

func NewUserDomain(fullName, email, username, password string) UserDomainInterface {
	return &userDomain{fullName, email, username, password}
}

type userDomain struct {
	fullName string
	email    string
	username string
	password string
	//birthday time.Time
}

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

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}
