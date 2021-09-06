package conssitnf3e

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/entities"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/ports"
)

type ruleI05Rej482 struct {
	ChaveAcesso entities.ChaveAcesso
}

func NewRnI05Rej482(chaveAcesso entities.ChaveAcesso) ports.RulePort {
	return &ruleI05Rej482{ChaveAcesso: chaveAcesso}
}

func (e *ruleI05Rej482) Validate() entities.ResultadoProcessamento {
	var resp = entities.ResultadoProcessamento{
		CStat:       "100",
		XMotivo:     "Ok",
		Complemento: "ruleI05Rej482",
	}

	if e.ChaveAcesso.Nsiteautoriz != "0" {
		resp.CStat = "482"
		resp.XMotivo = "Rejeição: Site de autorização inválido"
	}

	return resp
}
