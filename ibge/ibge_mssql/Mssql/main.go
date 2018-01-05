package Mssql

import (
	_ "github.com/denisenkom/go-mssqldb"
	"database/sql"
	"fmt"
	"log"
	"github.com/italiviocorrea/golang/commons"
	"os"
)

var Database *sql.DB

func init()  {

	log.Println("Inicializa variaveis de ambiente, caso nao exista")
	setVarEnv()

	log.Println("Iniciando as configuracoes da aplicacao.")
	commons.StartUp("MSSQL_API")

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
			commons.AppConfig.DBHost,
			commons.AppConfig.DBUser,
			commons.AppConfig.DBPwd,
			commons.AppConfig.DBPort,
			commons.AppConfig.Database)

	var err error

	// Create connection pool
	Database, err = sql.Open("sqlserver", connString)

	if err != nil {
		log.Fatal("Error creating connection pool:", err.Error())
	}

	log.Println("Connectado : "+connString)

}

func setVarEnv() {

	if os.Getenv("MSSQL_API_SERVER") == "" {
		os.Setenv("MSSQL_API_SERVER", "0.0.0.0:8080")
	}

	if os.Getenv("MSSQL_API_DBHOST") == "" {
		os.Setenv("MSSQL_API_DBHOST", "localhost")
	}

	if os.Getenv("MSSQL_API_DBPORT") == "" {
		os.Setenv("MSSQL_API_DBPORT", "1433")
	}

	if os.Getenv("MSSQL_API_DBUSER") == "" {
		os.Setenv("MSSQL_API_DBUSER", "sa")
	}

	if os.Getenv("MSSQL_API_DBPWD") == "" {
		os.Setenv("MSSQL_API_DBPWD ", "senha#123")
	}

	if os.Getenv("MSSQL_API_DATABASE") == "" {
		os.Setenv("MSSQL_API_DATABASE", "dbibgeapi")
	}

	if os.Getenv("MSSQL_API_CONTEXT") == "" {
		os.Setenv("MSSQL_API_CONTEXT ", "/ibge/v3")
	}

	if os.Getenv("MSSQL_API_LOGLEVEL") == "" {
		os.Setenv("MSSQL_API_LOGLEVEL", "4")
	}

}
