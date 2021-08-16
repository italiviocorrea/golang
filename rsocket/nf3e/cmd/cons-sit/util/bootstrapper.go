package util

import "log"

func StartUp(varEnvPrefix string) {

	log.Println("Inicializando configuracoes...")
	initConfig(varEnvPrefix)

	log.Println("Configurando nivel de log...")
	setLogLevel(Level(AppConfig.LogLevel))

}
