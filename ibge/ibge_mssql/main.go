package main

import (
	"com/ItalivioCorrea/commons"
	"com/ItalivioCorrea/ibge/ibge_mssql/routers"
	"log"
	"net/http"
	"com/ItalivioCorrea/ibge/ibge_mssql/Mssql"
)

func main() {

	// Pegando uma conexao com o MSSQL
	MssqlDatabase := Mssql.Database
	defer MssqlDatabase.Close()

	// Criando as rotas
	log.Println("Iniciando rotas...")
	router := routers.InitRoutes()

	// Criando e configurando o servidor
	log.Println("Configurando o servidor..." + commons.AppConfig.Server)

	server := &http.Server{
		Addr:    commons.AppConfig.Server,
		Handler: router,
	}

	log.Println("Executando o servidor...")

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("Error on creating listener: ", err)
	}
	log.Println("Servidor pronto...")

}
