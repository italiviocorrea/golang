package conssitnf3e

import (
	dtos2 "github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/models/dtos"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/models/vos"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/ports"
	"strconv"
	"time"
)

type rnI04Rej236 struct {
	ConsSitNF3e dtos2.ConsSitNF3e
}

func NewRnI04Rej236(consSitNF3e dtos2.ConsSitNF3e) ports.Supplier {
	return &rnI02Rej226{ConsSitNF3e: consSitNF3e}
}

func (e *rnI04Rej236) Validate() dtos2.ResultadoProcessamento {
	var resp = dtos2.ResultadoProcessamento{
		CStat:       "100",
		XMotivo:     "Ok",
		Complemento: "",
	}
	ano, _ := strconv.Atoi("20" + e.ConsSitNF3e.ChNF3eDecode.Aamm[0:2])
	data := time.Now()
	anoCorrente := data.Year()
	mes, _ := strconv.Atoi(e.ConsSitNF3e.ChNF3eDecode.Aamm[2:4])
	var cnpj = vos.Cnpj{Value: e.ConsSitNF3e.ChNF3eDecode.Cnpj}

	if !cnpj.IsValid() {
		resp.XMotivo = "Rejeição: Chave de Acesso inválida [Motivo: CNPJ zerado ou inválido]"
	} else if ano < 2019 || ano > anoCorrente {
		resp.XMotivo = "Rejeição: Chave de Acesso inválida [Motivo: Ano < 2019 ou maior que o atual]"
	} else if mes < 1 || mes > 12 {
		resp.XMotivo = "Rejeição: Chave de Acesso inválida [Motivo: Mês inválido (0 ou > 12)]"
	} else if e.ConsSitNF3e.ChNF3eDecode.Mod != "66" {
		resp.XMotivo = "Rejeição: Chave de Acesso inválida [Motivo: Modelo diferente de 66, Número zerado]"
	} else if e.ConsSitNF3e.ChNF3eDecode.Tpemis != "1" && e.ConsSitNF3e.ChNF3eDecode.Tpemis != "2" {
		resp.XMotivo = "Rejeição: Chave de Acesso inválida [Motivo: Tipo de emissão inválido]"
	} else if e.ConsSitNF3e.ChNF3eDecode.Cuf != "50" {
		resp.XMotivo = "Rejeição: Chave de Acesso inválida [Motivo: UF inválida]"
	} else if !e.ConsSitNF3e.ChNF3eDecode.IsValidDv() {
		resp.XMotivo = "Rejeição: Chave de Acesso inválida [Motivo: DV inválido]"
	}
	if resp.XMotivo != "Ok" {
		resp.CStat = "236"
	}

	return resp
}
