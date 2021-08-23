package service

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/model"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/rule"
)

type Nf3eSituacaoService interface {
	Consultar(consSitNF3e rule.ConsSitNF3e) (model.Nf3eSituacao, error)
}
