package conssitnf3e

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/configs"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/models/entities"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/ports"
)

type rnI02Rej226 struct {
	ConsSitNF3e entities.ConsSitNF3e
}

func NewRnI02Rej226(consSitNF3e entities.ConsSitNF3e) ports.SupplierPort {
	return &rnI02Rej226{ConsSitNF3e: consSitNF3e}
}

func (r *rnI02Rej226) Validate() entities.ResultadoProcessamento {
	var resp = entities.ResultadoProcessamento{
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
