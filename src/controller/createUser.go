package controller

import (
	"Teste/src/configuration/validation"
	"Teste/src/controller/model/request"
	"Teste/src/model"
	"Teste/src/view"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
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

	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
}
