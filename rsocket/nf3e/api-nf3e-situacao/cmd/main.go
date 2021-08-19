package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/Cassandra"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/persistence"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/util"
	"github.com/rsocket/rsocket-go"
	"github.com/rsocket/rsocket-go/payload"
	"github.com/rsocket/rsocket-go/rx/mono"
	"log"
	"time"
)

func main() {

	log.Println("Criando uma sessao cassandra.")
	session := Cassandra.Session
	defer session.Close()

	err := rsocket.Receive().
		Acceptor(func(ctx context.Context, setup payload.SetupPayload, sendingSocket rsocket.CloseableRSocket) (rsocket.RSocket, error) {
			// bind responder
			return rsocket.NewAbstractSocket(
				rsocket.RequestResponse(func(msg payload.Payload) mono.Mono {
					nf3e, err := persistence.FindByChNF3e(msg.DataUTF8())
					if err != nil {
						return mono.Just(payload.NewString(err.Error(), time.Now().String()))
					}
					j, _ := JSONMarshal(nf3e)
					log.Println(nf3e)
					return mono.Just(payload.New(j, []byte(time.Now().String())))
				}),
			), nil
		}).
		Transport(rsocket.TCPServer().SetAddr(util.AppConfig.Server).Build()).
		Serve(context.Background())

	log.Fatalln(err)

}

func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}
