package service

import (
	"Teste/configuration/rest_error"
	"Teste/model"
	"gorm.io/gorm"
	"strconv"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return UserService{db: db}
}

func (ud *UserService) CreateUserServices(userModel *model.User) *rest_error.RestError {

	err := ud.db.First(&ud, "email = ?", userModel.Email).Error // db.First(&user, "username = ?", user.Username).Error
	err = ud.db.First(&ud, "username = ?", userModel.Username).Error

	if err != nil {
		ud.db.Create(&userModel)
	} else {
		return rest_error.NewBadRequestError("Email ou Username existente.")
	}
	return nil
}

func (ud *UserService) DeleteUserServices(userID string) *rest_error.RestError {
	userid, _ := strconv.Atoi(userID)

	if userid <= 0 {
		return rest_error.NewBadRequestError("ID invalido (menor ou iqual a 0).")
	}
	userModel := &model.User{}
	err := ud.db.Delete(&userModel, userID).Error
	if err != nil {
		return rest_error.NewNotFoundError("ID não encontrado.")
	}
	return nil
}

func (ud *UserService) FindUserIDServices(userID string) (*model.User, *rest_error.RestError) {

	userid, _ := strconv.Atoi(userID)
	if userid <= 0 {
		return nil, rest_error.NewNotFoundError("ID invalido (menor ou iqual a 0).")
	}
	userModel := &model.User{}

	err := ud.db.First(&userModel, userID).Error
	if err != nil {
		return nil, rest_error.NewNotFoundError("ID não encontrado.")
	}
	return userModel, nil
}

func (ud *UserService) FindUserEmailServices(userEmail string) (*model.User, *rest_error.RestError) {
	userModel := &model.User{
		Email: userEmail,
	}
	if userEmail == "" {
		return nil, rest_error.NewNotFoundError("Email vazil")
	} else {
		err := ud.db.First(&userModel, "email = ?", userEmail).Error
		if err != nil {
			return nil, rest_error.NewNotFoundError("Email não encontrado.")
		}
		return userModel, nil
	}
}

func (ud *UserService) HowMuchUsers() (int, *rest_error.RestError) {
	userModel := &model.User{}

	err := ud.db.Last(userModel).Error
	if err != nil {
		return 0, rest_error.NewNotFoundError("Não tem Usuarios")
	}
	return int(userModel.ID), nil
}

func (ud *UserService) ListUserIDServices(userID string) (*model.User, *rest_error.RestError) {

	userModel := &model.User{}

	ud.db.First(&userModel, userID)
	if userModel.ID == 0 {
		empty := &model.User{}
		userid, _ := strconv.Atoi(userID)
		empty.ID = uint(userid)
		empty.FullName = "nonexistent"
		empty.Email = "nonexistent"
		empty.Username = "nonexistent"
		return empty, rest_error.NewNotFoundError("Usuario não encontrado.")
	} else {
		return userModel, nil
	}

}

func (ud *UserService) UpdateUserServices(userID string, userModel *model.User) *rest_error.RestError {

	userid, _ := strconv.Atoi(userID)
	if userid <= 0 {
		return rest_error.NewBadRequestError("ID invalido (menor ou iqual a 0).")
	}
	lastUserModel := &model.User{}

	ud.db.First(&lastUserModel, userID)
	ud.db.Model(&lastUserModel).Updates(&userModel)

	return nil
}
