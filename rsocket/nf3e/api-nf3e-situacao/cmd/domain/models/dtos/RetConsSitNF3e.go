package dtos

type RetConsSitNF3e struct {
	Versao         string   `json:"versao"`
	TpAmb          string   `json:"tpAmb"`
	VerAplic       string   `json:"verAplic"`
	Cstat          string   `json:"cStat"`
	Xmotivo        string   `json:"xMotivo"`
	Cuf            string   `json:"cUF"`
	Protnf3e       string   `json:"protNF3e,omitempty"`
	Proceventonf3e []string `json:"procEventoNF3e,omitempty"`
}
