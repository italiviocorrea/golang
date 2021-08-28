package conssitnf3e

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/models/dtos"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/models/vos"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/rules/interfaces"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/utils"
	"github.com/rs/zerolog/log"
)

type regrasSupervisor struct {
	ConsSitNF3e dtos.ConsSitNF3e
}

func NewRegrasSupervisor(consSitNF3e dtos.ConsSitNF3e) interfaces.Supervisor {
	return &regrasSupervisor{ConsSitNF3e: consSitNF3e}
}

func (r *regrasSupervisor) Validate() []dtos.RespostaValidacao {
	var validationResponses []dtos.RespostaValidacao

	// decodificar a chave de acesso
	chave := vos.ChaveAcesso{}
	chave.Decode(r.ConsSitNF3e.ChNF3e)
	r.ConsSitNF3e.ChNF3eDecode = chave

	log.Info().
		Str("service", "api-nf3e-situacao").
		Str("module", "RegrasSupervisor").
		Str("chNF3e", r.ConsSitNF3e.ChNF3e).
		Msg(utils.JsonMarshal(r))

	// Criar os canais para receber as respostas
	ch252 := make(chan dtos.RespostaValidacao)
	ch226 := make(chan dtos.RespostaValidacao)
	ch478 := make(chan dtos.RespostaValidacao)
	ch236 := make(chan dtos.RespostaValidacao)
	ch482 := make(chan dtos.RespostaValidacao)

	// Executar as regras de forma concorrente usando GO rotinas
	go func() {
		defer close(ch252)
		rn := rnI01Rej252{ConsSitNF3e: r.ConsSitNF3e}
		ch252 <- rn.Validate()
	}()
	go func() {
		defer close(ch226)
		rn := rnI02Rej226{ConsSitNF3e: r.ConsSitNF3e}
		ch226 <- rn.Validate()
	}()
	go func() {
		defer close(ch236)
		rn := rnI04Rej236{ConsSitNF3e: r.ConsSitNF3e}
		ch236 <- rn.Validate()
	}()
	go func() {
		defer close(ch478)
		rn := rnI03Rej478{ConsSitNF3e: r.ConsSitNF3e}
		ch478 <- rn.Validate()
	}()
	go func() {
		defer close(ch482)
		rn := rnI05Rej482{ConsSitNF3e: r.ConsSitNF3e}
		ch482 <- rn.Validate()
	}()

	// pegar as respostas dos canais e anexa ao array de respostas
	validationResponses = append(validationResponses,
		utils.Reduce(ch252),
		utils.Reduce(ch226),
		utils.Reduce(ch236),
		utils.Reduce(ch478),
		utils.Reduce(ch482))

	log.Info().
		Str("service", "api-nf3e-situacao").
		Str("component", "RegrasSupervisor").
		Str("chNF3e", r.ConsSitNF3e.ChNF3e).
		Msgf("{RespostaValidacao:%s}", utils.JsonMarshal(validationResponses))

	return validationResponses
}
