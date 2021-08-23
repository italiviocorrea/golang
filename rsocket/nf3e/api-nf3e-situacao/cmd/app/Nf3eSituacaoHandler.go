package app

import (
	"bytes"
	"encoding/json"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/model"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/repository"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/rule"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/service"
	"github.com/rsocket/rsocket-go"
	"github.com/rsocket/rsocket-go/payload"
	"github.com/rsocket/rsocket-go/rx/mono"
	"log"
	"time"
)

type Nf3eSituacaoHandler struct {
	Service service.Nf3eSituacaoService
}

func (handler *Nf3eSituacaoHandler) Consultar(consSitNF3e rule.ConsSitNF3e) (model.Nf3eSituacao, error) {
	return handler.Service.Consultar(consSitNF3e)
}

func Responder() rsocket.RSocket {
	return rsocket.NewAbstractSocket(
		rsocket.RequestResponse(func(msg payload.Payload) mono.Mono {

			var consSitNF3e = rule.ConsSitNF3e{
				Versao:       "1.00",
				TpAmb:        "2",
				XServ:        "CONSULTAR",
				ChNF3e:       msg.DataUTF8(),
				ChNF3eDecode: model.ChaveAcesso{},
			}

			serv := service.NewNf3eSituacaoService(repository.NewNf3eSituacaoRepositoryStub())
			handler := Nf3eSituacaoHandler{serv}

			nf3e, err := handler.Consultar(consSitNF3e)

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
