package dtos

import "github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/models/vos"

type ConsSitNF3e struct {
	Versao       string `json:"versao"`
	TpAmb        string `json:"tpAmb"`
	XServ        string `json:"xServ"`
	ChNF3e       string `json:"chNF3e"`
	ChNF3eDecode vos.ChaveAcesso
}
