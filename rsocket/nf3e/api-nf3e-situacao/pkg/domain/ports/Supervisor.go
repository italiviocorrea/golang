package ports

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/models/dtos"
)

type Supervisor interface {
	Validate() []dtos.ResultadoProcessamento
}
