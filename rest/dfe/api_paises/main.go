package main

import (
	"./endpoints"
	"./models"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	// Pegando uma conexao com o MSSQL
	db := models.ConnectToDatabase()

	defer db.Close()

	db.AutoMigrate(&models.Paises{})

	// configura o servidor
	srv := &http.Server{
		Addr:    os.Getenv("API_SRV_ADDR"),
		Handler: endpoints.Startup(),
	}

	// executa o servidor http
	go func() {
		// conexões do serviço
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Aguarda o sinal de interrupção, para desligar normalmente o servidor,
	// com um tempo limite de 15 segundos.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")

}


