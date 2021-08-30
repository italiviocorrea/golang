package dtos

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/models/vos"
)

type ConsSitNF3e struct {
	Versao       string          `xml:"versao"`
	TpAmb        string          `xml:"tpAmb"`
	XServ        string          `xml:"xServ"`
	ChNF3e       string          `xml:"chNF3e"`
	ChNF3eDecode vos.ChaveAcesso `xml:"chNF3eDecode,omitempty"`
}
