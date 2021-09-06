package conssitnf3e

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/configs"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/entities"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/ports"
)

type ruleI02Rej226 struct {
	ChaveAcesso entities.ChaveAcesso
}

func NewRuleI02Rej226(chaveAcesso entities.ChaveAcesso) ports.RulePort {
	return &ruleI02Rej226{ChaveAcesso: chaveAcesso}
}

func (r *ruleI02Rej226) Validate() entities.ResultadoProcessamento {
	var resp = entities.ResultadoProcessamento{
		CStat:       "100",
		XMotivo:     "Ok",
		Complemento: "ruleI02Rej226",
	}

	if r.ChaveAcesso.Cuf != configs.Get().CUF {
		resp.CStat = "226"
		resp.XMotivo = "Rejeição: Código da UF do Emitente diverge da UF autorizadora"
	}
	return resp
}
