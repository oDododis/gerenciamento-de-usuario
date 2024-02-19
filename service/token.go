package service

import (
	"Teste/configuration/rest_error"
	"Teste/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TokenService struct {
	db *gorm.DB
}

func NewTokenService(db *gorm.DB) UserService {
	return UserService{db: db}
}

func (ud *TokenService) LoginServices(userModel *model.User) (string, *rest_error.RestError) {
	token := uuid.New().String()

	if userModel.Email == "" {
		return "", rest_error.NewNotFoundError("Email vazil")
	}
	lastUd := &model.User{}
	err := ud.db.First(&lastUd, "email = ?", userModel.Email).Error
	if err != nil {
		return "", rest_error.NewNotFoundError("Email não existe.")
	} else if lastUd.Password != userModel.Password {
		return "", rest_error.NewForbiddenError("Senha incorreta.")
	} else {
		autentication := &model.Token{
			Token: token,
			UID:   lastUd.ID,
		}
		if err = ud.db.AutoMigrate(&model.Token{}); err != nil {
			return "", rest_error.NewInternalServerError("Não iniciou o Banco de Dados em service/login")
		}
		ud.db.Create(&autentication)

		return autentication.Token, nil
	}
}

func (ud *TokenService) TokenAutentication(autenticationToken string) *rest_error.RestError {

	token := &model.Token{
		Token: autenticationToken,
		UID:   0,
	}
	if err := ud.db.First(&token, "token = ?", token.Token).Error; err != nil {
		return rest_error.NewUnauthorizedRequestError("TokenServiceInterface errado.")
	} else {
		return nil
	}
}
