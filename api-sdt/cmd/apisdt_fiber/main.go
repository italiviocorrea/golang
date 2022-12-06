package main

import (
	"api-sdt/internal/app/config"
	"api-sdt/internal/app/database/mongodb"
	"api-sdt/internal/app/transport/rest_fiber"
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
	app := rest_fiber.New(cfg, dbc.Client)
	app.Start()

}
