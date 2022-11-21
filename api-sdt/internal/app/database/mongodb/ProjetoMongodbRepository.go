package mongodb

import (
	"api-sdt/internal/app/config"
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
	projetoColecao := mongo.Database(cfg.DbName).Collection("projetos")
	return &ProjetoRepository{
		projetoColecao: projetoColecao,
		ctx:            context.TODO(),
	}
}

func (p ProjetoRepository) Create(projeto *entities.Projeto) (*entities.Projeto, error) {
	_, err := p.projetoColecao.InsertOne(p.ctx, projeto)
	if err != nil {
		return nil, errors.Wrap(err, "Erro ao inserir projeto")
	}
	return projeto, nil
}

func (p ProjetoRepository) FindByName(nome string) (*entities.Projeto, error) {

	var prjFound *entities.Projeto

	filter := bson.D{primitive.E{Key: "nome", Value: nome}}

	err := p.projetoColecao.FindOne(p.ctx, filter).Decode(&prjFound)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Info("Projeto não encontrado !")
			return nil, errors.Wrap(err, "Não existe projeto com este nome!")
		}
		log.Info("Erro ao pesquisar o projeto no MongoDB")
		return nil, errors.Wrap(err, "Erro ao pesquisar projeto por nome")
	}
	return prjFound, nil
}

func (p ProjetoRepository) FindAll() ([]*entities.Projeto, error) {
	var prjFounds []*entities.Projeto
	cursor, err := p.projetoColecao.Find(p.ctx, bson.D{})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Info("Não existe projetos cadastrados !")
			return nil, errors.Wrap(err, "Não existe projeto com este nome!")
		}
		return nil, errors.Wrap(err, "Erro ao recupurar todos os projetos")
	}
	if err = cursor.All(p.ctx, &prjFounds); err != nil {
		return nil, errors.Wrap(err, "Erro ao ler todos os projetos")
	}
	if len(prjFounds) == 0 {
		return nil, nil
	}
	log.Info("Encontrados projetos")
	return prjFounds, nil
}

/**

 */
func (p ProjetoRepository) Update(nome string, projeto *entities.Projeto) (*entities.Projeto, error) {
	log.Info(nome)
	filter := bson.D{{"nome", nome}}
	_, err := p.projetoColecao.ReplaceOne(p.ctx, filter, projeto)
	if err != nil {
		return nil, errors.Wrap(err, "Não foi possível atualizar o projeto")
	}
	return projeto, nil
}

func (p ProjetoRepository) Delete(nome string) error {
	filter := bson.D{primitive.E{Key: "nome", Value: nome}}
	_, err := p.projetoColecao.DeleteOne(p.ctx, filter)
	if err != nil {
		return errors.Wrap(err, "Projeto não foi excluído!")
	}
	return nil
}
