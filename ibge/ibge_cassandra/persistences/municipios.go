package persistences

import (
"com/ItalivioCorrea/commons"
	"log"
	"strconv"
	"com/ItalivioCorrea/ibge/ibge_cassandra/Cassandra"
	"fmt"
	"com/ItalivioCorrea/ibge/models"
)


func CreateMunicipio(municipio models.Municipio) (int64, error) {

	err := Cassandra.Session.Query("INSERT INTO municipios (codigo, nome, uf) VALUES (?,?,{codigo: ?, nome: ?, sigla: ?})",
		municipio.Codigo,municipio.Nome,municipio.Uf.Codigo, municipio.Uf.Nome, municipio.Uf.Sigla).Exec()

	if err != nil {
		log.Printf("Erro ao inserir municipio: " + err.Error())
		return -1, err
	}

	return 1, nil
}

func UpdateMunicipio(municipio models.Municipio)  (int64, error) {

	err := Cassandra.Session.Query("UPDATE municipios SET nome = ?, uf = {codigo: ?, nome: ?, sigla: ?} WHERE codigo= ?",
		municipio.Nome,municipio.Uf.Codigo,municipio.Uf.Nome,municipio.Uf.Sigla,municipio.Codigo).Exec()


	if err != nil {
		log.Printf("Erro ao atualizar municipio: " + err.Error())
		return -1, err
	}

	return 1, nil
}

func DeleteMunicipio(municipio models.Municipio)  (int64, error) {

	err := Cassandra.Session.Query("delete from municipios where codigo = ?",municipio.Codigo).Exec()

	if err != nil {
		fmt.Printf("Erro ao excluir a municipio: " + err.Error())
		return -1, err
	}

	return 1, nil
}

func GetAllMunicipio(page_num int, page_size int) []models.MunicipiosResponse {

	//skips := page_size * (page_num - 1)


	rows := Cassandra.Session.Query("select codigo,nome,uf.sigla from municipios").Iter()

	municipios := []models.MunicipiosResponse{}
	municipio := models.MunicipiosResponse{}

	// Iterate through the result set.
	for rows.Scan(&municipio.Codigo,&municipio.Nome,&municipio.Uf) {

		// adiciona os links
		links := []commons.Link{commons.Link{Name: "self", Method: "GET", Href: commons.AppConfig.Context+"/municipios/" + strconv.FormatInt(municipio.Codigo,10)}}
		municipio.Links = links

		municipios = append(municipios, municipio)
	}

	return municipios
}

func GetMunicipioByCode(municipio models.Municipio) (models.MunicipioResponse, error) {


	// Execute query
	// fazer a busca no banco de dados
	var b models.MunicipioResponse

	err := Cassandra.Session.Query("select codigo,nome,uf.codigo,uf.nome,uf.sigla from municipios where codigo = ?",municipio.Codigo).Scan(&b.Codigo,&b.Nome,&b.Uf.Codigo,&b.Uf.Nome,&b.Uf.Sigla)

	if err != nil {
		log.Printf("Error reading rows: " + err.Error())
		return b, err
	}

	// adiciona os links
	links := []commons.Link{commons.Link{Name: "self", Method: "GET", Href: commons.AppConfig.Context+"/municipios/" + strconv.FormatInt(b.Codigo,10)},
		commons.Link{Name: "update", Method: "PUT", Href: commons.AppConfig.Context+"/municipios/" + strconv.FormatInt(b.Codigo,10)},
		commons.Link{Name: "remove", Method: "DELETE", Href: commons.AppConfig.Context+"/municipios/" + strconv.FormatInt(b.Codigo,10)}}
	b.Links = links
	// retorna a resposta
	return b, err
}

func GetMunicipioCountPage(page_size int) int {


	var total int
	// fazer a busca no banco de dados
	err := Cassandra.Session.Query("select count(*) from municipios").Scan(&total)

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