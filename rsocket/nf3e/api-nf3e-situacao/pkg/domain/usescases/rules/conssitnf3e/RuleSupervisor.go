package conssitnf3e

import (
	entities2 "github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/entities"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/entities/vos"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/ports"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/utils"
	"github.com/rs/zerolog/log"
)

type ruleSupervisor struct {
	ConsSitNF3e entities2.ConsSitNF3e
}

func NewRuleSupervisor(consSitNF3e entities2.ConsSitNF3e) ports.RuleSupervisorPort {
	return &ruleSupervisor{ConsSitNF3e: consSitNF3e}
}

func (r *ruleSupervisor) Validate() []entities2.ResultadoProcessamento {
	var validationResponses []entities2.ResultadoProcessamento

	// decodificar a chave de acesso
	chave := vos.ChaveAcesso{}
	chave.Decode(r.ConsSitNF3e.ChNF3e)
	r.ConsSitNF3e.ChNF3eDecode = chave

	log.Info().
		Str("service", "api-nf3e-situacao").
		Str("module", "conssitnf3e.RuleSupervisor").
		Str("chNF3e", r.ConsSitNF3e.ChNF3e).
		Msg(utils.JsonMarshal(r))

	// Criar os canais para receber as respostas
	ch252 := make(chan entities2.ResultadoProcessamento)
	ch226 := make(chan entities2.ResultadoProcessamento)
	ch478 := make(chan entities2.ResultadoProcessamento)
	ch236 := make(chan entities2.ResultadoProcessamento)
	ch482 := make(chan entities2.ResultadoProcessamento)

	// Executar as regras de forma concorrente usando GO rotinas
	go func() {
		defer close(ch252)
		rn := ruleI01Rej252{ConsSitNF3e: r.ConsSitNF3e}
		ch252 <- rn.Validate()
	}()
	go func() {
		defer close(ch226)
		rn := ruleI02Rej226{ConsSitNF3e: r.ConsSitNF3e}
		ch226 <- rn.Validate()
	}()
	go func() {
		defer close(ch236)
		rn := ruleI04Rej236{ConsSitNF3e: r.ConsSitNF3e}
		ch236 <- rn.Validate()
	}()
	go func() {
		defer close(ch478)
		rn := ruleI03Rej478{ConsSitNF3e: r.ConsSitNF3e}
		ch478 <- rn.Validate()
	}()
	go func() {
		defer close(ch482)
		rn := ruleI05Rej482{ConsSitNF3e: r.ConsSitNF3e}
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
		Str("component", "conssitnf3e.RuleSupervisor").
		Str("chNF3e", r.ConsSitNF3e.ChNF3e).
		Msgf("{ResultadoProcessamento:%s}", utils.JsonMarshal(validationResponses))

	return validationResponses
}
