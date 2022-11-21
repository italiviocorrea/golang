package validators

import (
	"errors"
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	validator "gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

var (
	uni *ut.UniversalTranslator
)

// CustomValidator ..
type CustomValidator struct {
	Validator *validator.Validate
	Trans     ut.Translator
}

//Validate is func test valid object
func (cv *CustomValidator) Validate(i interface{}) error {

	err := cv.Validator.Struct(i)
	if err != nil {

		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			return errors.New(e.Translate(cv.Trans))
		}
	}
	return nil
}

//InitCustomValidator is func init CustomValid
func InitCustomValidator() *CustomValidator {
	en := pt_BR.New()
	uni = ut.New(en, en)
	trans, _ := uni.GetTranslator("pt_BR")

	validate := validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)
	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "Por favor, insira o valor {0}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	validate.RegisterTranslation("min", trans, func(ut ut.Translator) error {
		return ut.Add("min", "Insira um valor {0} maior que {1} (caracter/numérico).", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("min", fe.Field(), fe.Param())

		return t
	})

	validate.RegisterTranslation("max", trans, func(ut ut.Translator) error {
		return ut.Add("max", "Insira um valor {0} menor que {1} (caractere/numérico).", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("max", fe.Field(), fe.Param())

		return t
	})

	validate.RegisterTranslation("excludesall", trans, func(ut ut.Translator) error {
		return ut.Add("excludesall", "Insira o valor {0} sem o caractere {1}.", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("excludesall", fe.Field(), fe.Param())

		return t
	})

	validate.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "O valor {0} é inválido.", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())

		return t
	})
	return &CustomValidator{Validator: validate, Trans: trans}
}
