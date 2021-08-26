package services

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/models/dtos"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/models/entities"
)

type Nf3eSituacaoService interface {
	GetNf3eSituacao(consSitNF3e dtos.ConsSitNF3e) (entities.Nf3eSituacao, error)
}
