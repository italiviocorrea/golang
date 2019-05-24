package endpoints

import (
	"../models"
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)


/*
  Configuração Default da rota.
*/
func Startup() *gin.Engine {

	setVarEnv()

	// iniciar as rotas para os recursos
	router := gin.New()

	// Rota estatica para a documentação do swagger
	router.Use(static.Serve(os.Getenv("API_DOC_CONTEXT"), static.LocalFile("./swagger", true)))

	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())

	// Rotas para os endpoint da API
	r1 := router.Group(os.Getenv("API_SRV_CONTEXT"))
	{
		r1.GET("/paises", getPaisesAll)
		r1.GET("/paises/:codigo", getPaisesByCode)
		r1.POST("/paises", createPaises)
		r1.PUT("/paises/:codigo", updatePaises)
		r1.DELETE("/paises/:codigo", deletePaises)
	}

	// Rotas para os endpoints de monitoramento
	r2 := router.Group(os.Getenv("API_MON_CONTEXT"))
	{
		r2.GET("/metrics", gin.WrapH(promhttp.Handler()))
		r2.GET("/status", getStatus)
	}

	// Rotas para os endpoints de eventos.
	r3 := router.Group(os.Getenv("API_EVT_CONTEXT"))
	{
		r3.GET("/paises/:nsu", getEventosPaises)
	}

	return router
}

// Handlers Functions

/**
	Função Handler (manipulador) da requisição HTTP/HTTPS, para atender a solicitação de consulta
	aos eventos streaming de manutenção da tabela de países. Este serviço retornará, a partir de um offset, todas as
	inclusões, alterações e exclusões de registros ocorridas na tabela de Países. Todos eventos serão mantidos por um ano,
	nos servidores Kafka.
 */
func getEventosPaises(c *gin.Context) {

	// declara as variáveis de resposta e mensagem.
	var mensagem models.Mensagem
	var resposta models.PaisesResposta

	// Gera a mensagem
	mensagem.Codigo = http.StatusOK
	mensagem.Mensagem = "Eventos encontrados com sucesso! "
	mensagem.DataHora = time.Now()

	// Gera a resposta
	resposta.Mensagens = append(resposta.Mensagens, mensagem)

	log.Println("Lendo o log de eventos da tabela países.")
	c.JSON(http.StatusOK, resposta)

}

/**
	Handler/Manipulador de requisições para criar/inserir (POST) um país na tabela países.
 */
func createPaises(c *gin.Context) {

	// declara as variáveis
	var pais models.Paises
	var resposta models.PaisesResposta
	var mensagem models.Mensagem

	// versao da resposta
	resposta.Versao = os.Getenv("API_RSP_VERSION")

	// Bind o payload no objeto país
	if err := c.ShouldBindJSON(&pais); err != nil {

		mensagem.Codigo = http.StatusBadRequest
		mensagem.Mensagem = "Dados de entrada inválido! " + err.Error()
		mensagem.DataHora = time.Now()

		resposta.Mensagens = append(resposta.Mensagens, mensagem)

		c.JSONP(http.StatusBadRequest, resposta)
		return
	}

	// Valida os dados do país
	if err := pais.Validate(); err != nil {

		mensagem.Codigo = http.StatusBadRequest
		mensagem.Mensagem = "Erro na validação dos dados : " + err.Error()
		mensagem.DataHora = time.Now()

		resposta.Mensagens = append(resposta.Mensagens, mensagem)

		c.JSONP(http.StatusBadRequest, resposta)
		return

	}

	// Salva o novo país
	if _, err := pais.Create(); err != nil {

		mensagem.Codigo = http.StatusInternalServerError
		mensagem.Mensagem = "Não foi possível gravar o país : " + err.Error()
		mensagem.DataHora = time.Now()

		resposta.Mensagens = append(resposta.Mensagens, mensagem)

		c.JSONP(http.StatusInternalServerError, resposta)
		return

	}

	// Gera a mensagem
	mensagem.Codigo = http.StatusCreated
	mensagem.Mensagem = "País salvo com sucesso! "
	mensagem.DataHora = time.Now()

	// Gera a resposta
	resposta.Dados = append(resposta.Dados, pais)
	resposta.Mensagens = append(resposta.Mensagens, mensagem)
	resposta.Links = append(resposta.Links, models.ListAllLink())

	log.Println("País gravado e gerando as respostas ...")
	c.JSON(http.StatusCreated, resposta)

}

