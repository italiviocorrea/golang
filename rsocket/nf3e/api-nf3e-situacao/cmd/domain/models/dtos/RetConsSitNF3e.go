package dtos

type RetConsSitNF3e struct {
	Xmlns          string   `xml:"xmlns,attr"`
	Versao         string   `xml:"versao,attr"`
	TpAmb          string   `xml:"tpAmb"`
	VerAplic       string   `xml:"verAplic"`
	Cstat          string   `xml:"cStat"`
	Xmotivo        string   `xml:"xMotivo"`
	Cuf            string   `xml:"cUF"`
	Protnf3e       string   `xml:"protNF3e,omitempty"`
	Proceventonf3e []string `xml:"procEventoNF3e,omitempty"`
}
