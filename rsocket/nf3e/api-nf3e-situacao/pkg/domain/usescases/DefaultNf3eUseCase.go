package usescases

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/configs"
	entities2 "github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/entities"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/ports"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/usescases/rules/conssitnf3e"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/utils"
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
func (repo *defaultNf3eSituacaoUseCase) GetNf3eSituacao(consSitNF3e entities2.ConsSitNF3e) (entities2.Nf3eSituacao, error) {
	//consSitNF3e.Validar()
	regrasSupervisor := conssitnf3e.NewRuleSupervisor(consSitNF3e)
	resps := regrasSupervisor.Validate()

	if utils.IsRejects(resps) {
		resp := utils.Take(utils.FilterRejects(resps), 1)[0]
		return entities2.Nf3eSituacao{
			Versao:         configs.Get().VersaoLeiaute,
			TpAmb:          configs.Get().TpAmb,
			Cstat:          resp.CStat,
			Xmotivo:        resp.XMotivo,
			Cuf:            configs.Get().CUF,
			Protnf3e:       "",
			Proceventonf3e: nil,
		}, nil
	} else {
		nf3e, err := repo.Nf3eSituacaoRepository.FindByID(consSitNF3e.ChNF3e)
		if err != nil {
			return entities2.Nf3eSituacao{
				Versao:         configs.Get().VersaoLeiaute,
				Cstat:          "999",
				Xmotivo:        "Rejeição: Erro não catalogado",
				Cuf:            configs.Get().CUF,
				Protnf3e:       "",
				Proceventonf3e: nil,
			}, nil

		}
		return nf3e, err
	}
}
