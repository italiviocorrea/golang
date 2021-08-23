package model

type Nf3eSituacao struct {
	Chnf3e         string   `json:"chNF3e,omitempty"`
	Versao         string   `json:"versao,omitempty"`
	Tpamb          string   `json:"tpAmb,omitempty"`
	Cstat          string   `json:"cStat,omitempty"`
	Xmotivo        string   `json:"xMotivo,omitempty"`
	Cuf            string   `json:"cUF,omitempty"`
	Protnf3e       string   `json:"protNF3e,omitempty"`
	Proceventonf3e []string `json:"procEventoNF3e,omitempty"`
}
