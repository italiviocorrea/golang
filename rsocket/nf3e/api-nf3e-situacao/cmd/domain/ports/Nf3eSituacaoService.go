package ports

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/models/dtos"
)

type Nf3eSituacaoService interface {
	GetNf3eSituacao(consSitNF3e dtos.ConsSitNF3e) (dtos.RetConsSitNF3e, error)
}
