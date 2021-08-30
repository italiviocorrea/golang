package ports

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/internal/adapters/primary/dtos"
	dtos2 "github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/models/entities"
)

type Nf3eSituacaoServicePort interface {
	GetNf3eSituacao(consSitNF3e dtos2.ConsSitNF3e) (dtos.RetConsSitNF3e, error)
}
