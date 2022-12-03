package usecases

import (
	"api-sdt/internal/app/config"
	"api-sdt/internal/domain/dtos"
	"api-sdt/internal/domain/entities"
	"api-sdt/internal/domain/ports"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjetoUseCase struct {
	projetoRepository ports.ProjetoRepositoryPort
	cfg               *config.Settings
}

func New(cfg *config.Settings, projetoRepository ports.ProjetoRepositoryPort) ports.ProjetoUseCasePort {
	return &ProjetoUseCase{
		projetoRepository: projetoRepository,
		cfg:               cfg,
	}
}

func (p ProjetoUseCase) Create(projeto *entities.Projeto) (*entities.Projeto, *dtos.Error) {

	projetExist, err := p.projetoRepository.FindByName(projeto.Nome)

	if projetExist != nil {
		return nil, &dtos.Error{
			Message: "Projeto já foi criado no banco de dados",
			Code:    400,
			Name:    "PROJETO_EXIST",
			Error:   err,
		}
	}

	log.Info("Gravando o projeto no banco de dados")

	projeto.ID = primitive.NewObjectID()
	_, err = p.projetoRepository.Create(projeto)

	if err != nil {
		return nil, &dtos.Error{
			Message: "Não foi possível criar o projeto",
			Code:    500,
			Name:    "SERVER_ERROR",
			Error:   err,
		}
	}
	return projeto, nil
}

func (p ProjetoUseCase) FindByName(nome string) (*entities.Projeto, *dtos.Error) {

	projetExist, err := p.projetoRepository.FindByName(nome)

	if projetExist == nil {
		return nil, &dtos.Error{
			Message: "Projeto não encontrado no banco de dados",
			Code:    404,
			Name:    "PROJETO_NOT_FOUD",
		}
	}

	if err != nil {
		return nil, &dtos.Error{
			Message: "Não foi possível pesquisar no banco de dados",
			Code:    500,
			Name:    "SERVER_ERROR",
			Error:   err,
		}
	}

	return projetExist, nil

}

func (p ProjetoUseCase) FindAll() ([]*entities.Projeto, *dtos.Error) {

	projetos, err := p.projetoRepository.FindAll()

	if err != nil {
		return nil, &dtos.Error{
			Message: "Não foi possível pesquisar no banco de dados",
			Code:    500,
			Name:    "SERVER_ERROR",
			Error:   err,
		}
	}

	if len(projetos) == 0 {
		return nil, &dtos.Error{
			Message: "Não foi encontrado nenhum projeto no banco de dados",
			Code:    404,
			Name:    "PROJETO_NOT_FOUD",
		}
	}
	log.Info("Tem proejtos")
	return projetos, nil
}

func (p ProjetoUseCase) Update(nome string, projeto *entities.Projeto) (*entities.Projeto, *dtos.Error) {
	projetoUpdate, err := p.projetoRepository.Update(nome, projeto)

	if err != nil {
		return nil, &dtos.Error{
			Message: "Não foi possível atualizar este projeto",
			Code:    400,
			Name:    "PROJETO_NOUPDATE",
			Error:   err,
		}
	}
	log.Info("Projeto atualizado com sucesso!")
	return projetoUpdate, nil
}

func (p ProjetoUseCase) Delete(nome string) *dtos.Error {

	_, err1 := p.FindByName(nome)

	if err1 != nil {
		return &dtos.Error{
			Message: "Não foi encontrado este projeto",
			Code:    404,
			Name:    "PROJETO_NOT_FOUND",
			Error:   err1.Error,
		}
	}

	err := p.projetoRepository.Delete(nome)

	if err != nil {
		return &dtos.Error{
			Message: "Não foi excluir este projeto",
			Code:    400,
			Name:    "PROJETO_NODELETE",
			Error:   err,
		}
	}
	return nil
}
