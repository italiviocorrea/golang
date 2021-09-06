package conssitnf3e

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/entities"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/ports"
)

type ruleI04Rej236 struct {
	ChaveAcesso entities.ChaveAcesso
}

func NewRuleI04Rej236(chaveAcesso entities.ChaveAcesso) ports.RulePort {
	return &ruleI04Rej236{ChaveAcesso: chaveAcesso}
}

func (e *ruleI04Rej236) Validate() entities.ResultadoProcessamento {

	return e.ChaveAcesso.Validate()
}
