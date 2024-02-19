package service

import (
	"Teste/configuration/rest_error"
	"Teste/model"
	"errors"
	"gorm.io/gorm"
	"strconv"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (ud *UserService) CreateUserServices(userModel *model.User) *rest_error.RestError {
	var lastUserModel *model.User
	lastUserModel = nil
	err := ud.db.First(&lastUserModel, "email = ?", userModel.Email).Error
	if lastUserModel.ID != 0 {
		return rest_error.NewBadRequestError("Existing email.")
	}

	err = ud.db.First(&lastUserModel, "username = ?", userModel.Username).Error
	if lastUserModel.ID != 0 {
		return rest_error.NewBadRequestError("Existing username.")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return rest_error.NewRestError(err.Error(), err.Error(), 505, nil)
	}
	ud.db.Create(&userModel)
	return nil
}

func (ud *UserService) DeleteUserServices(userID string) *rest_error.RestError {
	userid, _ := strconv.Atoi(userID)

	if userid <= 0 {
		return rest_error.NewBadRequestError("Invalid ID (less than or equal to 0).")
	}
	userModel := &model.User{}
	err := ud.db.Delete(&userModel, userID).Error
	if err != nil {
		return rest_error.NewNotFoundError("ID not found.")
	}
	return nil
}

func (ud *UserService) FindUserIDServices(userID string) (*model.User, *rest_error.RestError) {

	userid, _ := strconv.Atoi(userID)
	if userid <= 0 {
		return nil, rest_error.NewNotFoundError("Invalid ID (less than or equal to 0).")
	}
	userModel := &model.User{}

	err := ud.db.First(&userModel, userID).Error
	if err != nil {
		return nil, rest_error.NewNotFoundError("ID not found.")
	}
	return userModel, nil
}

func (ud *UserService) FindUserEmailServices(userEmail string) (*model.User, *rest_error.RestError) {
	userModel := &model.User{
		Email: userEmail,
	}
	if userEmail == "" {
		return nil, rest_error.NewNotFoundError("Empty email.")
	} else {
		err := ud.db.First(&userModel, "email = ?", userEmail).Error
		if err != nil {
			return nil, rest_error.NewNotFoundError("Email not found.")
		}
		return userModel, nil
	}
}

func (ud *UserService) HowMuchUsers() (int, *rest_error.RestError) {
	userModel := &model.User{}

	err := ud.db.Last(userModel).Error
	if err != nil {
		return 0, rest_error.NewNotFoundError("There are no users.")
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
		return empty, rest_error.NewNotFoundError("User not found.")
	} else {
		return userModel, nil
	}
}

func (ud *UserService) UpdateUserServices(userID string, userModel *model.User) *rest_error.RestError {

	userid, _ := strconv.Atoi(userID)
	if userid <= 0 {
		return rest_error.NewBadRequestError("Invalid ID (less than or equal to 0).")
	}
	lastUserModel := &model.User{}

	err := ud.db.First(&lastUserModel, "email = ?", userModel.Email).Error
	if lastUserModel.ID != 0 && lastUserModel.ID != userModel.ID {
		return rest_error.NewBadRequestError("Existing email.")
	}
	lastUserModel.ID = 0
	err = ud.db.First(&lastUserModel, "username = ?", userModel.Username).Error
	if lastUserModel.ID != 0 && lastUserModel.ID != userModel.ID {
		return rest_error.NewBadRequestError("Existing username.")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return rest_error.NewRestError(err.Error(), err.Error(), 505, nil)
	}
	ud.db.Model(&lastUserModel).Updates(&userModel)

	return nil
}
