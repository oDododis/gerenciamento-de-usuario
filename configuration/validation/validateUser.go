package validation

//Melhora a vualisação do erro, colocando as mesagens de erro nos campos criados no RestError, validando os erros

import (
	"Teste/configuration/rest_error"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	translate ut.Translator
)

//Inicia a validação

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		translate, _ = unt.GetTranslator("en")
		err := en_translation.RegisterDefaultTranslations(val, translate)
		if err != nil {
			return
		}
	}
}

// Coloca os error validados nos campos corretos em RestError

func ValidateUserError(validationErr error) *rest_error.RestError {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return rest_error.NewBadRequestError("Invalid field type.")
	} else if errors.As(validationErr, &jsonValidationError) {
		var errorsCauses []rest_error.Causes

		for _, e := range validationErr.(validator.ValidationErrors) {
			cause := rest_error.Causes{
				Message: e.Translate(translate),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return rest_error.NewBadRequestValidationError("Invalid field sent.", errorsCauses)
	} else {
		return rest_error.NewBadRequestError("Error trying to convert a field.")
	}
}
