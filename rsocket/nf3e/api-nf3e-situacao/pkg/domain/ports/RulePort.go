package ports

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/models/entities"
)

type RulePort interface {
	Validate() entities.ResultadoProcessamento
}
