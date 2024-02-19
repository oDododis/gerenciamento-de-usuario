package request

import "Teste/model"

// Requisita o Email e a sernha com restrições

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,containsany=!@#$%&*()_+"`
}

func (userRequestLogin *LoginRequest) ConvertRequestLoginToModel() *model.User {
	user := &model.User{
		Email:    userRequestLogin.Email,
		Password: userRequestLogin.Password,
	}
	return user
}
