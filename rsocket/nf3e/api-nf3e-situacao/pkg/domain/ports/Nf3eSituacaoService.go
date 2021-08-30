package ports

import (
	dtos2 "github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/models/dtos"
)

type Nf3eSituacaoService interface {
	GetNf3eSituacao(consSitNF3e dtos2.ConsSitNF3e) (dtos2.RetConsSitNF3e, error)
}
