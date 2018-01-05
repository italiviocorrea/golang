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

func CreateMunicipio(w http.ResponseWriter, r *http.Request) {

	log.Println("Criando um Municipio ...")

	var dataResource models.Municipio

	// Decodifica a entrada JSON da Municipio
	err := json.NewDecoder(r.Body).Decode(&dataResource)

	if err != nil {
		commons.DisplayAppError(
			w,
			err,
			"Nao foi possivel decodificar os dados do Municipio! (1)",
			500,
		)
		return
	}

	municipio := &dataResource
	// Cria um novo DatStore para manipular o MongoDB.
	dataStore := Mongodb.NewDataStore()
	// Adiciona o mgo.Session.Close()
	defer dataStore.Close()
	// Pega a mgo.Collection da "Municipios"
	col := dataStore.Collection("municipios")
	// Cria uma instancia de MunicipioStore
	municipioStore := persistences.MunicipioStore{C: col}
	// Insere um Municipio
	err = municipioStore.Create(municipio)

	if err != nil {
		commons.DisplayAppError(
			w,
			err,
			"Nao foi possivel criar Municipio! (2)",
			500,
		)
		return
	}

	municipioResp, err := municipioStore.GetByCode(municipio.Codigo)

	j, err := json.Marshal(models.MunicipioResource{Data: municipioResp,
		Status:  commons.StatusResponse{Code: http.StatusCreated, Message: "Municipio criado com sucesso!"},
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

func UpdateMunicipio(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	codigo, err := strconv.ParseInt(vars["codigo"],10,0)

	if err != nil {
		commons.DisplayAppError(
			w,
			err,
			"Codigo nao foi informado!",
			500,
		)
		return
	}

	var dataResource models.Municipio

	// Decodifica a entrada JSON da Municipio
	err1 := json.NewDecoder(r.Body).Decode(&dataResource)

	if err1 != nil {
		commons.DisplayAppError(
			w,
			err1,
			"Invalido dados da Municipio!",
			500,
		)
		return
	}
	// Atualiza a Municipio
	municipio := dataResource
	municipio.Codigo = codigo

	dataStore := Mongodb.NewDataStore()
	defer dataStore.Close()

	col := dataStore.Collection("municipios")

	municipioStore := persistences.MunicipioStore{C: col}

	// Atualiza a Municipio
	if err := municipioStore.Update(municipio); err != nil {
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

func GetMunicipios(w http.ResponseWriter, r *http.Request) {

	// Calcula os dados de paginacao
	pageOpts := commons.GetPagination(r)

	log.Println("Consulta todas Municipios ...")

	dataStore := Mongodb.NewDataStore()
	defer dataStore.Close()
	col := dataStore.Collection("municipios")
	municipioStore := persistences.MunicipioStore{C: col}

	municipios := municipioStore.GetAll(pageOpts.Offset, pageOpts.Limit)

	// gera a resposta
	j, err := json.Marshal(models.MunicipiosResource{Data: municipios,
		Pagination: commons.GetLinkPagination(pageOpts, municipioStore.GetCountPage(pageOpts.Limit),"municipios"),
		Status:     commons.StatusResponse{Code: 200, Message: "Sucesso"},
		Version:    "2.0",
		Link:       commons.Link{Name: "create", Method: "POST", Href: commons.AppConfig.Context+"/municipios"}})

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

func GetMunicipioByID(w http.ResponseWriter, r *http.Request) {

	// Get id from the incoming url
	vars := mux.Vars(r)
	codigo,err := strconv.ParseInt(vars["codigo"],10,0)

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

	col := dataStore.Collection("municipios")
	municipioStore := persistences.MunicipioStore{C: col}

	municipio, err1 := municipioStore.GetByCode(codigo)

	if err1 != nil {
		commons.ResponseMessageWithJSON(
			w,
			"Nao foi encontrado Municipio",
			http.StatusNotFound,
		)

		return
	}
	j, err := json.Marshal(models.MunicipioResource{Data: municipio,
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

func DeleteMunicipio(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	codigo,err := strconv.ParseInt(vars["codigo"],10,0)

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
	col := dataStore.Collection("municipios")
	municipioStore := persistences.MunicipioStore{C: col}

	// Exclui a Municipio
	err1 := municipioStore.Delete(codigo)

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
