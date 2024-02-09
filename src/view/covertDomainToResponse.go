package view

import (
	"Teste/src/controller/model/response"
	"Teste/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:       "",
		FullName: userDomain.GetFullName(),
		Email:    userDomain.GetEmail(),
		Username: userDomain.GetUsername(),
		//Birthday: time.Time{},
	}
}
