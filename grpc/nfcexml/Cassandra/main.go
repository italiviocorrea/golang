package Cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"os"
	"strings"
	"time"
)

var Session *gocql.Session

func init() {

	var err error

	log.Println("Inicializa variaveis de ambiente, caso nao exista")
	setVarEnv()

	log.Println("Iniciando as configuracoes da aplicacao.")

	hosts := strings.Split(os.Getenv("API_DB_HOST"), ",")

	log.Println(hosts)

	cluster := gocql.NewCluster(hosts...)
	//cluster := gocql.NewCluster("cassandra-1:9042","cassandra-2:9042","cassandra-3:9042")
	//cluster.DisableInitialHostLookup = true
	cluster.ProtoVersion = 4

	cluster.ConnectTimeout = time.Second * 10

	//cluster.Authenticator = gocql.PasswordAuthenticator{Username: os.Getenv("API_DB_USER"),
	//	Password: os.Getenv("API_DB_PWD")} //replace the username and password fields with their real settings.

	cluster.Keyspace = os.Getenv("API_DB_NAME")
	cluster.Consistency = gocql.Quorum

	Session, err = cluster.CreateSession()

	if err != nil {
		panic(err)
	}

	//createDBandTable()

	fmt.Println("Conexao com o Cassandra esta pronto.")

}

func createDBandTable() {

	// create keyspaces
	err := Session.Query("CREATE KEYSPACE IF NOT EXISTS ks_nfce WITH REPLICATION = {'class' : 'SimpleStrategy', 'replication_factor' : 1};").Exec()
	if err != nil {
		log.Println(err)
		return
	}

	// create table
	err = Session.Query("CREATE TABLE IF NOT EXISTS ks_nfce.nfcexml (id UUID, dadosXML blob, PRIMARY KEY (id));").Exec()
	if err != nil {
		log.Println(err)
		return
	}

}
func setVarEnv() {

	if os.Getenv("API_SRV_ADDR") == "" {
		os.Setenv("API_SRV_ADDR", "0.0.0.0:50051")
	}

	if os.Getenv("API_DB_HOST") == "" {
		os.Setenv("API_DB_HOST", "localhost:19042,localhost:29042,localhost:39042")
	}

	if os.Getenv("API_DB_USER") == "" {
		os.Setenv("API_DB_USER", "admin")
	}

	if os.Getenv("API_DB_PWD") == "" {
		os.Setenv("API_DB_PWD ", "senha#123")
	}

	if os.Getenv("API_DB_NAME") == "" {
		os.Setenv("API_DB_NAME", "ks_nfce")
	}

}
