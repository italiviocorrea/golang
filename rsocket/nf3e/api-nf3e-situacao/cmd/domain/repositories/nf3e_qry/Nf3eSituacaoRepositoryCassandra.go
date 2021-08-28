package nf3e_qry

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/configs/db"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/models/entities"
	"log"
)

type nf3eSituacaoRepositoryCassandra struct {
	DB db.ClientDB
}

func NewNf3eSituacaoRepositoryCassandra(DB db.ClientDB) Nf3eSituacaoRepositoryInterface {
	return &nf3eSituacaoRepositoryCassandra{
		DB: DB,
	}
}

func (n *nf3eSituacaoRepositoryCassandra) FindByID(chnf3e string) (entities.Nf3eSituacao, error) {

	var nf3e entities.Nf3eSituacao

	err := n.DB.DB().Query("select chnf3e,versao,tpamb,cstat,xmotivo,cuf,protnf3e,proceventonf3e FROM nf3e_situacao where chnf3e = ?", chnf3e).
		Scan(&nf3e.Chnf3e, &nf3e.Versao, &nf3e.Tpamb, &nf3e.Cstat, &nf3e.Xmotivo, &nf3e.Cuf, &nf3e.Protnf3e, &nf3e.Proceventonf3e)

	if err != nil {
		log.Printf("Error reading rows: " + err.Error())
		return nf3e, err
	}
	return nf3e, nil

}
