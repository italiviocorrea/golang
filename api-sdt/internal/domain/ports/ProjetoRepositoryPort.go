package ports

import (
	"api-sdt/internal/domain/entities"
	"context"
)

type ProjetoRepositoryPort interface {
	Create(ctx context.Context, projeto *entities.Projeto) (*entities.Projeto, error)
	FindByName(ctx context.Context, nome string) (*entities.Projeto, error)
	FindAll(ctx context.Context) ([]*entities.Projeto, error)
	Update(ctx context.Context, nome string, projeto *entities.Projeto) (*entities.Projeto, error)
	Delete(ctx context.Context, nome string) error
}
