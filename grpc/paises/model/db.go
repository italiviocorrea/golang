package model

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"log"
	"os"
	"strconv"
)


var db *gorm.DB


/**
Conecta com o banco de dados
*/
func ConnectToDatabase() *gorm.DB {

	setVarEnv()

	var err error

	db, err = gorm.Open("mssql",connectionString())

	setPool()

	if err != nil {
		log.Fatal("Error connection database:", err.Error())
	}

	log.Println("Connectado ao banco de dados")

	return db

}

/**
Define o pool de conexao do banco de dados
*/
func setPool() {


	if i, err :=  strconv.Atoi(os.Getenv("API_DB_MINPOOL")); err == nil {
		db.DB().SetMaxIdleConns(i)
	}

	if i, err :=  strconv.Atoi(os.Getenv("API_DB_MAXPOOL")); err == nil {
		db.DB().SetMaxOpenConns(i)
	}


}

// Cria a string de conexao
func connectionString() string  {

	log.Println("Inicializa variaveis de ambiente, caso nao exista")
	setVarEnv()

	log.Println("Iniciando as configuracoes da aplicacao.")

	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		os.Getenv("API_DB_USER"),
		os.Getenv("API_DB_PWD"),
		os.Getenv("API_DB_HOST"),
		os.Getenv("API_DB_PORT"),
		os.Getenv("API_DB_NAME")	)

	log.Println(connString)

	return connString
}

// Fixa as variaveis de ambiente
func setVarEnv() {

	if os.Getenv("API_SRV_ADDR") == "" {
		os.Setenv("API_SRV_ADDR", "0.0.0.0:50051")
	}

	if os.Getenv("API_DB_HOST") == "" {
		os.Setenv("API_DB_HOST", "localhost")
	}

	if os.Getenv("API_DB_PORT") == "" {
		os.Setenv("API_DB_PORT", "1433")
	}

	if os.Getenv("API_DB_USER") == "" {
		os.Setenv("API_DB_USER", "sa")
	}

	if os.Getenv("API_DB_PWD") == "" {
		os.Setenv("API_DB_PWD", "Senha123")
	}

	if os.Getenv("API_DB_NAME") == "" {
		os.Setenv("API_DB_NAME", "dbpaises")
	}

	if os.Getenv("API_DB_MINPOOL") == "" {
		os.Setenv("API_DB_MINPOOL", "10")
	}

	if os.Getenv("API_DB_MAXPOOL") == "" {
		os.Setenv("API_DB_MAXPOOL", "128")
	}

}

