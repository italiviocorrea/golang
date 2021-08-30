package ports

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/internal/adapters/primary/dtos"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/models/entities"
)

type Nf3eSituacaoServicePort interface {
	GetNf3eSituacao(consSitNF3e entities.ConsSitNF3e) (dtos.RetConsSitNF3e, error)
}
