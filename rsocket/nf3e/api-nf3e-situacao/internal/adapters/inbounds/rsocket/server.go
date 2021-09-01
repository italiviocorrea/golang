package rsocket

import (
	"context"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/configs"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/internal/adapters/inbounds/commons"
	rsocket2 "github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/internal/adapters/inbounds/handlers/rsocket"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/internal/adapters/outbounds/db"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/internal/adapters/outbounds/repositories/nf3e_qry"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/usescases"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/utils"
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
				Str("component", "rsocket.server").
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
			Str("component", "rsocket.server").
			Msg("Erro Fatal no servidor RSocket")
		panic(err)
	}

}

func responder(clientDB db.ClientDB) rsocket.RSocket {
	return rsocket.NewAbstractSocket(
		rsocket.RequestResponse(func(msg payload.Payload) mono.Mono {

			repo := nf3e_qry.NewNf3eSituacaoRepositoryCassandra(clientDB)
			serv := usescases.NewNf3eSituacaoUseCase(repo)
			handler := rsocket2.NewRSocketHandler(serv)

			nf3e, err := handler.GetNf3eSituacao(msg)

			if err != nil {
				log.Err(err).
					Str("service", "api-nf3e-situacao").
					Str("component", "rsocket.server").
					Msgf("Erro ao pesquisar chave de acesso (%s)", msg)

				return mono.Just(payload.NewString(
					utils.JsonMarshal(commons.RetConsSitNF3eRejeitada("999", "Rejeição: Erro não catalogado")),
					time.Now().String()))
			}

			//j, _ := JSONMarshal(nf3e)
			j, _ := commons.XmlMarshal(nf3e)

			return mono.Just(payload.New(j, []byte(time.Now().String())))
		}),
	)
}
