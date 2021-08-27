package interfaces

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/models/dtos"
)

type Supplier interface {
	Validate() dtos.RespostaValidacao
}
