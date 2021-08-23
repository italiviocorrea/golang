package service

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/model"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/repository"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/rule"
)

/*
	Implementação de serviço, pode haver muitas implementações do mesmo serviço
	Observe que aqui o repo é adicionado como dependência
*/
type DefaultNf3eSituacaoService struct {
	repo repository.Nf3eSituacaoRepository
}

/*
	Implementar todas as validações
*/
func (service DefaultNf3eSituacaoService) Consultar(consSitNF3e rule.ConsSitNF3e) (model.Nf3eSituacao, error) {
	consSitNF3e.Validar()
	return service.repo.Consultar(consSitNF3e.ChNF3e)
}

func NewNf3eSituacaoService(repo repository.Nf3eSituacaoRepository) DefaultNf3eSituacaoService {
	return DefaultNf3eSituacaoService{repo}
}
