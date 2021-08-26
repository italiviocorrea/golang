package dtos

type ConsSitNF3e struct {
	Versao       string `json:"versao"`
	TpAmb        string `json:"tpAmb"`
	XServ        string `json:"xServ"`
	ChNF3e       string `json:"chNF3e"`
	ChNF3eDecode ChaveAcesso
}
