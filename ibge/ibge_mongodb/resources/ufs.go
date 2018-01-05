package resources

import (
	"github.com/italiviocorrea/golang/commons"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/italiviocorrea/golang/ibge/ibge_mongodb/Mongodb"
	"github.com/italiviocorrea/golang/ibge/ibge_mongodb/persistences"
	"github.com/italiviocorrea/golang/ibge/ibge_mongodb/models"
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

	uf := &dataResource
	// Cria um novo DatStore para manipular o MongoDB.
	dataStore := Mongodb.NewDataStore()
	// Adiciona o mgo.Session.Close()
	defer dataStore.Close()
	// Pega a mgo.Collection da "UFs"
	col := dataStore.Collection("ufs")
	// Cria uma instancia de UfStore
	ufStore := persistences.UfStore{C: col}
	// Insere um UF
	err = ufStore.Create(uf)

	if err != nil {
		commons.DisplayAppError(
			w,
			err,
			"Nao foi possivel criar a UF! (2)",
			500,
		)
		return
	}

	ufResp, err := ufStore.GetByCode(uf.Codigo)

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

	vars := mux.Vars(r)

	codigo,err := strconv.Atoi(vars["codigo"])

	if err != nil {
		commons.DisplayAppError(
			w,
			err,
			"Codigo nao foi informado!",
			500,
		)
		return
	}

	var dataResource models.Uf

	// Decodifica a entrada JSON da UF
	err1 := json.NewDecoder(r.Body).Decode(&dataResource)
	if err1 != nil {
		commons.DisplayAppError(
			w,
			err1,
			"Invalido dados da UF!",
			500,
		)
		return
	}
	// Atualiza a UF
	uf := dataResource
	uf.Codigo = codigo

	dataStore := Mongodb.NewDataStore()
	defer dataStore.Close()

	col := dataStore.Collection("ufs")

	ufStore := persistences.UfStore{C: col}

	// Atualiza a UF
	if err := ufStore.Update(uf); err != nil {
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

func GetUfs(w http.ResponseWriter, r *http.Request) {

	// Calcula os dados de paginacao
	pageOpts := commons.GetPagination(r)

	log.Println("Consulta todas UFs ...")

	dataStore := Mongodb.NewDataStore()
	defer dataStore.Close()
	col := dataStore.Collection("ufs")
	ufStore := persistences.UfStore{C: col}

	ufs := ufStore.GetAll(pageOpts.Offset, pageOpts.Limit)

	// gera a resposta
	j, err := json.Marshal(models.UfsResource{Data: ufs,
		Pagination: commons.GetLinkPagination(pageOpts, ufStore.GetCountPage(pageOpts.Limit),"ufs"),
		Status:     commons.StatusResponse{Code: 200, Message: "Sucesso"},
		Version:    "2.0",
		Link:       commons.Link{Name: "create", Method: "POST", Href: commons.AppConfig.Context+"/ufs"}})

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

func GetUfByID(w http.ResponseWriter, r *http.Request) {

	// Get id from the incoming url
	vars := mux.Vars(r)
	codigo,err := strconv.Atoi(vars["codigo"])

	if err != nil {
		commons.DisplayAppError(
			w,
			err,
			"Codigo nao foi informado!",
			500,
		)
		return
	}

	dataStore := Mongodb.NewDataStore()
	defer dataStore.Close()

	col := dataStore.Collection("ufs")
	ufStore := persistences.UfStore{C: col}

	uf, err1 := ufStore.GetByCode(codigo)

	if err1 != nil {
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

	codigo,err := strconv.Atoi(vars["codigo"])

	if err != nil {
		commons.DisplayAppError(
			w,
			err,
			"Codigo nao foi informado!",
			500,
		)
		return
	}

	dataStore := Mongodb.NewDataStore()
	defer dataStore.Close()
	col := dataStore.Collection("ufs")
	ufStore := persistences.UfStore{C: col}

	// Exclui a UF
	err1 := ufStore.Delete(codigo)

	if err1 != nil {
		commons.DisplayAppError(
			w,
			err1,
			"Ocorreu um erro nao esperado",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
