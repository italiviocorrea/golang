package conssitnf3e

import (
	"fmt"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/models/dtos"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/rules/interfaces"
)

type regrasSupervisor struct {
	ConsSitNF3e dtos.ConsSitNF3e
}

func NewRegrasSupervisor(consSitNF3e dtos.ConsSitNF3e) interfaces.Supervisor {
	return &regrasSupervisor{ConsSitNF3e: consSitNF3e}
}

func (r *regrasSupervisor) Validate() []dtos.RespostaValidacao {
	var resps = []dtos.RespostaValidacao{}

	// decodificar a chave de acesso
	chave := dtos.ChaveAcesso{}
	chave.Decode(r.ConsSitNF3e.ChNF3e)
	r.ConsSitNF3e.ChNF3eDecode = chave
	fmt.Println(r)

	resp252 := regraI01Rej252(r.ConsSitNF3e)
	resp226 := regraI02Rej226(r.ConsSitNF3e)
	resp478 := regraI03Rej2478(r.ConsSitNF3e)
	resp236 := regraI04Rej236(r.ConsSitNF3e)
	resp482 := regraI05Rej482(r.ConsSitNF3e)

	resps = append(resps, resp252, resp226, resp478, resp236, resp482)

	for _, s := range resps {
		if s.CStat != "100" {
			fmt.Println(s)
		}
	}

	return resps
}

func regraI01Rej252(consSitNF3e dtos.ConsSitNF3e) dtos.RespostaValidacao {

	rej252 := rnI01Rej252{ConsSitNF3e: consSitNF3e}
	resp := rej252.Validate()

	return resp
}

func regraI02Rej226(consSitNF3e dtos.ConsSitNF3e) dtos.RespostaValidacao {

	rej := rnI02Rej226{ConsSitNF3e: consSitNF3e}
	resp := rej.Validate()

	return resp
}

func regraI03Rej2478(consSitNF3e dtos.ConsSitNF3e) dtos.RespostaValidacao {

	rej := rnI03Rej478{ConsSitNF3e: consSitNF3e}
	resp := rej.Validate()

	return resp
}

func regraI04Rej236(consSitNF3e dtos.ConsSitNF3e) dtos.RespostaValidacao {

	rej := rnI04Rej236{ConsSitNF3e: consSitNF3e}
	resp := rej.Validate()

	return resp
}

func regraI05Rej482(consSitNF3e dtos.ConsSitNF3e) dtos.RespostaValidacao {

	rej := rnI05Rej482{ConsSitNF3e: consSitNF3e}
	resp := rej.Validate()

	return resp
}
