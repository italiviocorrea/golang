package entities

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/models/vos"
)

type ConsSitNF3e struct {
	Versao       string
	TpAmb        string
	XServ        string
	ChNF3e       string
	ChNF3eDecode vos.ChaveAcesso
}
