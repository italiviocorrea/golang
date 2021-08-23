package repository

import "github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/model"

type Nf3eSituacaoRepository interface {
	Consultar(chnf3e string) (model.Nf3eSituacao, error)
}
