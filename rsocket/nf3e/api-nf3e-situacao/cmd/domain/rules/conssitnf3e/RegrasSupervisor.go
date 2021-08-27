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

	// Cria os canais para receber as respostas
	ch252 := make(chan dtos.RespostaValidacao)
	ch226 := make(chan dtos.RespostaValidacao)
	ch478 := make(chan dtos.RespostaValidacao)
	ch236 := make(chan dtos.RespostaValidacao)
	ch482 := make(chan dtos.RespostaValidacao)

	// Chamas as regras de forma concorrente usando GO rotinas
	go func() {
		ch252 <- regraI01Rej252(r.ConsSitNF3e)
		close(ch252)
	}()
	go func() {
		ch226 <- regraI02Rej226(r.ConsSitNF3e)
		close(ch226)
	}()
	go func() {
		ch236 <- regraI04Rej236(r.ConsSitNF3e)
		close(ch236)
	}()
	go func() {
		ch478 <- regraI03Rej478(r.ConsSitNF3e)
		close(ch478)
	}()
	go func() {
		ch482 <- regraI05Rej482(r.ConsSitNF3e)
		close(ch482)
	}()

	// pega as respostas dos canais e anexa ao array de respostas
	resps = append(resps, <-ch252, <-ch226, <-ch236, <-ch478, <-ch482)

	fmt.Println(resps)

	for _, s := range resps {
		if s.CStat != "100" {
			fmt.Println(s)
		}
	}

	return resps
}

func regraI01Rej252(consSitNF3e dtos.ConsSitNF3e) dtos.RespostaValidacao {
	rej := rnI01Rej252{ConsSitNF3e: consSitNF3e}
	resp := rej.Validate()
	return resp
}

func regraI02Rej226(consSitNF3e dtos.ConsSitNF3e) dtos.RespostaValidacao {
	rej := rnI02Rej226{ConsSitNF3e: consSitNF3e}
	resp := rej.Validate()
	return resp
}

func regraI03Rej478(consSitNF3e dtos.ConsSitNF3e) dtos.RespostaValidacao {

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
