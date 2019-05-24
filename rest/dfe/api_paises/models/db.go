package models

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"log"
	"os"
	"strconv"
	"time"
)

type (

	EventoResposta struct {
		Operacao string    `json:"operacao"        description:"Operacao informa o tipo de evento que gerou a modificacao."`
		DataHora time.Time `json:"dataHora"        description:"Data e Hora em que ocorreu o evento" `
		Nsu      int       `json:"nsu"             description:"Data e Hora em que ocorreu o evento" `
		Antes    Paises    `json:"antes,omitempty" description:"Dados do Pais antes da operacao" `
		Apos     Paises    `json:"apos,omitempty"  description:"Dados do Pais após a operacao" `
	}

	Mensagem struct {
		Codigo   int       `json:"codigo"   description:"codigo do pais"`
		Mensagem string    `json:"mensagem" description:"Descricao da mensagem"`
		DataHora time.Time `json:"dataHora" description:"Número serial unico do evento" `
	}

	Link struct {
		Nome   string `json:"nome,omitempty"   description:"nome do link"`
		Metodo string `json:"metodo" description:"nome do método HTTP"`
		Tipo   string `json:"tipo,omitempty"   description:"Tipo de midia esperado pelo link"`
		Href   string `json:"href"   description:"URI do link"`
		Rel    string `json:"rel"    description:"Permite informar a relacao do link com o recurso"`
	}

	Paginacao struct {
		Primeira     string `json:"primeira,omitempty"     description:"Caminho para o primeira pagina."`
		Proxima      string `json:"proxima,omitempty"      description:"Caminho para a proxima pagina, em relacao a atual."`
		Anterior     string `json:"anterior,omitempty"     description:"Caminho para a pagina anterior, em relacao a atual"`
		Ultima       string `json:"ultima,omitempty"       description:"Caminho para a ultima pagina."`
		TotalPaginas int    `json:"totalPaginas,omitempty" description:"Total de paginas disponiveis para listagem."`
	}

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
		os.Setenv("API_SRV_ADDR", "0.0.0.0:8080")
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

/**
	Gera o link create.
 */
func CreateLink() Link {

	var context = os.Getenv("API_SRV_CONTEXT")

	var link = Link{}

	link.Nome = "create"
	link.Metodo = "POST"
	link.Tipo = "application/json"
	link.Href = context+"/paises"
	link.Rel = "Novo"

	return link

}

/**
	Gera o link litar todos.
 */
func ListAllLink() Link {

	var context = os.Getenv("API_SRV_CONTEXT")

	var link = Link{}

	link.Nome = "ListarTodos"
	link.Metodo = "GET"
	link.Tipo = "application/json"
	link.Href = context+"/paises?pagina=1&limite=50"
	link.Rel = "Listar"

	return link

}
