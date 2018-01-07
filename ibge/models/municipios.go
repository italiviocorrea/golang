package models

import (
	"github.com/italiviocorrea/golang/commons"
	"errors"
	"regexp"
)

type (

	// Representacao da entidade Municipio
	Municipio struct {
		Codigo int64  `json:"codigo" description:"Codigo IBGE do Municipio"`
		Nome   string `json:"nome"   description:"Nome do Municipio"`
		Uf     Uf     `json:"uf"     description:"UF do Municipio"`
	}

	// Representacao da resposta quando for um so municipio
	MunicipioResponse struct {
		Codigo int64          `json:"codigo" description:"Codigo IBGE do Municipio"`
		Nome   string         `json:"nome"   description:"Nome do Municipio"`
		Uf     Uf             `json:"uf"     description:"UF do Municipio"`
		Links  []commons.Link `json:"_links"`
	}

	// Representacao da resposta quando for uma lista de municipio
	MunicipiosResponse struct {
		Codigo int64          `json:"codigo" description:"Codigo IBGE do Municipio"`
		Nome   string         `json:"nome"   description:"Nome do Municipio"`
		Uf     string         `json:"uf" 	   description:"UF do Municipio"`
		Links  []commons.Link `json:"_links"`
	}

	// Representacao da resposta do endpoint (ModelView)
	MunicipioResource struct {
		Data    MunicipioResponse `json:"data"`
		Status  commons.StatusResponse   `json:"status"`
		Version string                   `json:"version"`
	}

	MunicipiosResource struct {
		Data       []MunicipiosResponse  `json:"data"`
		Pagination commons.Pagination          `json:"pagination,omitempty"`
		Status     commons.StatusResponse      `json:"status"`
		Version    string                      `json:"version"`
		Link       commons.Link                `json:"link"`
	}


)

var (
	ErrInvalidCodigoMunicipio = errors.New("Codigo do Municipio invalido!")
	ErrInvalidNomeMunicipio = errors.New("Nome do Municipio invalido")
	ErrInvalidCodigoUFMunicipio = errors.New("Codigo da UF do Municipio invalido!")
	ErrInvalidNomeUFMunicipio = errors.New("Nome da UF do Municipio invalido")
	ErrInvalidSiglaUFMunicipio = errors.New("Sigla da UF do Municipio invalido")
)


// Validate - implementation of the RequestValidation interface
func (t Municipio) Validate() error {

	var validID = regexp.MustCompile(`[A-Z]{2}`)

	if t.Codigo <= 0  {
		return ErrInvalidCodigoMunicipio
	}
	if t.Nome == "" {
		return ErrInvalidNomeMunicipio
	}
	if t.Uf.Codigo <= 0  {
		return ErrInvalidCodigoUFMunicipio
	}
	if t.Uf.Nome == "" {
		return ErrInvalidNomeUFMunicipio
	}
	if t.Uf.Sigla == "" || !validID.MatchString(t.Uf.Sigla){
		return ErrInvalidSiglaUFMunicipio
	}
	return nil
}


