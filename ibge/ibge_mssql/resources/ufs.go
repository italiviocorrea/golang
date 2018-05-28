package resources

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/italiviocorrea/golang/commons"
	"github.com/italiviocorrea/golang/ibge/ibge_mssql/persistences"
	"github.com/italiviocorrea/golang/ibge/models"
	"log"
	"net/http"
	"strconv"
)

func CreateUf(w http.ResponseWriter, r *http.Request) {

	log.Println("Criando uma UF ...")

	var dataResource models.Uf
	// Decodifica a entrada JSON da UF
	err := json.NewDecoder(r.Body).Decode(&dataResource)

	if err != nil {
		commons.DisplayAppError(
			w,
			err,
			"Nao foi possivel decodificar os dados da UF! (1)",
			500,
		)
		return
	}

	// valida os dados
	err = dataResource.Validate()

	if err != nil {
		commons.DisplayAppError(
			w,
			err,
			"Erro na validação dos dados!",
			500,
		)
		return
	}

	// Insere um UF
	_, err = persistences.CreateUF(dataResource)

	if err != nil {
		commons.DisplayAppError(
			w,
			err,
			"Nao foi possivel criar a UF! (2)",
			500,
		)
		return
	}

	ufResp, err := persistences.GetUFByCode(dataResource.Codigo)
	j, err := json.Marshal(models.UfResource{Data: ufResp,
		Status:  commons.StatusResponse{Code: http.StatusCreated, Message: "UF criado com sucesso!"},
		Version: "2.0"})
	// if um erro ocorreu,
	// envia uma resposta JSON usando um funcao helper common.DisplayAppError
	if err != nil {
		commons.DisplayAppError(
			w,
			err,
			"Ocorreu um erro nao esperado",
			500,
		)
		return
	}

	commons.ResponseWithJSON(w, j, http.StatusCreated)

}

func UpdateUf(w http.ResponseWriter, r *http.Request) {

	// Get codigo from the incoming url
	vars := mux.Vars(r)

	codigo, _ := strconv.Atoi(vars["codigo"])

	var dataResource models.Uf
	// Decodifica a entrada JSON da UF
	err := json.NewDecoder(r.Body).Decode(&dataResource)

	if err != nil {
		commons.DisplayAppError(
			w,
			err,
			"Nao foi possivel decodificar os dados da UF!",
			500,
		)
		return
	}

	// valida os dados
	err = dataResource.Validate()

	if err != nil {
		commons.DisplayAppError(
			w,
			err,
			"Erro na validação dos dados!",
			500,
		)
		return
	}

	if codigo != dataResource.Codigo {
		commons.DisplayAppError(
			w,
			err,
			"Codigo do municipio informado e invalido!",
			500,
		)
		return
	}

	// Atualiza a UF
	_, err = persistences.UpdateUF(dataResource)

	if err != nil {
		commons.DisplayAppError(
			w,
			err,
			"Ocorreu um erro nao esperado :"+err.Error(),
			500,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

func GetUfs(w http.ResponseWriter, r *http.Request) {

	// Calcula os dados de paginacao
	pageOpts := commons.GetPagination(r)

	log.Println("Consulta todas UFs ...")

	log.Println("Page offset : "+ strconv.Itoa(pageOpts.Offset))
	log.Println("Page limit  : "+ strconv.Itoa(pageOpts.Limit))


	ufs := persistences.GetAllUF(pageOpts.Offset, pageOpts.Limit)

	// gera a resposta
	j, err := json.Marshal(models.UfsResource{Data: ufs,
		Pagination: commons.GetLinkPagination(pageOpts, persistences.GetUFCountPage(pageOpts.Limit), commons.AppConfig.Context+"/ufs"),
		Status:     commons.StatusResponse{Code: 200, Message: "Sucesso"},
		Version:    "2.0",
		Link:       commons.Link{Name: "create", Method: "POST", Href: commons.AppConfig.Context + "/ufs"}})

	if err != nil {
		commons.DisplayAppError(
			w,
			err,
			"Ocorreu um erro nao esperado",
			500,
		)
		return
	}

	commons.ResponseWithJSON(w, j, http.StatusOK)

}

func GetUfByCode(w http.ResponseWriter, r *http.Request) {

	// Get id from the incoming url
	vars := mux.Vars(r)
	codigo, _ := strconv.Atoi(vars["codigo"])

	uf, err := persistences.GetUFByCode(codigo)

	if err != nil {
		commons.ResponseMessageWithJSON(
			w,
			"Nao foi encontrado UF",
			http.StatusNotFound,
		)

		return
	}
	j, err := json.Marshal(models.UfResource{Data: uf,
		Status:  commons.StatusResponse{Code: 200, Message: "Sucesso"},
		Version: "2.0"})
	if err != nil {
		commons.DisplayAppError(
			w,
			err,
			"Ocorreu um erro nao esperado",
			500,
		)
		return
	}
	commons.ResponseWithJSON(w, j, http.StatusOK)
}

func DeleteUf(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	codigo, _ := strconv.Atoi(vars["codigo"])
	// Exclui a UF informada
	_, err := persistences.DeleteUF(codigo)
	if err != nil {
		commons.DisplayAppError(
			w,
			err,
			"Ocorreu um erro nao esperado",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
