package main

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/adapters/configs/db"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/adapters/endpoints"
)

func main() {

	// Pega a conexão com o banco de dados
	cassadraDB := db.NewCassandraClient()
	defer cassadraDB.DB().Close()

	// Inicia o servidor rsocket
	endpoints.Server(cassadraDB)

}
