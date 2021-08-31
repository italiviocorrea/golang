package rsocket

import (
	"encoding/json"
	"encoding/xml"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/configs"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/internal/adapters/inbounds/dtos"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/entities"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/ports"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/utils"
	"github.com/rs/zerolog/log"
	"github.com/rsocket/rsocket-go/payload"
	"gopkg.in/jeevatkm/go-model.v1"
)

type Nf3eSituacaoHandler interface {
	GetNf3eSituacao(msg payload.Payload) (dtos.RetConsSitNF3e, error)
}
type nf3eSituacaoHandler struct {
	Nf3eSituacaoUseCase ports.Nf3eSituacaoUseCasePort
}

func NewRSocketHandler(situacaoService ports.Nf3eSituacaoUseCasePort) Nf3eSituacaoHandler {
	return &nf3eSituacaoHandler{Nf3eSituacaoUseCase: situacaoService}
}

func (services *nf3eSituacaoHandler) GetNf3eSituacao(msg payload.Payload) (dtos.RetConsSitNF3e, error) {

	// Decodificar o Payload da mensagem
	//conSitNF3e := JsonUnmarshal(msg.DataUTF8())
	conSitNF3eDTO, err := XmlUnmarshall(msg.DataUTF8())

	if err != nil {
		return dtos.RetConsSitNF3e{
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
		nf3e, _ := services.Nf3eSituacaoUseCase.GetNf3eSituacao(mapperFromDTO(conSitNF3eDTO))

		retConsSitNF3e := mapperToDTO(nf3e)

		return retConsSitNF3e, nil
	}
}

/**
Copia os dados do DTO para a entidade
*/
func mapperFromDTO(conSitNF3eDTO dtos.ConsSitNF3eDTO) entities.ConsSitNF3e {
	consSitNF3e := entities.ConsSitNF3e{}
	model.Copy(&consSitNF3e, conSitNF3eDTO)
	return consSitNF3e
}

func mapperToDTO(nf3e entities.Nf3eSituacao) dtos.RetConsSitNF3e {

	retConsSitNF3e := dtos.RetConsSitNF3e{}

	model.Copy(&retConsSitNF3e, nf3e)

	// configura os campos : verAplic e Xmlns com valores default
	retConsSitNF3e.Xmlns = configs.Get().Xmlns
	retConsSitNF3e.VerAplic = configs.Get().VerAplic

	return retConsSitNF3e
}

/*
	Converte a string json do payload em Classe
*/
func JsonUnmarshal(payload string) dtos.ConsSitNF3eDTO {

	var conSitNF3e dtos.ConsSitNF3eDTO

	// Decodifica a entrada JSON para a entidade municipio
	err := json.Unmarshal([]byte(payload), &conSitNF3e)

	if err != nil {
		log.Err(err).
			Str("service", "api-nf3e-situacao").
			Str("component", "rsocket.handler").
			Msgf("Error convert JSON payload (%s)", payload)
	}

	log.Info().
		Str("service", "api-nf3e-situacao").
		Str("module", "rsocket.handler").
		Str("chNF3e", conSitNF3e.ChNF3e).
		Msg(utils.JsonMarshal(conSitNF3e))

	return conSitNF3e
}

func XmlUnmarshall(payload string) (dtos.ConsSitNF3eDTO, error) {

	var conSitNF3e dtos.ConsSitNF3eDTO

	err := xml.Unmarshal([]byte(payload), &conSitNF3e)
	if err != nil {
		log.Err(err).
			Str("service", "api-nf3e-situacao").
			Str("component", "rsocket.handler").
			Msgf("Error convert XML payload (%s)", payload)
		return conSitNF3e, err
	}

	log.Info().
		Str("service", "api-nf3e-situacao").
		Str("module", "rsocket.handler").
		Str("chNF3e", conSitNF3e.ChNF3e).
		Msg(utils.JsonMarshal(conSitNF3e))

	return conSitNF3e, nil
}
