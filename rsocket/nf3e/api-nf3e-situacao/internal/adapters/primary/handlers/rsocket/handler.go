package rsocket

import (
	"encoding/json"
	"encoding/xml"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/configs"
	dtos2 "github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/models/dtos"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/ports"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/utils"
	"github.com/rs/zerolog/log"
	"github.com/rsocket/rsocket-go/payload"
)

type Nf3eSituacaoHandler interface {
	GetNf3eSituacao(msg payload.Payload) (dtos2.RetConsSitNF3e, error)
}
type nf3eSituacaoHandler struct {
	Nf3eSituacaoService ports.Nf3eSituacaoService
}

func NewRSocketHandler(situacaoService ports.Nf3eSituacaoService) Nf3eSituacaoHandler {
	return &nf3eSituacaoHandler{Nf3eSituacaoService: situacaoService}
}

func (services *nf3eSituacaoHandler) GetNf3eSituacao(msg payload.Payload) (dtos2.RetConsSitNF3e, error) {

	// Decodificar o Payload da mensagem
	//conSitNF3e := JsonUnmarshal(msg.DataUTF8())
	conSitNF3e, err := XmlUnmarshall(msg.DataUTF8())

	if err != nil {
		return dtos2.RetConsSitNF3e{
			Versao:         configs.Get().VersaoLeiaute,
			Xmlns:          configs.Get().Xmlns,
			TpAmb:          configs.Get().TpAmb,
			VerAplic:       configs.Get().VerAplic,
			Cstat:          "999",
			Xmotivo:        "Rejeição: Erro não catalogado",
			Cuf:            configs.Get().CUF,
			Protnf3e:       "",
			Proceventonf3e: nil,
		}, nil
	} else {
		return services.Nf3eSituacaoService.GetNf3eSituacao(conSitNF3e)
	}
}

/*
	Converte a string json do payload em Classe
*/
func JsonUnmarshal(payload string) dtos2.ConsSitNF3e {

	var conSitNF3e dtos2.ConsSitNF3e

	// Decodifica a entrada JSON para a entidade municipio
	err := json.Unmarshal([]byte(payload), &conSitNF3e)

	if err != nil {
		log.Err(err).
			Str("service", "api-nf3e-situacao").
			Str("component", "rsocket_handler").
			Msgf("Error convert JSON payload (%s)", payload)
	}

	log.Info().
		Str("service", "api-nf3e-situacao").
		Str("module", "rsocket_handler").
		Str("chNF3e", conSitNF3e.ChNF3e).
		Msg(utils.JsonMarshal(conSitNF3e))

	return conSitNF3e
}

func XmlUnmarshall(payload string) (dtos2.ConsSitNF3e, error) {

	var conSitNF3e dtos2.ConsSitNF3e

	err := xml.Unmarshal([]byte(payload), &conSitNF3e)
	if err != nil {
		log.Err(err).
			Str("service", "api-nf3e-situacao").
			Str("component", "rsocket_handler").
			Msgf("Error convert XML payload (%s)", payload)
		return conSitNF3e, err
	}

	log.Info().
		Str("service", "api-nf3e-situacao").
		Str("module", "rsocket_handler").
		Str("chNF3e", conSitNF3e.ChNF3e).
		Msg(utils.JsonMarshal(conSitNF3e))

	return conSitNF3e, nil
}
