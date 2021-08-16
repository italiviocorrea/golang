package Cassandra

import (
	"fmt"
	"github.com/italiviocorrea/golang/rsocket/nf3e/cmd/cons-sit/util"
	"github.com/yugabyte/gocql"
	"log"
	"os"
	"strings"
	"time"
)

var Session *gocql.Session

func init() {

	log.Println("Inicializa variaveis de ambiente, caso nao exista")
	setVarEnv()

	log.Println("Iniciando as configuracoes da aplicacao.")
	util.StartUp("NF3E_API")

	hosts := strings.Split(util.AppConfig.DBHost, ",")

	log.Println(hosts)

	cluster := gocql.NewCluster(hosts...)

	cluster.Timeout = 12 * time.Second
	cluster.Keyspace = util.AppConfig.Database

	Session, _ = cluster.CreateSession()
	fmt.Println("Conexao com o Cassandra esta pronto.")
}

func setVarEnv() {

	if os.Getenv("NF3E_API_SERVER") == "" {
		os.Setenv("NF3E_API_SERVER", ":7878")
	}

	if os.Getenv("NF3E_API_DBHOST") == "" {
		os.Setenv("NF3E_API_DBHOST", "s1602.ms:9042")
	}

	if os.Getenv("NF3E_API_DBUSER") == "" {
		os.Setenv("NF3E_API_DBUSER", "admin")
	}

	if os.Getenv("NF3E_API_DBPWD") == "" {
		os.Setenv("NF3E_API_DBPWD ", "senha#123")
	}

	if os.Getenv("NF3E_API_DATABASE") == "" {
		os.Setenv("NF3E_API_DATABASE", "nf3e")
	}

	if os.Getenv("NF3E_API_CONTEXT") == "" {
		os.Setenv("NF3E_API_CONTEXT ", "/nf3/v1")
	}

	if os.Getenv("NF3E_API_LOGLEVEL") == "" {
		os.Setenv("NF3E_API_LOGLEVEL", "4")
	}

}
