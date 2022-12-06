package main

import (
	"api-sdt/internal/app/config"
	"api-sdt/internal/app/database/mongodb"
	"api-sdt/internal/app/transport/rest_echo"
	"github.com/labstack/gommon/log"
)

func main() {

	// Lê as configurações
	cfg := config.New()

	log.Info(cfg)

	// conecta com o mongodb
	dbc := mongodb.NewConnection(cfg)
	defer dbc.Disconnect()

	// inicia o servidor
	app := rest_echo.New(cfg, dbc.Client)
	app.Start()

}
