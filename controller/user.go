package controller

import (
	"Teste/configuration/validation"
	"Teste/controller/request"
	"Teste/controller/response"
	"Teste/service"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type UserController struct {
	userService  service.UserServiceInterface
	tokenServise service.TokenServiceInterface
}

func NewUserController(userServiceInterface service.UserServiceInterface, tokenServiceInterface service.TokenServiceInterface) UserController {
	return UserController{
		userService:  userServiceInterface,
		tokenServise: tokenServiceInterface,
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restError := validation.ValidateUserError(err)
		c.JSON(restError.Code, restError)
		return
	}

	user := userRequest.ConvertRequestToModel()
	user.Password = encryptPassword(userRequest.Password)

	if err := uc.userService.CreateUserServices(user); err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusCreated, response.ConvertDomainToResponse(user))
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	autenticationToken := c.Request.Header.Get("Authorization")
	token := strings.Split(autenticationToken, " ")

	if err := uc.tokenServise.TokenAutentication(token[1]); err != nil {
		c.JSON(err.Code, err)
		return
	}

	userID := c.Param("userID")
	if err := uc.userService.DeleteUserServices(userID); err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusAccepted, "UserServiceInterface "+userID+" excluido")
}

func (uc *UserController) FindUserID(c *gin.Context) {
	autenticationToken := c.Request.Header.Get("Authorization")
	token := strings.Split(autenticationToken, " ")
	if err := uc.tokenServise.TokenAutentication(token[1]); err != nil {

		c.JSON(err.Code, err)
		return
	}
	userID := c.Param("userID")
	userDomain, err := uc.userService.FindUserIDServices(userID)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusAccepted, response.ConvertDomainToResponse(userDomain))
}

func (uc *UserController) FindUserEmail(c *gin.Context) {
	userEmail := c.Param("userEmail")
	userDomain, err := uc.userService.FindUserEmailServices(userEmail)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusAccepted, response.ConvertDomainToResponse(userDomain))
}

func (uc *UserController) Login(c *gin.Context) {
	var userRequest request.LoginRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restError := validation.ValidateUserError(err)
		c.JSON(restError.Code, restError)
		return
	}

	user := userRequest.ConvertRequestLoginToModel()
	user.Password = encryptPassword(userRequest.Password)

	userDomain, err := uc.tokenServise.LoginServices(user)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusAccepted, "TokenServiceInterface: "+userDomain)
}

func (uc *UserController) UsersList(c *gin.Context) {

	userID, err := uc.userService.HowMuchUsers()
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	for i := 1; i <= userID; i++ {
		userDomain, err := uc.userService.FindUserIDServices(strconv.Itoa(i))
		if err != nil {
			c.JSON(err.Code, err)
			return
		}
		c.JSON(http.StatusAccepted, "=-=-=-=-=-=-=-=-=-=-=-=-=-=")
		c.JSON(http.StatusAccepted, response.ConvertDomainToResponse(userDomain))
	}
	c.JSON(http.StatusAccepted, "=-=-=-=-=-=-=-=-=-=-=-=-=-=")
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restError := validation.ValidateUserError(err)
		c.JSON(restError.Code, restError)
		return
	}

	autenticationToken := c.Request.Header.Get("Authorization")
	token := strings.Split(autenticationToken, " ")
	if err := uc.tokenServise.TokenAutentication(token[1]); err != nil {
		c.JSON(err.Code, err)
		return
	}

	user := userRequest.ConvertRequestToModel()
	user.Password = encryptPassword(userRequest.Password)

	userID := c.Param("userID")

	if err := uc.userService.UpdateUserServices(userID, user); err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusAccepted, response.ConvertDomainToResponse(user))
}

func encryptPassword(password string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(password))
	password = hex.EncodeToString(hash.Sum(nil))
	return password
}
