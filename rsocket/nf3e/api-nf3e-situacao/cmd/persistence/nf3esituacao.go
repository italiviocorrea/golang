package persistence

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/Cassandra"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/model"
	"log"
)

func FindByChNF3e(chnf3e string) (model.Nf3esituacao, error) {

	var resp model.Nf3esituacao

	err := Cassandra.Session.Query("select chnf3e,versao,tpamb,cstat,xmotivo,cuf,protnf3e,proceventonf3e FROM nf3e_situacao where chnf3e = ?", chnf3e).
		Scan(&resp.Chnf3e, &resp.Versao, &resp.Tpamb, &resp.Cstat, &resp.Xmotivo, &resp.Cuf, &resp.Protnf3e, &resp.Proceventonf3e)

	if err != nil {
		log.Printf("Error reading rows: " + err.Error())
		return resp, err
	}
	return resp, err

}
