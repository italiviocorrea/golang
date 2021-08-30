package conssitnf3e

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/configs"
	dtos2 "github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/models/dtos"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/ports"
)

type rnI01Rej252 struct {
	ConsSitNF3e dtos2.ConsSitNF3e
}

func NewRnI01Rej252(consSitNF3e dtos2.ConsSitNF3e) ports.Supplier {
	return &rnI01Rej252{ConsSitNF3e: consSitNF3e}
}

func (r *rnI01Rej252) Validate() dtos2.ResultadoProcessamento {
	var resp = dtos2.ResultadoProcessamento{
		CStat:       "100",
		XMotivo:     "Ok",
		Complemento: "",
	}

	if r.ConsSitNF3e.TpAmb != configs.Get().TpAmb {
		resp.CStat = "252"
		resp.XMotivo = "Rejeicao: Ambiente informado diverge do Ambiente de recebimento"
	}
	return resp
}