/**
	Handler/Manipulador de requisições para atualizar (PUT) um país na tabela países. Atualiza todos os campos.
	Pode ser melhorado para modificar somente os campos modificados.
 */

func updatePaises(c *gin.Context) {

	codigo := c.Param("codigo")

	// declara as variáveis
	var pais models.Paises
	var resposta models.PaisesResposta
	var mensagem models.Mensagem

	// versao da resposta
	resposta.Versao = os.Getenv("API_RSP_VERSION")

	// Bind o payload no objeto país
	if err := c.ShouldBindJSON(&pais); err != nil {

		mensagem.Codigo = http.StatusBadRequest
		mensagem.Mensagem = "Dados de entrada inválido! " + err.Error()
		mensagem.DataHora = time.Now()

		resposta.Mensagens = append(resposta.Mensagens, mensagem)

		c.JSONP(http.StatusBadRequest, resposta)
		return
	}

	if code,_ := strconv.Atoi(codigo); pais.Codigo != code {
		mensagem.Codigo = http.StatusBadRequest
		mensagem.Mensagem = "Codigo informado inválido! "
		mensagem.DataHora = time.Now()

		resposta.Mensagens = append(resposta.Mensagens, mensagem)

		c.JSONP(http.StatusBadRequest, resposta)
		return
	}

	// Valida os dados do país
	if err := pais.Validate(); err != nil {

		mensagem.Codigo = http.StatusBadRequest
		mensagem.Mensagem = "Erro na validação dos dados : " + err.Error()
		mensagem.DataHora = time.Now()

		resposta.Mensagens = append(resposta.Mensagens, mensagem)

		c.JSONP(http.StatusBadRequest, resposta)
		return

	}

	// Salva o novo país
	if err := pais.Save(); err != nil {

		mensagem.Codigo = http.StatusInternalServerError
		mensagem.Mensagem = "Não foi possível atualizar o país : " + err.Error()
		mensagem.DataHora = time.Now()

		resposta.Mensagens = append(resposta.Mensagens, mensagem)

		c.JSONP(http.StatusInternalServerError, resposta)
		return

	}

	// Gera a mensagem
	mensagem.Codigo = http.StatusNoContent
	mensagem.Mensagem = "País atualizado ("+strconv.Itoa(pais.Codigo)+") com sucesso! "
	mensagem.DataHora = time.Now()

	// Gera a resposta
	resposta.Dados = append(resposta.Dados, pais)
	resposta.Mensagens = append(resposta.Mensagens, mensagem)
	resposta.Links = append(resposta.Links, models.ListAllLink())

	log.Printf("====== Atualizando Países " + codigo + " ========")
	c.JSON(http.StatusNoContent, resposta)
}

/**
	Handler/Manipulador de requisições para excluir/remover (DELETE) um país da tabela países.
 */
