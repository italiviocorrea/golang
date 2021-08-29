package conssitnf3e

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/adapters/configs"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/models/dtos"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/ports"
)

type rnI02Rej226 struct {
	ConsSitNF3e dtos.ConsSitNF3e
}

func NewRnI02Rej226(consSitNF3e dtos.ConsSitNF3e) ports.Supplier {
	return &rnI02Rej226{ConsSitNF3e: consSitNF3e}
}

func (r *rnI02Rej226) Validate() dtos.RespostaValidacao {
	var resp = dtos.RespostaValidacao{
		CStat:       "100",
		XMotivo:     "Ok",
		Complemento: "",
	}

	if r.ConsSitNF3e.ChNF3eDecode.Cuf != configs.Get().CUF {
		resp.CStat = "226"
		resp.XMotivo = "Rejeição: Código da UF do Emitente diverge da UF autorizadora"
	}
	return resp
}
