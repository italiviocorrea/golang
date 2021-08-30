package services

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/configs"
	dtos2 "github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/models/dtos"
	ports2 "github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/ports"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/services/rules/conssitnf3e"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/utils"
)

/*
	Implementação de serviço, pode haver muitas implementações do mesmo serviço
	Observe que aqui o Nf3eSituacaoRepository é adicionado como dependência
*/
type defaultNf3eSituacaoService struct {
	Nf3eSituacaoRepository ports2.Nf3eSituacaoRepositoryInterface
}

func NewNf3eSituacaoService(nf3eSituacaoRepository ports2.Nf3eSituacaoRepositoryInterface) ports2.Nf3eSituacaoService {
	return &defaultNf3eSituacaoService{Nf3eSituacaoRepository: nf3eSituacaoRepository}
}

/*
	Implementar todas as validações
*/
func (repo *defaultNf3eSituacaoService) GetNf3eSituacao(consSitNF3e dtos2.ConsSitNF3e) (dtos2.RetConsSitNF3e, error) {
	//consSitNF3e.Validar()
	regrasSupervisor := conssitnf3e.NewRegrasSupervisor(consSitNF3e)
	resps := regrasSupervisor.Validate()

	if utils.IsRejects(resps) {
		resp := utils.Take(utils.FilterRejects(resps), 1)[0]
		return dtos2.RetConsSitNF3e{
			Versao:         configs.Get().VersaoLeiaute,
			Xmlns:          configs.Get().Xmlns,
			TpAmb:          configs.Get().TpAmb,
			VerAplic:       configs.Get().VerAplic,
			Cstat:          resp.CStat,
			Xmotivo:        resp.XMotivo,
			Cuf:            configs.Get().CUF,
			Protnf3e:       "",
			Proceventonf3e: nil,
		}, nil
	} else {
		nf3e, err := repo.Nf3eSituacaoRepository.FindByID(consSitNF3e.ChNF3e)
		if err != nil {
			return dtos2.RetConsSitNF3e{
				Versao:         configs.Get().VersaoLeiaute,
				Xmlns:          configs.Get().Xmlns,
				TpAmb:          configs.Get().TpAmb,
				VerAplic:       configs.Get().VerAplic,
				Cstat:          "999",
				Xmotivo:        "Rejeição: Erro não catalogado",
				Cuf:            configs.Get().CUF,
				Protnf3e:       "",
				Proceventonf3e: nil,
			}, nil

		}
		return dtos2.RetConsSitNF3e{
			Versao:         configs.Get().VersaoLeiaute,
			Xmlns:          configs.Get().Xmlns,
			TpAmb:          nf3e.Tpamb,
			VerAplic:       configs.Get().VerAplic,
			Cstat:          nf3e.Cstat,
			Xmotivo:        nf3e.Xmotivo,
			Cuf:            nf3e.Cuf,
			Protnf3e:       nf3e.Protnf3e,
			Proceventonf3e: nf3e.Proceventonf3e,
		}, err
	}
}
