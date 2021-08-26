package rules

import "github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/models/dtos"

type regrasSupervisor struct {
	ConsSitNF3e dtos.ConsSitNF3e
}

func NewRegrasSupervisor(consSitNF3e dtos.ConsSitNF3e) Supervisor {
	return &regrasSupervisor{ConsSitNF3e: consSitNF3e}
}

func (r *regrasSupervisor) Validate() []dtos.RespostaValidacao {
	var resps = []dtos.RespostaValidacao{}

	rej252 := rnI01Rej252{ConsSitNF3e: r.ConsSitNF3e}
	resps = append(resps, rej252.Validate())

	return resps
}
