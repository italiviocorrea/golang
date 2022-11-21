package ports

import "api-sdt/internal/domain/entities"

type ProjetoRepositoryPort interface {
	Create(projeto *entities.Projeto) (*entities.Projeto, error)
	FindByName(nome string) (*entities.Projeto, error)
	FindAll() ([]*entities.Projeto, error)
	Update(nome string, projeto *entities.Projeto) (*entities.Projeto, error)
	Delete(nome string) error
}
