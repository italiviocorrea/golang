package entities

type TipoAmbiente int

const (
	PRODUCAO TipoAmbiente = iota + 1
	HOMOLOGACAO
	DESENVOLVIMENTO
)

func (tp TipoAmbiente) IsValid() bool {
	switch tp {
	case PRODUCAO, HOMOLOGACAO, DESENVOLVIMENTO:
		return true
	}
	return false
}

func (tp TipoAmbiente) String() string {
	return [...]string{"PRODUCAO", "HOMOLOGACAO", "DESENVOLVIMENTO"}[tp-1]
}

func (tp TipoAmbiente) Index() int {
	return int(tp)
}
