package utils

import (
	"encoding/json"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/models/dtos"
	"github.com/rs/zerolog/log"
)

type Chan chan dtos.ResultadoProcessamento

func Reduce(in Chan) dtos.ResultadoProcessamento {
	resp := <-in
	return resp
}

func Count(in []dtos.ResultadoProcessamento) int {
	return len(in)
}

func Take(in []dtos.ResultadoProcessamento, nmax int) []dtos.ResultadoProcessamento {
	var out []dtos.ResultadoProcessamento
	index := 0
	for _, s := range in {
		index++
		if index <= nmax {
			out = append(out, s)
		}
	}
	return out
}

func FilterRejects(p []dtos.ResultadoProcessamento) []dtos.ResultadoProcessamento {
	var out []dtos.ResultadoProcessamento
	for _, resp := range p {
		if resp.CStat != "100" {
			out = append(out, resp)
		}
	}
	return out
}

func IsRejects(resps []dtos.ResultadoProcessamento) bool {

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
