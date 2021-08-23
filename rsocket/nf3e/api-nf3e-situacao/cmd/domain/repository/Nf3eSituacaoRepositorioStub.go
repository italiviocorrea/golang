package repository

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/config/Cassandra"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/domain/model"
	"log"
)

type Nf3eSituacaoRepositorioStub struct {
	nf3e model.Nf3eSituacao
}

func (stub Nf3eSituacaoRepositorioStub) Consultar(chnf3e string) (model.Nf3eSituacao, error) {

	err := Cassandra.Session.Query("select chnf3e,versao,tpamb,cstat,xmotivo,cuf,protnf3e,proceventonf3e FROM nf3e_situacao where chnf3e = ?", chnf3e).
		Scan(&stub.nf3e.Chnf3e, &stub.nf3e.Versao, &stub.nf3e.Tpamb, &stub.nf3e.Cstat, &stub.nf3e.Xmotivo, &stub.nf3e.Cuf, &stub.nf3e.Protnf3e, &stub.nf3e.Proceventonf3e)

	if err != nil {
		log.Printf("Error reading rows: " + err.Error())
		return stub.nf3e, err
	}
	return stub.nf3e, nil

}

func NewNf3eSituacaoRepositoryStub() Nf3eSituacaoRepositorioStub {
	return Nf3eSituacaoRepositorioStub{nf3e: model.Nf3eSituacao{
		Chnf3e:         "",
		Versao:         "",
		Tpamb:          "",
		Cstat:          "",
		Xmotivo:        "",
		Cuf:            "",
		Protnf3e:       "",
		Proceventonf3e: nil,
	}}
}

//func Consultar(chnf3e string) (Nf3esituacao, error) {
//
//	var resp Nf3esituacao
//
//	err := Cassandra.Session.Query("select chnf3e,versao,tpamb,cstat,xmotivo,cuf,protnf3e,proceventonf3e FROM nf3e_situacao where chnf3e = ?", chnf3e).
//		Scan(&resp.Chnf3e, &resp.Versao, &resp.Tpamb, &resp.Cstat, &resp.Xmotivo, &resp.Cuf, &resp.Protnf3e, &resp.Proceventonf3e)
//
//	if err != nil {
//		log.Printf("Error reading rows: " + err.Error())
//		return resp, err
//	}
//	return resp, err
//
//}
