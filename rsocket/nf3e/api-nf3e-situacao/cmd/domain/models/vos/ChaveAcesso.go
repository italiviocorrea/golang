package vos

import (
	"strconv"
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
