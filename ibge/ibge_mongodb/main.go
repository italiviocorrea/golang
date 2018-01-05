package main

import (
	"log"
	"com/ItalivioCorrea/commons"
	"com/ItalivioCorrea/ibge/ibge_mongodb/Mongodb"
	"com/ItalivioCorrea/ibge/ibge_mongodb/routers"
)

func main()  {

    // Pegando um sessao do cassandra
	log.Println("Criando uma sessao mongodb.")
	MongoDBSession := Mongodb.Session
	defer MongoDBSession.Close()

	// Configurando e iniciando o servidor HTTP
	commons.StartServer(commons.AppConfig.Server,routers.InitRoutes())

}


