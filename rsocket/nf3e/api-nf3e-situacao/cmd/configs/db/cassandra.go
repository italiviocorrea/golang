package db

import (
	"fmt"
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/cmd/configs"
	"github.com/yugabyte/gocql"
	"log"
	"strings"
	"time"
)

type ClientDB interface {
	DB() *gocql.Session
}

type cassandraDB struct {
	db *gocql.Session
}

func NewCassandraClient() ClientDB {

	var err error
	var db *gocql.Session

	hosts := strings.Split(configs.Get().DBHost, ",")

	log.Println(hosts)

	cluster := gocql.NewCluster(hosts...)

	cluster.Timeout = 12 * time.Second
	cluster.Keyspace = configs.Get().Database

	db, err = cluster.CreateSession()

	if err != nil {
		log.Println(fmt.Sprintf("Error to loading Database %s", err))
		return nil
	}

	fmt.Println("Database was connected.")

	return &cassandraDB{
		db: db,
	}
}

func (c cassandraDB) DB() *gocql.Session {
	return c.db
}
