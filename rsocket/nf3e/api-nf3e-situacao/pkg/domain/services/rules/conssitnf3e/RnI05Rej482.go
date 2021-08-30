package conssitnf3e

import (
	dtos2 "github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/models/dtos"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/ports"
	"log"
)

type rnI05Rej482 struct {
	ConsSitNF3e dtos2.ConsSitNF3e
}

func NewRnI05Rej482(consSitNF3e dtos2.ConsSitNF3e) ports.Supplier {
	return &rnI02Rej226{ConsSitNF3e: consSitNF3e}
}

func (e *rnI05Rej482) Validate() dtos2.ResultadoProcessamento {
	var resp = dtos2.ResultadoProcessamento{
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
