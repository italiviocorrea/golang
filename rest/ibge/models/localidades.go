package models

import (
	"errors"
	"github.com/italiviocorrea/golang/commons"
)

type (

	// Representacao da entidade Localidade
	Localidade struct {
		Codigo          int64  `json:"codigo" description:"Codigo IBGE do Municipio"`
		NomeMunicipio   string `json:"nome_municipio"   description:"Nome do Municipio"`
		NomeUf          string `json:"nome_uf"   description:"Nome da UF"`
		NomeLocalidade  string `json:"nome_localidade"   description:"Nome da Localidade"`
		NomeMesoRegiao  string `json:"nome_meso_regiao"   description:"Nome da Meso Região"`
		NomeMicroRegiao string `json:"nome_micro_regiao"   description:"Nome da Micro Região"`
		Longitude       string `json:"longitude"   description:"Longitude"`
		Latitude        string `json:"latitude"   description:"Latitude"`
		Altitude        string `json:"altitude"   description:"Altitude"`
	}

	// Representacao da resposta quando for uma só localidade
	LocalidadeResponse struct {
		Codigo          int64          `json:"codigo" description:"Codigo IBGE do Municipio"`
		NomeMunicipio   string         `json:"nome_municipio"   description:"Nome do Municipio"`
		NomeUf          string         `json:"nome_uf"   description:"Nome da UF"`
		NomeLocalidade  string         `json:"nome_localidade"   description:"Nome da Localidade"`
		NomeMesoRegiao  string         `json:"nome_meso_regiao"   description:"Nome da Meso Região"`
		NomeMicroRegiao string         `json:"nome_micro_regiao"   description:"Nome da Micro Região"`
		Longitude       string         `json:"longitude"   description:"Longitude"`
		Latitude        string         `json:"latitude"   description:"Latitude"`
		Altitude        string         `json:"altitude"   description:"Altitude"`
		Links           []commons.Link `json:"_links"`
	}

	// Representacao da resposta quando for uma lista de municipio
	LocalidadesResponse struct {
		Codigo         int64          `json:"codigo" description:"Codigo IBGE do Municipio"`
		NomeMunicipio  string         `json:"nome_municipio"   description:"Nome do Municipio"`
		NomeUf         string         `json:"nome_uf"   description:"Nome da UF"`
		NomeLocalidade string         `json:"nome_localidade"   description:"Nome da Localidade"`
		Links          []commons.Link `json:"_links"`
	}

	// Representacao da resposta do endpoint (ModelView)
	LocalidadeResource struct {
		Data    LocalidadeResponse     `json:"data"`
		Status  commons.StatusResponse `json:"status"`
		Version string                 `json:"version"`
	}

	LocalidadesResource struct {
		Data       []LocalidadesResponse  `json:"data"`
		Pagination commons.Pagination     `json:"pagination,omitempty"`
		Status     commons.StatusResponse `json:"status"`
		Version    string                 `json:"version"`
		Link       commons.Link           `json:"link"`
	}
)

var (
	ErrInvalidCodigoLocalidade = errors.New("Codigo do Localidade invalido!")
	ErrInvalidNomeLocalidade   = errors.New("Nome do Localidade invalido")
	//ErrInvalidCodigoUFLocalidade = errors.New("Codigo da UF do Localidade invalido!")
	ErrInvalidNomeUFLocalidade = errors.New("Nome da UF do Localidade invalido")
	//ErrInvalidSiglaUFLocalidade = errors.New("Sigla da UF do Localidade invalido")
)

// Validate - implementation of the RequestValidation interface
func (t Localidade) Validate() error {

	//var validID = regexp.MustCompile(`[A-Z]{2}`)

	if t.Codigo <= 0 {
		return ErrInvalidCodigoLocalidade
	}
	if t.NomeMunicipio == "" {
		return ErrInvalidNomeLocalidade
	}
	if t.NomeUf == "" {
		return ErrInvalidNomeUFLocalidade
	}
	return nil
}
