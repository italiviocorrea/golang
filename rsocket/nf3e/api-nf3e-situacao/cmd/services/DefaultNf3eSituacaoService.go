package services

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/models/dtos"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/models/entities"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/repositories/nf3e_qry"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/rules/conssitnf3e"
	"log"
)

/*
	Implementação de serviço, pode haver muitas implementações do mesmo serviço
	Observe que aqui o Nf3eSituacaoRepository é adicionado como dependência
*/
type defaultNf3eSituacaoService struct {
	Nf3eSituacaoRepository nf3e_qry.Nf3eSituacaoRepositoryInterface
}

func NewNf3eSituacaoService(nf3eSituacaoRepository nf3e_qry.Nf3eSituacaoRepositoryInterface) Nf3eSituacaoService {
	return &defaultNf3eSituacaoService{Nf3eSituacaoRepository: nf3eSituacaoRepository}
}

/*
	Implementar todas as validações
*/
func (repo *defaultNf3eSituacaoService) GetNf3eSituacao(consSitNF3e dtos.ConsSitNF3e) (entities.Nf3eSituacao, error) {
	//consSitNF3e.Validar()
	regrasSupervisor := conssitnf3e.NewRegrasSupervisor(consSitNF3e)
	resp := regrasSupervisor.Validate()

	log.Println(resp)

	return repo.Nf3eSituacaoRepository.FindByID(consSitNF3e.ChNF3e)
}
