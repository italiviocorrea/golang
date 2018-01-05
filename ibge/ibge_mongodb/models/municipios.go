package models

import (
	"github.com/italiviocorrea/golang/commons"
	"gopkg.in/mgo.v2/bson"
)

type (

	UFMunicipio struct {
		Codigo int    			`json:"codigo" description:"Codigo IBGE da UF"`
		Nome   string 			`json:"nome"   description:"Nome da UF"`
		Sigla  string 			`json:"sigla"  description:"Sigla da UF"`
	}

	// Representacao da entidade Municipio
	Municipio struct {
		ID     bson.ObjectId 	`bson:"_id,omitempty"`
		Codigo int64  			`json:"codigo" description:"Codigo IBGE do Municipio"`
		Nome   string 			`json:"nome"   description:"Nome do Municipio"`
		Uf     UFMunicipio		`json:"uf"     description:"UF do Municipio"`
	}

	// Representacao da resposta quando for um so municipio
	MunicipioResponse struct {
		ID     bson.ObjectId  `bson:"_id,omitempty"`
		Codigo int64          `json:"codigo" description:"Codigo IBGE do Municipio"`
		Nome   string         `json:"nome"   description:"Nome do Municipio"`
		Uf     UFMunicipio    `json:"uf"     description:"UF do Municipio"`
		Links  []commons.Link `json:"_links"`
	}

	// Representacao da resposta quando for uma lista de municipio
	MunicipiosResponse struct {
		ID     bson.ObjectId  `bson:"_id,omitempty"`
		Codigo int64          `json:"codigo" description:"Codigo IBGE do Municipio"`
		Nome   string         `json:"nome"   description:"Nome do Municipio"`
		Uf     string         `json:"uf" 	 description:"UF do Municipio"`
		Links  []commons.Link `json:"_links"`
	}

	// Representacao da resposta do endpoint (ModelView)
	MunicipioResource struct {
		Data    MunicipioResponse 		`json:"data"`
		Status  commons.StatusResponse  `json:"status"`
		Version string                  `json:"version"`
	}

	MunicipiosResource struct {
		Data       []MunicipiosResponse  	`json:"data"`
		Pagination commons.Pagination       `json:"pagination,omitempty"`
		Status     commons.StatusResponse   `json:"status"`
		Version    string                   `json:"version"`
		Link       commons.Link             `json:"link"`
	}


)
