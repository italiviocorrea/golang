package main

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/internal/adapters/primary/rsocket"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/internal/adapters/secondary/db"
)

func main() {

	// Pega a conexão com o banco de dados
	cassadraDB := db.NewCassandraClient()
	defer cassadraDB.DB().Close()

	// Inicia o servidor rsocket
	rsocket.Server(cassadraDB)

}
