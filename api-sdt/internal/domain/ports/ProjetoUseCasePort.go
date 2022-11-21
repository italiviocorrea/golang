package ports

import (
	"api-sdt/internal/domain/dtos"
	"api-sdt/internal/domain/entities"
)

type ProjetoUseCasePort interface {
	Create(projeto *entities.Projeto) (*entities.Projeto, *dtos.Error)
	FindByName(nome string) (*entities.Projeto, *dtos.Error)
	FindAll() ([]*entities.Projeto, *dtos.Error)
	Update(nome string, projeto *entities.Projeto) (*entities.Projeto, *dtos.Error)
	Delete(nome string) *dtos.Error
}
