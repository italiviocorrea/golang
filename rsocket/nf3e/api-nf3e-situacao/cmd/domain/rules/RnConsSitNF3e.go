package rules

//type RespostaValidacao struct {
//	CStat       string
//	XMotivo     string
//	Complemento string
//}
//
//type ConsSitNF3e struct {
//	Versao       string `json:"versao"`
//	TpAmb        string `json:"tpAmb"`
//	XServ        string `json:"xServ"`
//	ChNF3e       string `json:"chNF3e"`
//	ChNF3eDecode dtos.ChaveAcesso
//}
//
//func (e *ConsSitNF3e) Validar() {
//	// decodificar a chave de acesso
//	chave := dtos.ChaveAcesso{}
//	chave.Decode(e.ChNF3e)
//	e.ChNF3eDecode = chave
//	fmt.Println(e)
//	var resp478 = e.rnI03Rej478()
//	fmt.Println(resp478)
//	var resp236 = e.rnI04Rej236()
//	fmt.Println(resp236)
//
//}
//
///*
//	Validar se o Tipo do ambiente informado difere do ambiente do Web Nf3eSituacaoService
//*/
//func (e *ConsSitNF3e) rnI01Rej252() RespostaValidacao {
//
//	var resp = RespostaValidacao{
//		CStat:       "100",
//		XMotivo:     "Ok",
//		Complemento: "",
//	}
//
//	if e.TpAmb != configs.Get().TpAmb {
//		resp.CStat = "252"
//		resp.XMotivo = "Rejeicao: Ambiente informado diverge do Ambiente de recebimento"
//	}
//
//	return resp
//}
//
///*
//	Validar UF da chave de acesso difere da UF do Web Nf3eSituacaoService
//*/
//func (e *ConsSitNF3e) rnI02Rej226() RespostaValidacao {
//
//	var resp = RespostaValidacao{
//		CStat:       "100",
//		XMotivo:     "Ok",
//		Complemento: "",
//	}
//
//	if e.ChNF3eDecode.Cuf != configs.Get().CUF {
//		resp.CStat = "226"
//		resp.XMotivo = "Rejeição: Código da UF do Emitente diverge da UF autorizadora"
//	}
//
//	return resp
//}
//
///*
//	Validar se o ano – mês da chave de acesso está com atraso superior a 6 meses em relação ao ano–mês atual
//*/
//func (e *ConsSitNF3e) rnI03Rej478() RespostaValidacao {
//
//	var resp = RespostaValidacao{
//		CStat:       "100",
//		XMotivo:     "Ok",
//		Complemento: "",
//	}
//	anoMesChave, _ := strconv.Atoi("20" + e.ChNF3eDecode.Aamm)
//	anoMes6MesesAntes, _ := strconv.Atoi(getAnoMes6MesesAtras())
//
//	fmt.Printf("%d - %d", anoMesChave, anoMes6MesesAntes)
//
//	if anoMesChave < anoMes6MesesAntes {
//		resp.CStat = "478"
//		resp.XMotivo = "Rejeicao: Consulta a uma Chave de Acesso muito antiga"
//	}
//
//	return resp
//}
//
///*
//	Calcula a partir da data e hora atual, o ano e mês à 180 dias atrás.
//*/
//func getAnoMes6MesesAtras() string {
//	data := time.Now()
//	oneMonthLater := data.AddDate(0, -6, 0)
//	s := fmt.Sprintf("%04d%02d", oneMonthLater.Year(), int(oneMonthLater.Month()))
//	return s
//}
//
///*
//  	Validar chave de acesso
//	Retornar motivo da rejeição da Chave de Acesso:
//		- CNPJ zerado ou inválido,
//		- Ano < 2019 ou maior que atual,
//        - Mês inválido (0 ou > 12),
//		- Modelo diferente de 66,
//		- Número zerado,
//		- Tipo de emissão inválido,
//		- UF inválida
//		- ou DV inválido)
//	[Motivo: XXXXXXXXXXXX]
//*/
//func (e *ConsSitNF3e) rnI04Rej236() RespostaValidacao {
//
//	var resp = RespostaValidacao{
//		CStat:       "100",
//		XMotivo:     "Ok",
//		Complemento: "",
//	}
//
//	ano, _ := strconv.Atoi("20" + e.ChNF3eDecode.Aamm[0:2])
//	data := time.Now()
//	anoCorrente := data.Year()
//	mes, _ := strconv.Atoi(e.ChNF3eDecode.Aamm[2:4])
//	var cnpj = dtos.Cnpj{Value: e.ChNF3eDecode.Cnpj}
//
//	if !cnpj.IsValid() {
//		resp.CStat = "236"
//		resp.XMotivo = "Rejeição: Chave de Acesso inválida [Motivo: CNPJ zerado ou inválido]"
//	} else if ano < 2019 || ano > anoCorrente {
//		resp.CStat = "236"
//		resp.XMotivo = "Rejeição: Chave de Acesso inválida [Motivo: Ano < 2019 ou maior que o atual]"
//	} else if mes < 1 || mes > 12 {
//		resp.CStat = "236"
//		resp.XMotivo = "Rejeição: Chave de Acesso inválida [Motivo: Mês inválido (0 ou > 12)]"
//	} else if e.ChNF3eDecode.Mod != "66" {
//		resp.CStat = "236"
//		resp.XMotivo = "Rejeição: Chave de Acesso inválida [Motivo: Modelo diferente de 66, Número zerado]"
//	} else if e.ChNF3eDecode.Tpemis != "1" && e.ChNF3eDecode.Tpemis != "2" {
//		resp.CStat = "236"
//		resp.XMotivo = "Rejeição: Chave de Acesso inválida [Motivo: Tipo de emissão inválido]"
//	} else if e.ChNF3eDecode.Cuf != "50" {
//		resp.CStat = "236"
//		resp.XMotivo = "Rejeição: Chave de Acesso inválida [Motivo: UF inválida]"
//	} else if !e.ChNF3eDecode.IsValidDv() {
//		resp.CStat = "236"
//		resp.XMotivo = "Rejeição: Chave de Acesso inválida [Motivo: DV inválido]"
//	}
//
//	return resp
//
//}
//
///*
//	Validar se o Site de autorização da chave de acesso da NF3e difere do Site de Recebimento
//*/
//func (e *ConsSitNF3e) rnI05Rej482() RespostaValidacao {
//
//	var resp = RespostaValidacao{
//		CStat:       "100",
//		XMotivo:     "Ok",
//		Complemento: "",
//	}
//
//	if e.ChNF3eDecode.Nsiteautoriz != "0" {
//		resp.CStat = "482"
//		resp.XMotivo = "Rejeição: Site de autorização inválido"
//	}
//
//	return resp
//}
