package main

import (
	"github.com/italiviocorrea/golang/ibge/ibge_cassandra/Cassandra"
	"log"
	"github.com/italiviocorrea/golang/ibge/ibge_cassandra/routers"
	"github.com/italiviocorrea/golang/commons"
)

func main()  {

    // Pegando um sessao do cassandra
	log.Println("Criando uma sessao cassandra.")
	CassandraSession := Cassandra.Session
	defer CassandraSession.Close()

	// Configurando e iniciando o servidor HTTP
	commons.StartServer(commons.AppConfig.Server,routers.InitRoutes())

}


