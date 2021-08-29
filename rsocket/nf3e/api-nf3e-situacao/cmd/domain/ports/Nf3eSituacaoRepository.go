package ports

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/models/entities"
)

type Nf3eSituacaoRepositoryInterface interface {
	FindByID(chnf3e string) (entities.Nf3eSituacao, error)
}
