package response

//Exporta o Domain usando o Resonse

import (
	"Teste/model"
)

//Cria a conversão do Dominio para a Resposta

func ConvertDomainToResponse(userModel *model.User) UserResponse {
	return UserResponse{
		ID:       userModel.ID,
		FullName: userModel.FullName,
		Email:    userModel.Email,
		Username: userModel.Username,
	}
}
