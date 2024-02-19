package service

import (
	"Teste/configuration/rest_error"
	"Teste/model"
)

type UserServiceInterface interface {
	CreateUserServices(userModel *model.User) *rest_error.RestError
	DeleteUserServices(userID string) *rest_error.RestError
	FindUserIDServices(userID string) (*model.User, *rest_error.RestError)
	FindUserEmailServices(userEmail string) (*model.User, *rest_error.RestError)
	HowMuchUsers() (int, *rest_error.RestError)
	ListUserIDServices(userID string) (*model.User, *rest_error.RestError)
	UpdateUserServices(userID string, userModel *model.User) *rest_error.RestError
}
type TokenServiceInterface interface {
	LoginServices(userModel *model.User) (string, *rest_error.RestError)
	TokenAutentication(autenticationToken string) *rest_error.RestError
}
