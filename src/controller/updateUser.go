package controller

import (
	"Teste/src/configuration/validation"
	"Teste/src/controller/model/request"
	"Teste/src/model"
	"Teste/src/view"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Recebe o Request e manda atualizar os dados

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	//autenticationToken := c.Request.Header.Get("Authorization")
	//if err := uc.tokenService.TokenValidation(autenticationToken); err != nil {
	//	c.JSON(err.Code, err)
	//	return
	//}
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restError := validation.ValidateUserError(err)
		c.JSON(restError.Code, restError)
	}

	domain := model.NewUserDomain(
		userRequest.FullName,
		userRequest.Email,
		userRequest.Username,
		userRequest.Password,
		//userResquest.Birthday
	)
	userID := c.Param("userID")

	if err := uc.service.UpdateUserServices(userID, domain); err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusAccepted, view.ConvertDomainToResponse(domain))
}
