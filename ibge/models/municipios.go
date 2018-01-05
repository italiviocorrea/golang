package models

import "com/ItalivioCorrea/commons"

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
