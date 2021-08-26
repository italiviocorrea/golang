package rules

import "github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/models/dtos"

type Supervisor interface {
	Validate() []dtos.RespostaValidacao
}
