package commons

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/configs"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/internal/adapters/inbounds/dtos"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/entities"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/utils"
	"github.com/rs/zerolog/log"
	"gopkg.in/jeevatkm/go-model.v1"
)

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

/**
Copia os dados do DTO para a entidade
*/
func MapperFromDTO(conSitNF3eDTO dtos.ConsSitNF3eDTO) entities.ConsSitNF3e {
	consSitNF3e := entities.ConsSitNF3e{}
	model.Copy(&consSitNF3e, conSitNF3eDTO)
	return consSitNF3e
}

func MapperToDTO(nf3e entities.Nf3eSituacao) dtos.RetConsSitNF3e {

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

func RetConsSitNF3eRejeitada(cStat string, xMotivo string) dtos.RetConsSitNF3e {
	return dtos.RetConsSitNF3e{
		Versao:         configs.Get().VersaoLeiaute,
		Xmlns:          configs.Get().Xmlns,
		TpAmb:          configs.Get().TpAmb,
		VerAplic:       configs.Get().VerAplic,
		Cstat:          cStat,
		Xmotivo:        xMotivo,
		Cuf:            configs.Get().CUF,
		Protnf3e:       "",
		Proceventonf3e: nil,
	}
}
