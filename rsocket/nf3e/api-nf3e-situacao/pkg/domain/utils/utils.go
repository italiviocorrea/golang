package utils

import (
	"encoding/json"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/entities"
)

type Chan chan entities.ResultadoProcessamento

func Reduce(in Chan) entities.ResultadoProcessamento {
	resp := <-in
	return resp
}

func Count(in []entities.ResultadoProcessamento) int {
	return len(in)
}

func Take(in []entities.ResultadoProcessamento, nmax int) []entities.ResultadoProcessamento {
	var out []entities.ResultadoProcessamento
	index := 0
	for _, s := range in {
		index++
		if index <= nmax {
			out = append(out, s)
		}
	}
	return out
}

func TakeChan(in chan entities.ResultadoProcessamento, size int) []entities.ResultadoProcessamento {
	var out []entities.ResultadoProcessamento
	remaining := size
	for remaining > 0 {
		result := <-in
		remaining -= 1
		out = append(out, result)
	}
	return out
}

func FilterRejects(p []entities.ResultadoProcessamento) []entities.ResultadoProcessamento {
	var out []entities.ResultadoProcessamento
	for _, resp := range p {
		if resp.CStat != "100" {
			out = append(out, resp)
		}
	}
	return out
}

func IsRejects(resps []entities.ResultadoProcessamento) bool {

	rejects := FilterRejects(resps)
	rejectsCount := Count(rejects)

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
