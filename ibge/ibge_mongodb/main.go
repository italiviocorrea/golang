package main

import (
	"log"
	"github.com/italiviocorrea/golang/commons"
	"github.com/italiviocorrea/golang/ibge/ibge_mongodb/Mongodb"
	"github.com/italiviocorrea/golang/ibge/ibge_mongodb/routers"
)

func main()  {

    // Pegando um sessao do mongodb
	log.Println("Criando uma sessao mongodb.")
	MongoDBSession := Mongodb.Session
	defer MongoDBSession.Close()

	// Configurando e iniciando o servidor HTTP
	commons.StartServer(commons.AppConfig.Server,routers.InitRoutes())

}


