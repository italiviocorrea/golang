package model

import (
	"github.com/gocql/gocql"
	"github.com/italiviocorrea/golang/grpc/nfcexml/Cassandra"
	"log"
)

type Nfcexml struct {
	ID       gocql.UUID
	DadosXML []byte
}

func (t Nfcexml) Validate() error {

	return nil
}

/**
Grava o XML
*/
func (t *Nfcexml) Create() (int64, error) {

	err := Cassandra.Session.Query("INSERT INTO nfcexml (id, dadosXML) VALUES (?,?)",
		gocql.TimeUUID(), t.DadosXML).Exec()

	if err != nil {
		log.Printf("Erro ao inserir XML: " + err.Error())
		return -1, err
	}

	return 1, nil

}

/**
Exclui o XML correspondente ao ID
*/
func (t *Nfcexml) Delete() error {

	if err := Cassandra.Session.Query("delete from ks_nfce.nfcexml where id = ?", t.ID).Exec(); err != nil {
		log.Printf("Erro ao remover o país : " + err.Error())
		return err
	}

	return nil
}

/**
Retorna o XML correspondente ao ID informado
*/
func (t *Nfcexml) Find() error {

	if err := Cassandra.Session.Query("select id,dadosXML from ks_nfce.nfcexml where id = ?", t.ID).Scan(&t.ID, &t.DadosXML); err != nil {
		log.Println("Erro ao pesquisar pais :" + err.Error())
		return err
	}
	return nil
}

/**
Retorna todos XMLs da NFCe
*/
func (t Nfcexml) FindAll() []Nfcexml {

	var nfceList []Nfcexml
	var nfcexml Nfcexml

	rows := Cassandra.Session.Query("select id,dadosXML from ks_nfce.nfcexml").Iter()

	for rows.Scan(&nfcexml.ID, &nfcexml.DadosXML) {
		log.Println(nfcexml.DadosXML)
		nfceList = append(nfceList, nfcexml)
		nfcexml = Nfcexml{}
	}

	return nfceList

}
