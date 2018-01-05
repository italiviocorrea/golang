package persistences

import (
	"github.com/italiviocorrea/golang/commons"
	"github.com/italiviocorrea/golang/ibge/models"
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"github.com/italiviocorrea/golang/ibge/ibge_mssql/Mssql"
)


func CreateUF(uf models.Uf) (int64, error) {

	ctx := context.Background()

	// Check if database is alive.
	err := Mssql.Database.PingContext(ctx)

	if err != nil {
		log.Printf("Error pinging database: " + err.Error())
	}

	tsql := fmt.Sprintf("INSERT INTO ufs (codigo, nome, sigla) VALUES (@Codigo, @Nome, @Sigla);")

	// Execute non-query with named parameters
	result, err := Mssql.Database.ExecContext(
		ctx,
		tsql,
		sql.Named("Sigla", uf.Sigla),
		sql.Named("Nome", uf.Nome),
		sql.Named("Codigo", uf.Codigo))

	if err != nil {
		log.Printf("Error pinging database: " + err.Error())
		return -1, err
	}

	return result.RowsAffected()
}

func UpdateUF(uf models.Uf) (int64, error) {

	ctx := context.Background()

	// Check if database is alive.
	err := Mssql.Database.PingContext(ctx)

	if err != nil {
		log.Printf("Error pinging database: " + err.Error())
	}

	tsql := fmt.Sprintf("UPDATE ufs SET sigla = @Sigla, nome = @Nome WHERE codigo= @Codigo")

	// Execute non-query with named parameters
	result, err := Mssql.Database.ExecContext(
		ctx,
		tsql,
		sql.Named("Sigla", uf.Sigla),
		sql.Named("Nome", uf.Nome),
		sql.Named("Codigo", uf.Codigo))

	if err != nil {
		log.Printf("Error updating row: " + err.Error())
		return -1, err
	}

	return result.RowsAffected()
}

func DeleteUF(codigo int) (int64, error) {

	ctx := context.Background()

	// Check if database is alive.
	err := Mssql.Database.PingContext(ctx)

	if err != nil {
		log.Printf("Error pinging database: " + err.Error())
	}

	tsql := fmt.Sprintf("DELETE FROM ufs WHERE codigo=@Codigo;")

	// Execute non-query with named parameters
	result, err := Mssql.Database.ExecContext(ctx, tsql, sql.Named("Codigo", codigo))

	if err != nil {
		fmt.Printf("Error deleting row: " + err.Error())
		return -1, err
	}

	return result.RowsAffected()
}

func GetAllUF(page_num int, page_size int) []models.UFResponse {

	skips := page_size * (page_num - 1)

	ctx := context.Background()

	// Check if database is alive.
	err := Mssql.Database.PingContext(ctx)

	if err != nil {
		log.Printf("Error pinging database: " + err.Error())
	}

	tsql := fmt.Sprintf("SELECT * FROM ufs ORDER BY codigo OFFSET %d ROWS FETCH NEXT %d ROWS ONLY;", skips, page_size)

	// Execute query
	rows, err := Mssql.Database.QueryContext(ctx, tsql)

	if err != nil {
		log.Printf("Error reading rows: " + err.Error())
		return nil
	}

	defer rows.Close()

	ufs := []models.UFResponse{}

	// Iterate through the result set.
	for rows.Next() {
		uf := models.UFResponse{}

		// Get values from row.
		err := rows.Scan(&uf.Codigo, &uf.Nome, &uf.Sigla)

		if err != nil {
			log.Printf("Error reading rows: " + err.Error())
			return nil
		}

		// adiciona os links
		links := []commons.Link{commons.Link{Name: "self", Method: "GET", Href: commons.AppConfig.Context+"/ufs/" + strconv.Itoa(uf.Codigo)}}
		uf.Links = links

		ufs = append(ufs, uf)
	}

	return ufs
}

func GetUFByCode(codigo int) (models.UFResponse, error) {

	ctx := context.Background()

	// Check if database is alive.
	err := Mssql.Database.PingContext(ctx)

	if err != nil {
		log.Printf("Error pinging database: " + err.Error())
	}

	tsql := fmt.Sprintf("SELECT * FROM ufs WHERE codigo=@Codigo;")

	// Execute query
	var b models.UFResponse

	err = Mssql.Database.QueryRowContext(ctx, tsql, sql.Named("Codigo", codigo)).Scan(&b.Codigo, &b.Nome, &b.Sigla)

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

	ctx := context.Background()

	// Check if database is alive.
	err := Mssql.Database.PingContext(ctx)

	if err != nil {
		log.Printf("Error pinging database: " + err.Error())
	}

	tsql := fmt.Sprintf("SELECT count(*) FROM ufs;")

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