func deletePaises(c *gin.Context) {

	codigo := c.Param("codigo")

	// declara as variáveis
	var pais models.Paises
	var resposta models.PaisesResposta
	var mensagem models.Mensagem

	// versao da resposta
	resposta.Versao = os.Getenv("API_RSP_VERSION")

	// pega o codigo do pais e converte
	pais.Codigo, _ = strconv.Atoi(codigo)

	if pais.Codigo <= 100 || pais.Codigo >= 10000 {
		mensagem.Codigo = http.StatusPreconditionFailed
		mensagem.Mensagem = "Codigo do pais e inválido! (<= 100 ou >= 10000) "
		mensagem.DataHora = time.Now()

		resposta.Mensagens = append(resposta.Mensagens, mensagem)

		c.JSONP(http.StatusPreconditionFailed, resposta)
		return
	}

	// Bind o payload no objeto país
	if err := pais.Find(); err != nil {

		mensagem.Codigo = http.StatusNotFound
		mensagem.Mensagem = "Pais nao encontrado! " + err.Error()
		mensagem.DataHora = time.Now()

		resposta.Mensagens = append(resposta.Mensagens, mensagem)

		c.JSONP(http.StatusNotFound, resposta)
		return
	}

	if err := pais.Delete(); err != nil {

		mensagem.Codigo = http.StatusBadGateway
		mensagem.Mensagem = "Pais nao pode ser encontrado! " + err.Error()
		mensagem.DataHora = time.Now()

		resposta.Mensagens = append(resposta.Mensagens, mensagem)

		c.JSONP(http.StatusBadGateway, resposta)
		return
	}

	// Gera a mensagem de resposta
	mensagem.Codigo = http.StatusNoContent
	mensagem.Mensagem = "País ("+strconv.Itoa(pais.Codigo)+") excluído com sucesso! "
	mensagem.DataHora = time.Now()

	// Gera a resposta
	resposta.Dados = append(resposta.Dados, pais)

	// Adiciona a mensagem ao Array Mensagem do objeto resposta
	resposta.Mensagens = append(resposta.Mensagens, mensagem)
	resposta.Links = append(resposta.Links, models.ListAllLink())

	log.Printf("Pais (" + codigo + ") excluído com sucesso! ")

	// Escreve a resposta no objeto response do HTTP
	c.JSON(http.StatusNoContent, resposta)

}

/**
	Handler/Manipulador de requisições para retornar (GET) o estado do serviço.
 */
func getStatus(c *gin.Context) {

	// Declara as variáveis: resposta e mensagem
	var resposta models.PaisesResposta
	var mensagem models.Mensagem

	// Gera a mensagem de resposta
	mensagem.Codigo = 200
	mensagem.Mensagem = "Serviço em Operação"
	mensagem.DataHora = time.Now()

	// Pega da variável de ambiente a versao da resposta
	resposta.Versao = os.Getenv("API_RSP_VERSION")

	// Adiciona a mensagem ao Array Mensagem do objeto resposta
	resposta.Mensagens = append(resposta.Mensagens, mensagem)

	// Escreve a resposta no objeto response do HTTP
	c.JSONP(http.StatusOK, resposta)

	log.Println("====== status do serviço ======")
}

/**
	Handler/Manipulador de requisições para listar (GET) os país da tabela países, de forma paginada. A definição da
	paginação (Qual página desejada e quantidade de registro por página) deverá ser passada via query string, veja um
    exemplo na documentação do swagger.
 */
