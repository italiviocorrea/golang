package commons

import (
	"log"
	"github.com/gorilla/mux"
	"net/http"
)

func StartUp(varEnvPrefix string) {

	log.Println("Inicializando configuracoes...")
	initConfig(varEnvPrefix)

	log.Println("Configurando nivel de log...")
	setLogLevel(Level(AppConfig.LogLevel))

}

func StartServer(address string, handler *mux.Router)  {

	// Criando e configurando o servidor
	log.Println("Configurando o servidor..." + address)

	server := &http.Server{
		Addr:    address,
		Handler: handler,
	}

	log.Println("Executando o servidor...")

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("Erro ao criar o servidor: ", err)
	}


}
