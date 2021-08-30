package conssitnf3e

import (
	"fmt"
	dtos2 "github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/models/dtos"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/ports"
	"strconv"
	"time"
)

type rnI03Rej478 struct {
	ConsSitNF3e dtos2.ConsSitNF3e
}

func NewRnI03Rej478(consSitNF3e dtos2.ConsSitNF3e) ports.Supplier {
	return &rnI02Rej226{ConsSitNF3e: consSitNF3e}
}

func (e *rnI03Rej478) Validate() dtos2.ResultadoProcessamento {
	var resp = dtos2.ResultadoProcessamento{
		CStat:       "100",
		XMotivo:     "Ok",
		Complemento: "",
	}

	anoMesChave, _ := strconv.Atoi("20" + e.ConsSitNF3e.ChNF3eDecode.Aamm)
	anoMes6MesesAntes, _ := strconv.Atoi(getAnoMes6MesesAtras())

	fmt.Printf("%d - %d", anoMesChave, anoMes6MesesAntes)

	if anoMesChave < anoMes6MesesAntes {
		resp.CStat = "478"
		resp.XMotivo = "Rejeicao: Consulta a uma Chave de Acesso muito antiga"
	}

	return resp
}

/*
	Calcula a partir da data e hora atual, o ano e mês à 180 dias atrás.
*/
func getAnoMes6MesesAtras() string {
	data := time.Now()
	oneMonthLater := data.AddDate(0, -6, 0)
	s := fmt.Sprintf("%04d%02d", oneMonthLater.Year(), int(oneMonthLater.Month()))
	return s
}
