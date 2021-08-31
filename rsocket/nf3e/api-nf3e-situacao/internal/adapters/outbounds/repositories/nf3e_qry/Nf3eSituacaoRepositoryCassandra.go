package nf3e_qry

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/internal/adapters/outbounds/db"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/entities"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/pkg/domain/ports"
	"github.com/rs/zerolog/log"
)

type nf3eSituacaoRepositoryCassandra struct {
	DB db.ClientDB
}

func NewNf3eSituacaoRepositoryCassandra(DB db.ClientDB) ports.Nf3eSituacaoRepositoryPort {
	return &nf3eSituacaoRepositoryCassandra{
		DB: DB,
	}
}

func (n *nf3eSituacaoRepositoryCassandra) FindByID(chnf3e string) (entities.Nf3eSituacao, error) {

	var nf3e entities.Nf3eSituacao

	err := n.DB.DB().Query("select chnf3e,versao,tpamb,cstat,xmotivo,cuf,protnf3e,proceventonf3e FROM nf3e_situacao where chnf3e = ?", chnf3e).
		Scan(&nf3e.Chnf3e, &nf3e.Versao, &nf3e.TpAmb, &nf3e.Cstat, &nf3e.Xmotivo, &nf3e.Cuf, &nf3e.Protnf3e, &nf3e.Proceventonf3e)

	if err != nil {
		log.Err(err).
			Str("service", "api-nf3e-situacao").
			Str("component", "Nf3eSituacaoRepositoryCassandra").
			Str("chNF3e", chnf3e).
			Msg("Chave de acesso não localizada")
		return nf3e, err
	}
	return nf3e, nil

}
