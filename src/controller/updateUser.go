package controller

import (
	"Teste/src/configuration/validation"
	"Teste/src/controller/model/request"
	"Teste/src/model"
	"Teste/src/view"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restError := validation.ValidateUserError(err)
		c.JSON(restError.Code, restError)
	}

	domain := model.NewUserDomain(
		//userRequest gorm.Model
		userRequest.FullName,
		userRequest.Email,
		userRequest.Username,
		userRequest.Password,
		//userResquest.Birthday
	)
	ID := "2"
	if err := uc.service.UpdateUser(ID, domain); err != nil {

		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
}
