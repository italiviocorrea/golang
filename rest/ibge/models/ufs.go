package models

import (
	"github.com/italiviocorrea/golang/commons"
	"errors"
	"regexp"
)

type (

	// Representacao da entidade UF
	Uf struct {
		Codigo int    `json:"codigo" description:"Codigo IBGE da UF"`
		Nome   string `json:"nome"   description:"Nome da UF"`
		Sigla  string `json:"sigla"  description:"Sigla da UF"`
	}

	// Representacao da entidade de resposta da UF
	UFResponse struct {
		Codigo int            `json:"codigo" description:"Codigo IBGE da UF"`
		Nome   string         `json:"nome"   description:"Nome da UF"`
		Sigla  string         `json:"sigla"  description:"Sigla da UF"`
		Links  []commons.Link `json:"_links"`
	}

	// Representacao da resposta do endpoint (ModelView)
	UfResource struct {
		Data    []UFResponse      `json:"data"`
		Status  commons.StatusResponse `json:"status"`
		Version string                 `json:"version"`
	}

	UfsResource struct {
		Data       []UFResponse    `json:"data"`
		Pagination commons.Pagination     `json:"pagination,omitempty"`
		Status     commons.StatusResponse `json:"status"`
		Version    string                 `json:"version"`
		Link       commons.Link           `json:"link"`
	}

)

var (
	ErrInvalidCodigoUF = errors.New("Codigo da UF invalido!")
	ErrInvalidNomeUF = errors.New("Nome da UF invalido")
	ErrInvalidSiglaUF = errors.New("Sigla da UF invalida")
)


// Validate - implementation of the RequestValidation interface
func (t Uf) Validate() error {

	var validID = regexp.MustCompile(`[A-Z]{2}`)

	if t.Codigo <= 0  {
		return ErrInvalidCodigoUF
	}
	if t.Nome == "" {
		return ErrInvalidNomeUF
	}
	if t.Sigla == "" || !validID.MatchString(t.Sigla) {
		return ErrInvalidSiglaUF
	}
	return nil
}