func getPaisesAll(c *gin.Context) {

	// Pega do objeto request os parametros query string (pagina e limite)
	pageOpts := getPagination(c.Request)

	log.Println("Consulta todos paises ...")

	// Declara as variáveis: pais, paises (array), resposta e mensagem.
	var pais models.Paises
	var paises []models.Paises
	var resposta models.PaisesResposta
	var mensagem models.Mensagem

	// Pega da variável de ambiente a versao da resposta
	resposta.Versao = os.Getenv("API_RSP_VERSION")

	// Pesquisa a lista de paises no banco de dados.
	paises,err := pais.FindAll(pageOpts.Offset,pageOpts.Limit)

	// Verifica se houve erros na pesquisa.
	if err != nil {

		// Caso houve erros na pesquisa, gera a mensagem de retorno
		mensagem.Codigo = http.StatusNotFound
		mensagem.Mensagem = "Paises nao encontrado! " + err.Error()
		mensagem.DataHora = time.Now()

		// Adiciona a mensagem ao Array Mensagem do objeto resposta.
		resposta.Mensagens = append(resposta.Mensagens, mensagem)

		// Escreve a resposta no objeto response do HTTP
		c.JSONP(http.StatusNotFound, resposta)

		return

	}

	// Gera a paginacao e adiciona ao Array Paginação do objeto resposta
	resposta.Paginacao = getLinkPagination(pageOpts,pais.CountPage(pageOpts.Limit),pais.TableName())

	// Gera a mensagem
	mensagem.Codigo = http.StatusOK
	mensagem.Mensagem = "Países encontrados com sucesso! "
	mensagem.DataHora = time.Now()

	// Adiciona a mensagem ao Array Mensagem do objeto resposta
	resposta.Mensagens = append(resposta.Mensagens, mensagem)

	// Adiciona o registro encontrado ao Array Dados do objeto resposta.
	resposta.Dados = paises

	// Gera e adiciona os links global, ao Array Links do objeto resposta.
	resposta.Links = append(resposta.Links, models.CreateLink())

	log.Println("====== listando todos países ======")

	// Escreve a resposta no objeto response do HTTP
	c.JSON(http.StatusOK, resposta)
}

/**
	Handler/Manipulador de requisições para localizar (GET) um país na tabela países.
 */
func getPaisesByCode(c *gin.Context) {

	// Pega do path param o código do país a ser pesquisado.
	codigo := c.Param("codigo")

	// Declara os objetos : pais, resposta e mensagem.
	var pais models.Paises
	var resposta models.PaisesResposta
	var mensagem models.Mensagem

	// Pega da variável de ambiente a versao da resposta
	resposta.Versao = os.Getenv("API_RSP_VERSION")

	// pega o codigo do pais e converte de string para int e atribui a propriedade Codigo do objeto país.
	pais.Codigo, _ = strconv.Atoi(codigo)

	// Valida se o código informado no path esta na faixa de valores permitidos (de 100 a 9999).
	if pais.Codigo <= 100 || pais.Codigo >= 10000 {

		// caso seja um código inválido, gera a mensagem de falha.
		mensagem.Codigo = http.StatusPreconditionFailed
		mensagem.Mensagem = "Codigo do pais e inválido! (<= 100 ou >= 10000) "
		mensagem.DataHora = time.Now()

		// Adiciona a mensagem ao Array Mensagem do objeto resposta.
		resposta.Mensagens = append(resposta.Mensagens, mensagem)

		// Escreve a resposta no objeto response do HTTP
		c.JSONP(http.StatusPreconditionFailed, resposta)

		return
	}

	// Pesquisa no banco de dados o país.
	if err := pais.Find(); err != nil {

		// Caso o país não seja encontrado gera a mensagem resposta.
		mensagem.Codigo = http.StatusNotFound
		mensagem.Mensagem = "Pais nao encontrado! " + err.Error()
		mensagem.DataHora = time.Now()

		// Adiciona a mensagem ao objeto de resposta padrão
		resposta.Mensagens = append(resposta.Mensagens, mensagem)

		// Escreve a resposta no objeto response do HTTP
		c.JSONP(http.StatusNotFound, resposta)

		return
	}

	// Gera os links específico do Pais
	pais.SetLinks()

	// Gera a mensagem de OK
	mensagem.Codigo = http.StatusOK
	mensagem.Mensagem = "País encontrado com sucesso! "
	mensagem.DataHora = time.Now()

	// Adiciona o registro encontrado ao Array Dados do objeto resposta.
	resposta.Dados = append(resposta.Dados, pais)

	// Adiciona a mensagem ao Array Mensagem do objeto resposta
	resposta.Mensagens = append(resposta.Mensagens, mensagem)

	// Gera e adiciona os links global, ao Array Links do objeto resposta.
	resposta.Links = append(resposta.Links, models.ListAllLink())

	log.Println("País encontrado com sucesso ...")

	// Escreve a resposta no objeto response do HTTP
	c.JSON(http.StatusOK, resposta)
}

