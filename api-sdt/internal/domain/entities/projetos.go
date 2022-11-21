package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ambiente struct {
	Id   int8   `json:"id"`
	Nome string `json:"nome"`
	Uri  string `json:"uri"`
}

type Metodo struct {
	Nome              string `json:"nome"`
	Compactar         bool   `json:"compactar"`
	Envolope_template string `json:"envolope_Template"`
}

type Esquema struct {
	Nome      string `json:"nome"`
	Descricao string `json:"descricao"`
}

type Evento struct {
	Nome    string `json:"nome"`
	Esquema string `json:"esquema"`
}

type Servico struct {
	Nome       string    `json:"nome"`
	Nome_wsdl  string    `json:"nome_Wsdl"`
	Assincrono bool      `json:"assincrono"`
	Metodos    []Metodo  `json:"metodos"`
	Esquemas   []Esquema `json:"esquemas"`
	Eventos    []Evento  `json:"eventos"`
}

type Versao struct {
	Versao   string    `json:"versao"`
	Contexto string    `json:"contexto"`
	Servicos []Servico `json:"servicos"`
}

type Projeto struct {
	ID        primitive.ObjectID `bson:"_id"`
	Nome      string             `json:"nome"`
	Descricao string             `json:"descricao"`
	Modelo    int8               `json:"modelo"`
	Versoes   []Versao           `json:"versoes"`
	Ambientes []Ambiente         `json:"ambientes"`
}
