package main

import (
	"context"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/app"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/config"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/config/Cassandra"
	"github.com/rsocket/rsocket-go"
	"github.com/rsocket/rsocket-go/payload"
	"log"
)

func main() {

	log.Println("Criando uma sessao cassandra.")
	session := Cassandra.Session
	defer session.Close()

	err := rsocket.Receive().
		Acceptor(func(ctx context.Context, setup payload.SetupPayload, sendingSocket rsocket.CloseableRSocket) (rsocket.RSocket, error) {
			// bind responder
			return app.Responder(), nil
		}).
		Transport(rsocket.TCPServer().SetAddr(config.AppConfig.Server).Build()).
		Serve(context.Background())

	log.Fatalln(err)
}
