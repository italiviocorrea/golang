package usescases

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/configs"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/entities"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/ports"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/usescases/rules/conssitnf3e"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/utils"
	"github.com/rs/zerolog/log"
)

/*
	Implementação de serviço, pode haver muitas implementações do mesmo serviço
	Observe que aqui o Nf3eSituacaoRepository é adicionado como dependência
*/
type defaultNf3eSituacaoUseCase struct {
	Nf3eSituacaoRepository ports.Nf3eSituacaoRepositoryPort
}

func NewNf3eSituacaoUseCase(nf3eSituacaoRepository ports.Nf3eSituacaoRepositoryPort) ports.Nf3eSituacaoUseCasePort {
	return &defaultNf3eSituacaoUseCase{Nf3eSituacaoRepository: nf3eSituacaoRepository}
}

/*
	Implementar todas as validações
*/
func (repo *defaultNf3eSituacaoUseCase) GetNf3eSituacao(consSitNF3e entities.ConsSitNF3e) (entities.Nf3eSituacao, error) {
	//consSitNF3e.Validar()
	ruleSupervisor := conssitnf3e.NewRuleSupervisor(consSitNF3e)
	reps := ruleSupervisor.Validate()

	log.Info().
		Str("service", "api-nf3e-situacao").
		Str("module", "DefaultNf3eUseCase").
		Str("chNF3e", consSitNF3e.ChNF3e).
		Msgf("{Todas Respostas:%s}\n", utils.JsonMarshal(reps))

	if utils.IsRejects(reps) {
		// Caso haja rejeições pegar a primeira
		resp := utils.Take(utils.FilterRejects(reps), 1)[0]
		// Retorna a primeira rejeição
		return nf3eSituacaoRejeitada(resp.CStat, resp.XMotivo), nil
	} else {
		// Caso não haja rejeições buscar a NF3eSituação
		nf3e, err := repo.Nf3eSituacaoRepository.FindByID(consSitNF3e.ChNF3e)
		if err != nil {
			return nf3eSituacaoRejeitada("999", "Rejeição: Erro não catalogado"), nil
		}
		return nf3e, err
	}
}

func nf3eSituacaoRejeitada(cStat string, xMotivo string) entities.Nf3eSituacao {
	return entities.Nf3eSituacao{
		Versao:         configs.Get().VersaoLeiaute,
		Cstat:          cStat,
		Xmotivo:        xMotivo,
		Cuf:            configs.Get().CUF,
		Protnf3e:       "",
		Proceventonf3e: nil,
	}
}
