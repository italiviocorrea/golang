package entities

type TipoEmissao int

const (
	NORMAL TipoEmissao = iota + 1
	OFFLINE
)

func (t TipoEmissao) IsValid() bool {
	switch t {
	case NORMAL, OFFLINE:
		return true
	}
	return false
}

func (t TipoEmissao) String() string {
	return [...]string{"NORMAL", "OFFLINE"}[t-1]
}

func (t TipoEmissao) Index() int {
	return int(t)
}
