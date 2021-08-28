package handlers

import (
	"encoding/json"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/models/dtos"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/services"
	"github.com/rsocket/rsocket-go/payload"
	"log"
)

type Nf3eSituacaoHandler interface {
	GetNf3eSituacao(msg payload.Payload) (dtos.RetConsSitNF3e, error)
}
type nf3eSituacaoHandler struct {
	Nf3eSituacaoService services.Nf3eSituacaoService
}

func NewRSocketHandler(situacaoService services.Nf3eSituacaoService) Nf3eSituacaoHandler {
	return &nf3eSituacaoHandler{Nf3eSituacaoService: situacaoService}
}

func (services *nf3eSituacaoHandler) GetNf3eSituacao(msg payload.Payload) (dtos.RetConsSitNF3e, error) {

	// Decodificar o Payload da mensagem
	conSitNF3e := JsonUnmarshal(msg.DataUTF8())

	return services.Nf3eSituacaoService.GetNf3eSituacao(conSitNF3e)
}

/*
	Converte a string json do payload em Classe
*/
func JsonUnmarshal(payload string) dtos.ConsSitNF3e {

	var conSitNF3e dtos.ConsSitNF3e

	// Decodifica a entrada JSON para a entidade municipio
	err := json.Unmarshal([]byte(payload), &conSitNF3e)

	if err != nil {
		log.Println("Erro ao pegar ao o payload")
	}

	log.Println(conSitNF3e)

	return conSitNF3e
}
