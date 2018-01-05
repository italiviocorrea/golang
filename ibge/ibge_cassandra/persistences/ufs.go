package persistences

import (
	"com/ItalivioCorrea/commons"
	"com/ItalivioCorrea/ibge/ibge_cassandra/Cassandra"
	"log"
	"fmt"
	"strconv"
	"com/ItalivioCorrea/ibge/models"
)

func CreateUF(uf models.Uf)  (int64, error) {

	err := Cassandra.Session.Query("INSERT INTO ufs (codigo, nome, sigla) VALUES (?,?,?)",uf.Codigo,uf.Nome,uf.Sigla).Exec()

	if err != nil {
		log.Printf("Erro ao inserir UF: " + err.Error())
		return -1, err
	}

	return 1, nil
}

func UpdateUF(uf models.Uf) (int64, error) {

	err := Cassandra.Session.Query("UPDATE ufs SET sigla = ?, nome = ? WHERE codigo= ?",uf.Sigla,uf.Nome,uf.Codigo).Exec()


	if err != nil {
		log.Printf("Erro ao atualizar UF: " + err.Error())
		return -1, err
	}

	return 1, nil
}

func DeleteUF(uf models.Uf) (int64, error) {

	err := Cassandra.Session.Query("delete from ufs where codigo = ?",uf.Codigo).Exec()

	if err != nil {
		fmt.Printf("Erro ao excluir a UF: " + err.Error())
		return -1, err
	}

	return 1, nil
}

func GetAllUF(page_num int, page_size int) []models.UFResponse {

	//skips := page_size * (page_num - 1)


	rows := Cassandra.Session.Query("select codigo,nome,sigla from ufs").Iter()

	ufs := []models.UFResponse{}
	uf := models.UFResponse{}

	// Iterate through the result set.
	for rows.Scan(&uf.Codigo,&uf.Nome,&uf.Sigla) {

		// adiciona os links
		links := []commons.Link{commons.Link{Name: "self", Method: "GET", Href: commons.AppConfig.Context+"/ufs/" + strconv.Itoa(uf.Codigo)}}
		uf.Links = links

		ufs = append(ufs, uf)
	}

	return ufs
}

func GetByCode(uf models.Uf) (models.UFResponse, error) {


	// Execute query
	// fazer a busca no banco de dados
	var b models.UFResponse

	err := Cassandra.Session.Query("select codigo,nome,sigla from ufs where codigo = ?",uf.Codigo).Scan(&b.Codigo,&b.Nome,&b.Sigla)

	if err != nil {
		log.Printf("Error reading rows: " + err.Error())
		return b, err
	}

	// adiciona os links
	links := []commons.Link{commons.Link{Name: "self", Method: "GET", Href: commons.AppConfig.Context+"/ufs/" + strconv.Itoa(b.Codigo)},
		commons.Link{Name: "update", Method: "PUT", Href: commons.AppConfig.Context+"/ufs/" + strconv.Itoa(b.Codigo)},
		commons.Link{Name: "remove", Method: "DELETE", Href: commons.AppConfig.Context+"/ufs/" + strconv.Itoa(b.Codigo)}}
	b.Links = links
	// retorna a resposta
	return b, err
}

func GetUFCountPage(page_size int) int {


	var total int
	// fazer a busca no banco de dados
	err := Cassandra.Session.Query("select count(*) from ufs").Scan(&total)

	if err != nil {
		total = 1
	}
	if total > page_size {
		total = total / page_size
	} else {
		total = 1
	}
	return total
}
