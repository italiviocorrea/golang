package app

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/configs"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/configs/db"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/repositories/nf3e_qry"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/handlers"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/services"
	"github.com/rsocket/rsocket-go"
	"github.com/rsocket/rsocket-go/payload"
	"github.com/rsocket/rsocket-go/rx/mono"
	"log"
	"time"
)

func Server(clientDB db.ClientDB) {

	host := configs.Get().Server
	port := configs.Get().Port

	err := rsocket.Receive().
		OnStart(func() {
			log.Println("server start success! ")
		}).
		Acceptor(func(ctx context.Context, setup payload.SetupPayload, sendingSocket rsocket.CloseableRSocket) (rsocket.RSocket, error) {
			// bind responder
			return responder(clientDB), nil
		}).
		Transport(rsocket.TCPServer().SetHostAndPort(host, port).Build()).
		Serve(context.Background())

	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

}

func responder(clientDB db.ClientDB) rsocket.RSocket {
	return rsocket.NewAbstractSocket(
		rsocket.RequestResponse(func(msg payload.Payload) mono.Mono {

			repo := nf3e_qry.NewNf3eSituacaoRepositoryCassandra(clientDB)
			serv := services.NewNf3eSituacaoService(repo)
			handler := handlers.NewRSocketHandler(serv)

			nf3e, err := handler.GetNf3eSituacao(msg)

			if err != nil {
				return mono.Just(payload.NewString(err.Error(), time.Now().String()))
			}

			j, _ := JSONMarshal(nf3e)

			log.Println(nf3e)

			return mono.Just(payload.New(j, []byte(time.Now().String())))
		}),
	)
}

func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}
