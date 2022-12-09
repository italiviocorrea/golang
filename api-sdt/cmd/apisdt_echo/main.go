package main

import (
	"api-sdt/internal/app/config"
	"api-sdt/internal/app/database/mongodb"
	"api-sdt/internal/app/trace"
	"api-sdt/internal/app/transport/rest_echo"
	"context"
	"github.com/labstack/gommon/log"
)

func main() {

	// Lê as configurações
	cfg := config.New()

	// Habilita o trancing
	ctx := context.Background()

	prv, err := trace.NewProvider(trace.ProviderConfig{
		JaegerEndpoint: cfg.JaegerEndpoint,
		ServiceName:    "API-SDT",
		ServiceVersion: "1.0.0",
		Environment:    cfg.Env,
		Disabled:       false,
	})

	if err != nil {
		log.Fatal(err)
	}
	defer prv.Close(ctx)

	// conecta com o mongodb
	dbc := mongodb.NewConnection(cfg)
	defer dbc.Disconnect()

	// inicia o servidor
	app := rest_echo.New(cfg, dbc.Client)
	app.Start()

}
