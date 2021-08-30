package conssitnf3e

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/models/entities"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/ports"
	"log"
)

type rnI05Rej482Rule struct {
	ConsSitNF3e entities.ConsSitNF3e
}

func NewRnI05Rej482(consSitNF3e entities.ConsSitNF3e) ports.RulePort {
	return &rnI02Rej226Rule{ConsSitNF3e: consSitNF3e}
}

func (e *rnI05Rej482Rule) Validate() entities.ResultadoProcessamento {
	var resp = entities.ResultadoProcessamento{
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
