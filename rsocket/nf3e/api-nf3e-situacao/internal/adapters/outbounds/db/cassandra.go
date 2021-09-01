package db

import (
	"github.com/italiviocorrea/golang/rsocket/nf3e/api-nf3e-situacao/configs"
	"github.com/rs/zerolog/log"
	"github.com/yugabyte/gocql"
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

	cluster := gocql.NewCluster(hosts...)

	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: configs.Get().DBUser,
		Password: configs.Get().DBPwd,
	}

	cluster.Timeout = 12 * time.Second
	cluster.Keyspace = configs.Get().Database
	//cluster.PoolConfig.HostSelectionPolicy = TokenAwareHostPolicy(RoundRobinHostPolicy())
	//cluster.ReconnectInterval = 1 * time.Second

	db, err = cluster.CreateSession()

	if err != nil {
		//log.Println(fmt.Sprintf("Error to loading Database %s", err))
		log.Fatal().Err(err).
			Str("service", "api-nf3e-situacao").
			Str("component", "cassandra").
			Str("hosts", configs.Get().DBHost).
			Str("keyspace", configs.Get().Database).
			Msg("Erro ao carregar o banco de dados cassandra.")
		return nil
	}

	log.Info().
		Str("service", "api-nf3e-situacao").
		Str("component", "cassandra").
		Msg("O banco de dados foi conectado.")

	return &cassandraDB{
		db: db,
	}
}

func (c cassandraDB) DB() *gocql.Session {
	return c.db
}
