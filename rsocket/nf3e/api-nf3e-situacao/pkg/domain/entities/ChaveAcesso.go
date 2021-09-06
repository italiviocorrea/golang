package entities

import (
	"strconv"
	"time"
)

type ChaveAcesso struct {
	Cuf          string
	Aamm         string
	Cnpj         string
	Mod          string
	Serie        string
	Nnf          string
	Tpemis       string
	Nsiteautoriz string
	Cnf3e        string
	Cdv          string
	ChNF3e       string
}

func (a *ChaveAcesso) Decode(chNF3e string) {
	a.Cuf = chNF3e[0:2]
	a.Aamm = chNF3e[2:6]
	a.Cnpj = chNF3e[6:20]
	a.Mod = chNF3e[20:22]
	a.Serie = chNF3e[22:25]
	a.Nnf = chNF3e[25:34]
	a.Tpemis = chNF3e[34:35]
	a.Nsiteautoriz = chNF3e[35:36]
	a.Cnf3e = chNF3e[36:43]
	a.Cdv = chNF3e[43:44]
	a.ChNF3e = chNF3e
}

/*
	Valida o digito verificador da chave de acesso.
*/
func (a *ChaveAcesso) IsValidDv() bool {

	var pesos []int = []int{4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	var chNF3eSemDv = a.ChNF3e[0:43]
	var ponderacoes = 0

	var aChNF3eSemDv []string = []string{}
	for i := 0; i < len(chNF3eSemDv); i++ {
		aChNF3eSemDv = append(aChNF3eSemDv, chNF3eSemDv[i:i+1])
	}

	for i := 0; i < len(chNF3eSemDv); i++ {
		var dig, _ = strconv.Atoi(aChNF3eSemDv[i])
		ponderacoes += dig * pesos[i]
	}
	var dv = 11 - (ponderacoes % 11)
	chNF3eSemDv += strconv.Itoa(dv)

	if a.ChNF3e != chNF3eSemDv {
		return false
	}

	return true
}

/*
	Valida a Chave de acesso
*/
func (a ChaveAcesso) Validate() ResultadoProcessamento {

	var resp = ResultadoProcessamento{
		CStat:       "100",
		XMotivo:     "Ok",
		Complemento: "ruleI04Rej236",
	}

	for _, rule := range validationRules {
		result := rule(a)
		if result != "" {
			return ResultadoProcessamento{
				CStat:       "236",
				XMotivo:     result,
				Complemento: "ChaveAcesso",
			}
		}
	}
	return resp
}

var validationRules = map[int]func(m ChaveAcesso) string{
	0: func(m ChaveAcesso) string {
		var cnpj = Cnpj{Value: m.Cnpj}
		if !cnpj.IsValid() {
			return "Rejeição: Chave de Acesso inválida [Motivo: CNPJ zerado ou inválido]"
		}
		return ""

	},
	1: func(m ChaveAcesso) string {
		ano, _ := strconv.Atoi("20" + m.Aamm[0:2])
		data := time.Now()
		anoCorrente := data.Year()
		if ano < 2019 || ano > anoCorrente {
			return "Rejeição: Chave de Acesso inválida [Motivo: Ano < 2019 ou maior que o atual]"
		}
		return ""
	},
	2: func(m ChaveAcesso) string {
		mes, _ := strconv.Atoi(m.Aamm[2:4])
		if mes < 1 || mes > 12 {
			return "Rejeição: Chave de Acesso inválida [Motivo: Mês inválido (0 ou > 12)]"
		}
		return ""
	},
	3: func(m ChaveAcesso) string {
		if m.Mod != "66" {
			return "Rejeição: Chave de Acesso inválida [Motivo: Modelo diferente de 66, Número zerado]"
		}
		return ""
	},
	4: func(m ChaveAcesso) string {
		if m.Tpemis != "1" && m.Tpemis != "2" {
			return "Rejeição: Chave de Acesso inválida [Motivo: Tipo de emissão inválido]"
		}
		return ""
	},
	5: func(m ChaveAcesso) string {
		if m.Cuf != "50" {
			return "Rejeição: Chave de Acesso inválida [Motivo: UF inválida]"
		}
		return ""
	},
	6: func(m ChaveAcesso) string {
		if !m.IsValidDv() {
			return "Rejeição: Chave de Acesso inválida [Motivo: DV inválido]"
		}
		return ""
	},
}
