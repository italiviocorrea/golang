package app

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/configs"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/configs/db"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/models/dtos"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/repositories/nf3e_qry"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/utils"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/handlers"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/services"
	"github.com/rs/zerolog/log"
	"github.com/rsocket/rsocket-go"
	"github.com/rsocket/rsocket-go/payload"
	"github.com/rsocket/rsocket-go/rx/mono"
	"strconv"
	"time"
)

func Server(clientDB db.ClientDB) {

	host := configs.Get().Server
	port := configs.Get().Port

	err := rsocket.Receive().
		OnStart(func() {
			log.Info().
				Str("service", "api-nf3e-situacao").
				Str("component", "rsocket_server").
				Str("host", host).
				Str("port", strconv.Itoa(port)).
				Msg("Servidor RSocket Iniciado com sucesso!")
		}).
		Acceptor(func(ctx context.Context, setup payload.SetupPayload, sendingSocket rsocket.CloseableRSocket) (rsocket.RSocket, error) {
			// bind responder
			return responder(clientDB), nil
		}).
		Transport(rsocket.TCPServer().SetHostAndPort(host, port).Build()).
		Serve(context.Background())

	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "api-nf3e-situacao").
			Str("component", "rsocket_server").
			Msg("Erro Fatal no servidor RSocket")
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
				log.Err(err).
					Str("service", "api-nf3e-situacao").
					Str("component", "rsocket_server").
					Msgf("Erro ao pesquisar chave de acesso (%s)", msg)

				return mono.Just(payload.NewString(utils.JsonMarshal(dtos.RetConsSitNF3e{
					Versao:         configs.Get().VersaoLeiaute,
					Xmlns:          configs.Get().Xmlns,
					TpAmb:          configs.Get().TpAmb,
					VerAplic:       configs.Get().VerAplic,
					Cstat:          "999",
					Xmotivo:        "Rejeição: Erro não catalogado",
					Cuf:            configs.Get().CUF,
					Protnf3e:       "",
					Proceventonf3e: nil,
				}), time.Now().String()))
			}

			//j, _ := JSONMarshal(nf3e)
			j, _ := XmlMarshal(nf3e)

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

func XmlMarshal(t interface{}) ([]byte, error) {

	buffer := &bytes.Buffer{}
	enc := xml.NewEncoder(buffer)
	enc.Indent("  ", "    ")
	err := enc.Encode(t)
	return buffer.Bytes(), err

}
