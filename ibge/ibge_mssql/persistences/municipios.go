package persistences

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/italiviocorrea/golang/commons"
	"github.com/italiviocorrea/golang/ibge/ibge_mssql/Mssql"
	"github.com/italiviocorrea/golang/ibge/models"
	"log"
	"strconv"
)

func CreateMunicipio(municipio models.Municipio) (int64, error) {

	ctx := context.Background()

	// Check if dbMunicipio is alive.
	err := Mssql.Database.PingContext(ctx)

	if err != nil {
		log.Printf("Error pinging Municipio: " + err.Error())
	}

	tsql := fmt.Sprintf("INSERT INTO municipios (codigo, nome, uf_codigo) VALUES (@Codigo, @Nome, @UfCodigo);")

	// Execute non-query with named parameters
	result, err := Mssql.Database.ExecContext(
		ctx,
		tsql,
		sql.Named("UfCodigo", municipio.Uf.Codigo),
		sql.Named("Nome", municipio.Nome),
		sql.Named("Codigo", municipio.Codigo))

	if err != nil {
		log.Printf("Error pinging dbMunicipio: " + err.Error())
		return -1, err
	}

	return result.RowsAffected()
}

func UpdateMunicipio(municipio models.Municipio) (int64, error) {

	ctx := context.Background()

	// Check if dbMunicipio is alive.
	err := Mssql.Database.PingContext(ctx)
	if err != nil {
		log.Printf("Error pinging Municipio: " + err.Error())
	}

	tsql := fmt.Sprintf("UPDATE municipios SET uf_codigo = @UfCodigo, nome = @Nome WHERE codigo= @Codigo")

	// Execute non-query with named parameters
	result, err := Mssql.Database.ExecContext(
		ctx,
		tsql,
		sql.Named("UfCodigo", municipio.Uf.Codigo),
		sql.Named("Nome", municipio.Nome),
		sql.Named("Codigo", municipio.Codigo))

	if err != nil {
		log.Printf("Error updating row: " + err.Error())
		return -1, err
	}

	return result.RowsAffected()
}

func DeleteMunicipio(codigo int64) (int64, error) {

	ctx := context.Background()

	// Check if dbMunicipio is alive.
	err := Mssql.Database.PingContext(ctx)
	if err != nil {
		log.Printf("Error pinging dbMunicipio: " + err.Error())
	}

	tsql := fmt.Sprintf("DELETE FROM municipios WHERE codigo=@Codigo;")

	// Execute non-query with named parameters
	result, err := Mssql.Database.ExecContext(ctx, tsql, sql.Named("Codigo", codigo))
	if err != nil {
		fmt.Printf("Error deleting row: " + err.Error())
		return -1, err
	}

	return result.RowsAffected()
}

func GetAllMunicipio(page_num int, page_size int) []models.MunicipiosResponse {

	skips := page_size * (page_num - 1)

	ctx := context.Background()

	// Check if dbMunicipio is alive.
	err := Mssql.Database.PingContext(ctx)
	if err != nil {
		log.Printf("Error pinging Municipio: " + err.Error())
	}

	tsql := fmt.Sprintf("SELECT mnc.codigo, mnc.nome, uf.sigla"+
		" FROM municipios mnc INNER JOIN ufs uf ON mnc.uf_codigo = uf.codigo"+
		" ORDER BY mnc.codigo OFFSET %d ROWS FETCH NEXT %d ROWS ONLY;", skips, page_size)

	// Execute query
	rows, err := Mssql.Database.QueryContext(ctx, tsql)
	if err != nil {
		log.Printf("Error reading rows: " + err.Error())
		return nil
	}

	defer rows.Close()

	municipios := []models.MunicipiosResponse{}

	// Iterate through the result set.
	for rows.Next() {
		municipio := models.MunicipiosResponse{}

		// Get values from row.
		err := rows.Scan(&municipio.Codigo, &municipio.Nome, &municipio.Uf)
		if err != nil {
			log.Printf("Error reading rows: " + err.Error())
			return nil
		}

		// adiciona os links
		links := []commons.Link{commons.Link{Name: "self", Method: "GET", Href: commons.AppConfig.Context + "/municipios/" + strconv.FormatInt(municipio.Codigo, 10)}}
		municipio.Links = links

		municipios = append(municipios, municipio)
	}

	return municipios
}

func GetMunicipioByCode(codigo int64) (models.MunicipioResponse, error) {

	ctx := context.Background()

	// Check if dbMunicipio is alive.
	err := Mssql.Database.PingContext(ctx)

	if err != nil {
		log.Printf("Error pinging Municipio: " + err.Error())
	}

	tsql := fmt.Sprintf("SELECT mnc.codigo, mnc.nome, mnc.uf_codigo, uf.nome, uf.sigla" +
		" FROM municipios mnc INNER JOIN ufs uf ON mnc.uf_codigo = uf.codigo" +
		" WHERE mnc.codigo=@Codigo;")

	// Execute query
	var b models.MunicipioResponse

	err = Mssql.Database.QueryRowContext(ctx, tsql, sql.Named("Codigo", codigo)).Scan(&b.Codigo, &b.Nome, &b.Uf.Codigo, &b.Uf.Nome, &b.Uf.Sigla)

	if err != nil {
		log.Printf("Error reading rows: " + err.Error())
		return b, err
	}

	// adiciona os links
	links := []commons.Link{commons.Link{Name: "self", Method: "GET", Href: commons.AppConfig.Context + "/municipios/" + strconv.FormatInt(b.Codigo, 10)},
		commons.Link{Name: "update", Method: "PUT", Href: commons.AppConfig.Context + "/municipios/" + strconv.FormatInt(b.Codigo, 10)},
		commons.Link{Name: "remove", Method: "DELETE", Href: commons.AppConfig.Context + "/municipios/" + strconv.FormatInt(b.Codigo, 10)}}
	b.Links = links
	// retorna a resposta
	return b, err
}

func GetMunicipioCountPage(page_size int) int {

	ctx := context.Background()

	// Check if dbMunicipio is alive.
	err := Mssql.Database.PingContext(ctx)

	if err != nil {
		log.Printf("Error pinging Municipio: " + err.Error())
	}

	tsql := fmt.Sprintf("SELECT count(*) FROM municipios;")

	var total int

	// fazer a busca no banco de dados
	err = Mssql.Database.QueryRowContext(ctx, tsql).Scan(&total)

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
