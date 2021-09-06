package entities

import (
	"strconv"
)

type Cnpj struct {
	Value string
}

func (c *Cnpj) IsValid() bool {

	cnpj_size := len(c.Value)
	cnpj_zero, _ := strconv.Atoi(c.Value)

	if cnpj_size != 14 || cnpj_zero == 0 {
		return false
	}

	var cnpjCalc = c.Value[0:12]
	var soma = 0

	var aCnpj []string = []string{}
	for i := 0; i < cnpj_size-2; i++ {
		aCnpj = append(aCnpj, cnpjCalc[i:i+1])
	}

	// calcula o primeiro digito
	var peso1 []int = []int{6, 7, 8, 9, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < cnpj_size-2; i++ {
		var dig, _ = strconv.Atoi(aCnpj[i])
		soma += dig * peso1[i]
	}

	var dig = soma % 11
	cnpjCalc += strconv.Itoa(dig)

	// calcula o segundo digito
	var peso2 []int = []int{5, 6, 7, 8, 9, 2, 3, 4, 5, 6, 7, 8, 9}
	var soma2 = 0

	for i := 0; i < cnpj_size-2; i++ {
		var dig, _ = strconv.Atoi(aCnpj[i])
		soma2 += dig * peso2[i]
	}

	var dig2 = (soma2 % 11)
	cnpjCalc += strconv.Itoa(dig2)

	if c.Value != cnpjCalc {
		return false
	}

	return true
}
