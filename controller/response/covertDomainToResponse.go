package response

import (
	"Teste/model"
)

// Converte o Model para Response

func ConvertModelToResponse(userModel *model.User) UserResponse {
	return UserResponse{
		ID:       userModel.ID,
		FullName: userModel.FullName,
		Email:    userModel.Email,
		Username: userModel.Username,
	}
}
