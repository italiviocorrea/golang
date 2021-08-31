package conssitnf3e

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/entities"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/entities/vos"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/ports"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/utils"
)

type ruleSupervisor struct {
	ConsSitNF3e entities.ConsSitNF3e
}

func NewRuleSupervisor(consSitNF3e entities.ConsSitNF3e) ports.RuleSupervisorPort {
	return &ruleSupervisor{ConsSitNF3e: consSitNF3e}
}

func (r *ruleSupervisor) Validate() []entities.ResultadoProcessamento {
	var validationResponses []entities.ResultadoProcessamento

	// decodificar a chave de acesso
	chave := vos.ChaveAcesso{}
	chave.Decode(r.ConsSitNF3e.ChNF3e)
	r.ConsSitNF3e.ChNF3eDecode = chave

	// Criar os canais para receber as respostas
	qtRules := 5                                                 // quantidade rules a serem executadas de forma concorrente
	chRsp := make(chan entities.ResultadoProcessamento, qtRules) // cria um canal com buffer para 5 regras
	defer close(chRsp)

	// Executar as regras de forma concorrente usando GO rotinas
	go func() {
		rn := ruleI01Rej252{ConsSitNF3e: r.ConsSitNF3e}
		chRsp <- rn.Validate()
	}()
	go func() {
		rn := ruleI02Rej226{ConsSitNF3e: r.ConsSitNF3e}
		chRsp <- rn.Validate()
	}()
	go func() {
		rn := ruleI04Rej236{ConsSitNF3e: r.ConsSitNF3e}
		chRsp <- rn.Validate()
	}()
	go func() {
		rn := ruleI03Rej478{ConsSitNF3e: r.ConsSitNF3e}
		chRsp <- rn.Validate()
	}()
	go func() {
		rn := ruleI05Rej482{ConsSitNF3e: r.ConsSitNF3e}
		chRsp <- rn.Validate()
	}()

	// pegar todas as respostas que estão no canal, e atribui ao array de resultado processamento
	validationResponses = utils.TakeChan(chRsp, qtRules)

	return validationResponses
}
