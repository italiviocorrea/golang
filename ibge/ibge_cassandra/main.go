package main

import (
	"com/ItalivioCorrea/ibge/ibge_cassandra/Cassandra"
	"log"
	"com/ItalivioCorrea/ibge/ibge_cassandra/routers"
	"com/ItalivioCorrea/commons"
)

func main()  {

    // Pegando um sessao do cassandra
	log.Println("Criando uma sessao cassandra.")
	CassandraSession := Cassandra.Session
	defer CassandraSession.Close()

	// Configurando e iniciando o servidor HTTP
	commons.StartServer(commons.AppConfig.Server,routers.InitRoutes())

}


