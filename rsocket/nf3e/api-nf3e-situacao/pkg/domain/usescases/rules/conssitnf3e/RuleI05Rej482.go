package conssitnf3e

import (
	entities2 "github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/entities"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/ports"
	"log"
)

type ruleI05Rej482 struct {
	ConsSitNF3e entities2.ConsSitNF3e
}

func NewRnI05Rej482(consSitNF3e entities2.ConsSitNF3e) ports.RulePort {
	return &ruleI02Rej226{ConsSitNF3e: consSitNF3e}
}

func (e *ruleI05Rej482) Validate() entities2.ResultadoProcessamento {
	var resp = entities2.ResultadoProcessamento{
		CStat:       "100",
		XMotivo:     "Ok",
		Complemento: "",
	}

	log.Println(e.ConsSitNF3e.ChNF3eDecode.Nsiteautoriz)

	if e.ConsSitNF3e.ChNF3eDecode.Nsiteautoriz != "0" {
		resp.CStat = "482"
		resp.XMotivo = "Rejeição: Site de autorização inválido"
	}

	return resp
}
