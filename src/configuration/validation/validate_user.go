package validation

import (
	"Teste/src/configuration/rest_error"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate  = validator.New()
	translate ut.Translator
)

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
func ValidateUserError(validation_err error) *rest_error.RestError {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_error.NewBadRequestError("giTipo de campo invalido")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []rest_error.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_error.Causes{
				Message: e.Translate(translate),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return rest_error.NewBadRequestValidationError("Alguma coisa ta errada", errorsCauses)
	} else {
		return rest_error.NewBadRequestError("Erro tentando converter um campo")
	}
}
