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

// Estrutura do Controle

type UserController struct {
	userService  service.UserServiceInterface
	tokenServise service.TokenServiceInterface
}

// Cria um novo controle

func NewUserController(userServiceInterface service.UserServiceInterface, tokenServiceInterface service.TokenServiceInterface) *UserController {
	return &UserController{
		userService:  userServiceInterface,
		tokenServise: tokenServiceInterface,
	}
}

//Cria o usuário

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
	c.JSON(http.StatusCreated, response.ConvertModelToResponse(user))
}

//Exclui o usuário

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
	c.JSON(http.StatusAccepted, "User "+userID+" was deleted.")
}

//Procura o usuário por ID

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
	c.JSON(http.StatusAccepted, response.ConvertModelToResponse(userDomain))
}

// Procura o usuário por Email

func (uc *UserController) FindUserEmail(c *gin.Context) {
	autenticationToken := c.Request.Header.Get("Authorization")
	token := strings.Split(autenticationToken, " ")
	if err := uc.tokenServise.TokenAutentication(token[1]); err != nil {
		c.JSON(err.Code, err)
		return
	}

	userEmail := c.Param("userEmail")
	userDomain, err := uc.userService.FindUserEmailServices(userEmail)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusAccepted, response.ConvertModelToResponse(userDomain))
}

// Faz o login do usuário e retorna um token

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

// Cria uma lista dos usuários existentes no banco de dados

func (uc *UserController) UsersList(c *gin.Context) {
	autenticationToken := c.Request.Header.Get("Authorization")
	token := strings.Split(autenticationToken, " ")
	if err := uc.tokenServise.TokenAutentication(token[1]); err != nil {
		c.JSON(err.Code, err)
		return
	}

	userID, err := uc.userService.HowMuchUsers()
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	for i := 1; i <= userID; i++ {
		userModel, err := uc.userService.ListUserIDServices(strconv.Itoa(i))

		c.JSON(http.StatusAccepted, "=-=-=-=-=-=-=-=-=-=-=-=-=-=")
		c.JSON(http.StatusAccepted, response.ConvertModelToResponse(userModel))
		if err != nil {
			c.JSON(err.Code, err)
		}
	}
	c.JSON(http.StatusAccepted, "=-=-=-=-=-=-=-=-=-=-=-=-=-=")
}

// Atualiza as informações do usuário

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
	userid, _ := strconv.Atoi(userID)
	user.ID = uint(userid)

	if err := uc.userService.UpdateUserServices(userID, user); err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusAccepted, response.ConvertModelToResponse(user))
}

// criptografa a senha do usuário usando md5

func encryptPassword(password string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(password))
	password = hex.EncodeToString(hash.Sum(nil))
	return password
}
