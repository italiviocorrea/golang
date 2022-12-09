package ports

import (
	"api-sdt/internal/domain/dtos"
	"api-sdt/internal/domain/entities"
	"context"
)

type ProjetoUseCasePort interface {
	Create(ctx context.Context, projeto *entities.Projeto) (*entities.Projeto, *dtos.Error)
	FindByName(ctx context.Context, nome string) (*entities.Projeto, *dtos.Error)
	FindAll(ctx context.Context) ([]*entities.Projeto, *dtos.Error)
	Update(ctx context.Context, nome string, projeto *entities.Projeto) (*entities.Projeto, *dtos.Error)
	Delete(ctx context.Context, nome string) *dtos.Error
}
