package mongodb

import (
	"api-sdt/internal/app/config"
	"api-sdt/internal/app/trace"
	"api-sdt/internal/domain/entities"
	"api-sdt/internal/domain/ports"
	"context"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjetoRepository struct {
	projetoColecao *mongo.Collection
	ctx            context.Context
}

func NewRepository(cfg *config.Settings, mongo *mongo.Client) ports.ProjetoRepositoryPort {
	_, span := trace.NewSpan(context.TODO(), "ProjetoRepository.NewRepository")
	defer span.End()

	projetoColecao := mongo.Database(cfg.DbName).Collection("projetos")
	return &ProjetoRepository{
		projetoColecao: projetoColecao,
		ctx:            context.TODO(),
	}
}

func (p ProjetoRepository) Create(ctx context.Context, projeto *entities.Projeto) (*entities.Projeto, error) {
	_, span := trace.NewSpan(ctx, "ProjetoRepository.Create")
	defer span.End()

	_, err := p.projetoColecao.InsertOne(ctx, projeto)
	if err != nil {
		trace.AddSpanError(span, err)
		return nil, errors.Wrap(err, "Erro ao inserir projeto")
	}
	return projeto, nil
}

func (p ProjetoRepository) FindByName(ctx context.Context, nome string) (*entities.Projeto, error) {

	_, span := trace.NewSpan(ctx, "ProjetoRepository.FindByName")
	defer span.End()

	var prjFound *entities.Projeto

	filter := bson.D{primitive.E{Key: "nome", Value: nome}}

	err := p.projetoColecao.FindOne(ctx, filter).Decode(&prjFound)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			trace.AddSpanError(span, err)
			return nil, errors.Wrap(err, "Não existe projeto com este nome!")
		}
		trace.AddSpanError(span, err)
		return nil, errors.Wrap(err, "Erro ao pesquisar projeto por nome")
	}
	return prjFound, nil
}

func (p ProjetoRepository) FindAll(ctx context.Context) ([]*entities.Projeto, error) {
	_, span := trace.NewSpan(ctx, "ProjetoRepository.FindAll")
	defer span.End()

	var prjFounds []*entities.Projeto
	cursor, err := p.projetoColecao.Find(ctx, bson.D{})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			trace.AddSpanError(span, err)
			return nil, errors.Wrap(err, "Não existe projeto com este nome!")
		}
		trace.AddSpanError(span, err)
		return nil, errors.Wrap(err, "Erro ao recupurar todos os projetos")
	}
	if err = cursor.All(p.ctx, &prjFounds); err != nil {
		trace.AddSpanError(span, err)
		return nil, errors.Wrap(err, "Erro ao ler todos os projetos")
	}
	if len(prjFounds) == 0 {
		return nil, nil
	}
	log.Info("Encontrados projetos")
	return prjFounds, nil
}

/*
*
 */
func (p ProjetoRepository) Update(ctx context.Context, nome string, projeto *entities.Projeto) (*entities.Projeto, error) {
	_, span := trace.NewSpan(ctx, "ProjetoRepository.Update")
	defer span.End()

	filter := bson.D{{"nome", nome}}
	_, err := p.projetoColecao.ReplaceOne(ctx, filter, projeto)
	if err != nil {
		trace.AddSpanError(span, err)
		return nil, errors.Wrap(err, "Não foi possível atualizar o projeto")
	}
	return projeto, nil
}

func (p ProjetoRepository) Delete(ctx context.Context, nome string) error {
	_, span := trace.NewSpan(ctx, "ProjetoRepository.Delete")
	defer span.End()

	filter := bson.D{primitive.E{Key: "nome", Value: nome}}
	_, err := p.projetoColecao.DeleteOne(ctx, filter)
	if err != nil {
		trace.AddSpanError(span, err)
		return errors.Wrap(err, "Projeto não foi excluído!")
	}
	return nil
}
