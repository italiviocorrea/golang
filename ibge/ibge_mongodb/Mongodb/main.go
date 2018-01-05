package Mongodb

import (
	"gopkg.in/mgo.v2"
	"time"
	"log"
	"github.com/italiviocorrea/golang/commons"
	"os"
	"fmt"
)

var Session *mgo.Session

func init() {

	var err error

	log.Println("Inicializa variaveis de ambiente, caso nao exista")
	setVarEnv()

	log.Println("Iniciando as configuracoes da aplicacao.")
	commons.StartUp("MONGODB_API")

	// Iniciando sessao do banco de dados MongoDB
	Session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{commons.AppConfig.DBHost},
		Username: commons.AppConfig.DBUser,
		Password: commons.AppConfig.DBPwd,
		Timeout:  60 * time.Second,
	})

	if err != nil {
		panic(err)
	}

	// Criando indices
	addIndexes()

	fmt.Println("Conexao com o MongoDB esta pronto.")
}

// Add indexes into MongoDB
func addIndexes() {
	var err error

	ufIndex1 := mgo.Index{
		Key:        []string{"codigo"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}
	municipioIndex1 := mgo.Index{
		Key:        []string{"codigo"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}
	// Add indexes into MongoDB
	session := Session.Copy()
	defer session.Close()
	ufCol := session.DB(commons.AppConfig.Database).C("ufs")
	municipioCol := session.DB(commons.AppConfig.Database).C("municipios")

	// cria indice codigo para UF
	err = ufCol.EnsureIndex(ufIndex1)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}
	log.Println("Indice para UF criado com sucesso")

	// cria indice codigo para Municipio
	err = municipioCol.EnsureIndex(municipioIndex1)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}
	log.Println("Indice para Municipio criado com sucesso")

}

// DataStore for MongoDB
type DataStore struct {
	MongoSession *mgo.Session
}

// Close closes a mgo.Session value.
// Used to add defer statements for closing the copied session.
func (ds *DataStore) Close() {
	ds.MongoSession.Close()
}

// Collection returns mgo.collection for the given name
func (ds *DataStore) Collection(name string) *mgo.Collection {
	return ds.MongoSession.DB(commons.AppConfig.Database).C(name)
}

// NewDataStore creates a new DataStore object to be used for each HTTP request.
func NewDataStore() *DataStore {
	session := Session.Copy()
	dataStore := &DataStore{
		MongoSession: session,
	}
	return dataStore
}

func setVarEnv() {

	if os.Getenv("MONGODB_API_SERVER") == "" {
		os.Setenv("MONGODB_API_SERVER", "0.0.0.0:8080")
	}

	if os.Getenv("MONGODB_API_DBHOST") == "" {
		os.Setenv("MONGODB_API_DBHOST", "localhost")
	}

	if os.Getenv("MONGODB_API_DBUSER") == "" {
		os.Setenv("MONGODB_API_DBUSER", "")
	}

	if os.Getenv("MONGODB_API_DBPWD") == "" {
		os.Setenv("MONGODB_API_DBPWD ", "")
	}

	if os.Getenv("MONGODB_API_DATABASE") == "" {
		os.Setenv("MONGODB_API_DATABASE", "ibgeapi")
	}

	if os.Getenv("MONGODB_API_CONTEXT") == "" {
		os.Setenv("MONGODB_API_CONTEXT ", "/ibge/v3")
	}

	if os.Getenv("MONGODB_API_LOGLEVEL") == "" {
		os.Setenv("MONGODB_API_LOGLEVEL", "4")
	}

}
