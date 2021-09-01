package rsockets

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/internal/adapters/inbounds/commons"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/internal/adapters/inbounds/dtos"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/ports"
	"github.com/rsocket/rsocket-go/payload"
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
	conSitNF3eDTO, err := commons.XmlUnmarshall(msg.DataUTF8())

	if err != nil {
		return commons.RetConsSitNF3eRejeitada("999", "Rejeição: Erro não catalogado"), nil
	} else {
		nf3e, _ := services.Nf3eSituacaoUseCase.GetNf3eSituacao(commons.MapperFromDTO(conSitNF3eDTO))
		retConsSitNF3e := commons.MapperToDTO(nf3e)
		return retConsSitNF3e, nil
	}
}
