package view

//Exporta o Domain usando o Resonse

import (
	"Teste/src/controller/model/response"
	"Teste/src/model"
)

//Cria a conversão do Dominio para a Resposta

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:       userDomain.GetID(),
		FullName: userDomain.GetFullName(),
		Email:    userDomain.GetEmail(),
		Username: userDomain.GetUsername(),
		//Birthday: time.Time{},
	}
}
