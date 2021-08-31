package ports

import (
	entities2 "github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/entities"
)

type Nf3eSituacaoUseCasePort interface {
	GetNf3eSituacao(consSitNF3e entities2.ConsSitNF3e) (entities2.Nf3eSituacao, error)
}
