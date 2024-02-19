package request

import "Teste/model"

// Requisita o User com restriçoes

type UserRequest struct {
	FullName string `json:"fullName" binding:"required,min=3,max=150"`
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required,min=3,max=150"`
	Password string `json:"password" binding:"required,min=8,containsany=!@#$%&*()_+"`
}

// Converte o Request para Model

func (userRequest *UserRequest) ConvertRequestToModel() *model.User {
	user := &model.User{
		FullName: userRequest.FullName,
		Email:    userRequest.Email,
		Username: userRequest.Username,
		Password: userRequest.Password,
	}
	return user
}
