package conssitnf3e

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/models/dtos"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/rules/interfaces"
	"log"
)

type rnI05Rej482 struct {
	ConsSitNF3e dtos.ConsSitNF3e
}

func NewRnI05Rej482(consSitNF3e dtos.ConsSitNF3e) interfaces.Supplier {
	return &rnI02Rej226{ConsSitNF3e: consSitNF3e}
}

func (e *rnI05Rej482) Validate() dtos.RespostaValidacao {
	var resp = dtos.RespostaValidacao{
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
