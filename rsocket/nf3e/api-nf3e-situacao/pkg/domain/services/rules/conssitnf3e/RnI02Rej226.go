package conssitnf3e

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/configs"
	dtos2 "github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/models/dtos"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/ports"
)

type rnI02Rej226 struct {
	ConsSitNF3e dtos2.ConsSitNF3e
}

func NewRnI02Rej226(consSitNF3e dtos2.ConsSitNF3e) ports.Supplier {
	return &rnI02Rej226{ConsSitNF3e: consSitNF3e}
}

func (r *rnI02Rej226) Validate() dtos2.ResultadoProcessamento {
	var resp = dtos2.ResultadoProcessamento{
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
