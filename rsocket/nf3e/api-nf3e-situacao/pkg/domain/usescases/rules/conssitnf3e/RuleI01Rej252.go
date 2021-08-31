package conssitnf3e

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/configs"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/entities"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/ports"
)

type ruleI01Rej252 struct {
	ConsSitNF3e entities.ConsSitNF3e
}

func NewRuleI01Rej252(consSitNF3e entities.ConsSitNF3e) ports.RulePort {
	return &ruleI01Rej252{ConsSitNF3e: consSitNF3e}
}

func (r *ruleI01Rej252) Validate() entities.ResultadoProcessamento {
	var resp = entities.ResultadoProcessamento{
		CStat:       "100",
		XMotivo:     "Ok",
		Complemento: "ruleI01Rej252",
	}

	if r.ConsSitNF3e.TpAmb != configs.Get().TpAmb {
		resp.CStat = "252"
		resp.XMotivo = "Rejeicao: Ambiente informado diverge do Ambiente de recebimento"
	}
	return resp
}
