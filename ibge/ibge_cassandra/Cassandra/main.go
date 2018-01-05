package Cassandra

import (
	"github.com/gocql/gocql"
	"fmt"
	"github.com/italiviocorrea/golang/commons"
	"strings"
	"log"
	"os"
)

var Session *gocql.Session

func init() {

	var err error

	log.Println("Inicializa variaveis de ambiente, caso nao exista")
	setVarEnv()

	log.Println("Iniciando as configuracoes da aplicacao.")
	commons.StartUp("CASSANDRA_API")

	hosts := strings.Split(commons.AppConfig.DBHost, ",")

	log.Println(hosts)

	cluster := gocql.NewCluster(hosts...)
	//cluster := gocql.NewCluster("cassandra-1:9042","cassandra-2:9042","cassandra-3:9042")
    //cluster.DisableInitialHostLookup = true
	cluster.ProtoVersion = 4

	cluster.Keyspace = commons.AppConfig.Database

	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexao com o Cassandra esta pronto.")
}

func setVarEnv() {

	if os.Getenv("CASSANDRA_API_SERVER") == "" {
		os.Setenv("CASSANDRA_API_SERVER", "0.0.0.0:8080")
	}

	if os.Getenv("CASSANDRA_API_DBHOST") == "" {
		os.Setenv("CASSANDRA_API_DBHOST", "localhost:19042,localhost:29042,localhost:39042")
	}

	if os.Getenv("CASSANDRA_API_DBUSER") == "" {
		os.Setenv("CASSANDRA_API_DBUSER", "admin")
	}

	if os.Getenv("CASSANDRA_API_DBPWD") == "" {
		os.Setenv("CASSANDRA_API_DBPWD ", "senha#123")
	}

	if os.Getenv("CASSANDRA_API_DATABASE") == "" {
		os.Setenv("CASSANDRA_API_DATABASE", "ibgeapi")
	}

	if os.Getenv("CASSANDRA_API_CONTEXT") == "" {
		os.Setenv("CASSANDRA_API_CONTEXT ", "/ibge/v3")
	}

	if os.Getenv("CASSANDRA_API_LOGLEVEL") == "" {
		os.Setenv("CASSANDRA_API_LOGLEVEL", "4")
	}

}
