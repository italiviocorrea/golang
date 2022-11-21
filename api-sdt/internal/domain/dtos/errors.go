package dtos

type Error struct {
	Message    string   `json:"message"`
	Code       int      `json:"code"`
	Name       string   `json:"name"`
	Error      error    `json:"-"`
	Validation []string `json:"validation"`
}

func BindError() *Error {
	return &Error{
		Message: "Erro ao processar a requisição",
		Code:    400,
		Name:    "BIND_ERROR",
	}
}

func ValidationError(errors []string) *Error {
	return &Error{
		Message:    "Ocorreu erro de validação",
		Code:       400,
		Name:       "VALIDAÇÃO",
		Validation: errors,
	}
}
