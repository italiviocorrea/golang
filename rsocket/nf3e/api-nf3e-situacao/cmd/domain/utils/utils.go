package utils

import (
	"encoding/json"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/models/dtos"
	"github.com/rs/zerolog/log"
)

type Chan chan dtos.RespostaValidacao

func Reduce(in Chan) dtos.RespostaValidacao {
	resp := <-in
	return resp
}

func Count(in []dtos.RespostaValidacao) int {
	return len(in)
}

func Take(in []dtos.RespostaValidacao, nmax int) []dtos.RespostaValidacao {
	var out []dtos.RespostaValidacao
	index := 0
	for _, s := range in {
		index++
		if index <= nmax {
			out = append(out, s)
		}
	}
	return out
}

func FilterRejects(p []dtos.RespostaValidacao) []dtos.RespostaValidacao {
	var out []dtos.RespostaValidacao
	for _, resp := range p {
		if resp.CStat != "100" {
			out = append(out, resp)
		}
	}
	return out
}

func IsRejects(resps []dtos.RespostaValidacao) bool {

	rejects := FilterRejects(resps)
	rejectsCount := Count(rejects)

	log.Info().
		Str("service", "api-nf3e-situacao").
		Str("module", "utils").
		Msgf("{Rejeicoes:%s}\n", JsonMarshal(rejects))

	if rejectsCount > 0 {
		return true
	} else {
		return false
	}
}

func JsonMarshal(v interface{}) string {
	e, _ := json.Marshal(v)
	return string(e)
}
