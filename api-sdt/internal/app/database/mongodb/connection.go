package mongodb

import (
	"api-sdt/internal/app/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type Connection struct {
	Client *mongo.Client
	ctx    context.Context
}

func NewConnection(cfg *config.Settings) Connection {

	uri := fmt.Sprintf("mongodb://%s/%s/", cfg.DbHost, cfg.DbName)

	log.Println(uri)

	credentials := options.Credential{
		Username: cfg.DbUser,
		Password: cfg.DbPwd,
	}

	clientOpts := options.Client().ApplyURI(uri).SetAuth(credentials)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Println("Erro ao conectar ao banco mongodb")
		log.Fatal(err)
	}

	fmt.Println("Conectado ao banco de dados")

	return Connection{
		Client: client,
		ctx:    ctx,
	}
}

func (c Connection) Disconnect() {
	c.Client.Disconnect(c.ctx)
}
