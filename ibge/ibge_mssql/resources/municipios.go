package resources

import (
	"com/ItalivioCorrea/commons"
	"com/ItalivioCorrea/ibge/models"
	"com/ItalivioCorrea/ibge/ibge_mssql/persistences"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func CreateMunicipio(w http.ResponseWriter, r *http.Request) {

	log.Println("Criando um Municipio ...")

	var dataResource models.Municipio

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

	_, err = persistences.CreateMunicipio(dataResource)

	if err != nil {
		commons.DisplayAppError(
			w,
			err,
			"Nao foi possivel criar Municipio! (2)",
			500,
		)
		return
	}

	munResp, err := persistences.GetMunicipioByCode(dataResource.Codigo)

	j, err := json.Marshal(models.MunicipioResource{Data: munResp,
		Status:  commons.StatusResponse{Code: http.StatusCreated, Message: "Municipio criado com sucesso!"},
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

	commons.ResponseWithJSON(w, j, http.StatusCreated)

}

func UpdateMunicipio(w http.ResponseWriter, r *http.Request) {

	var dataResource models.Municipio

	// Decodifica a entrada JSON da MUNICIPIO
	err := json.NewDecoder(r.Body).Decode(&dataResource)

	if err != nil {
		commons.DisplayAppError(
			w,
			err,
			"Invalido dados do Municipio!",
			500,
		)
		return
	}

	// Atualiza a MUNICIPIO
	_, err = persistences.UpdateMunicipio(dataResource)

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

func GetMunicipios(w http.ResponseWriter, r *http.Request) {

	// Calcula os dados de paginacao
	pageOpts := commons.GetPagination(r)

	log.Println("Consulta todos os Municipios ...")

	municipios := persistences.GetAllMunicipio(pageOpts.Offset, pageOpts.Limit)

	// gera a resposta
	j, err := json.Marshal(models.MunicipiosResource{Data: municipios,
		Pagination: commons.GetLinkPagination(pageOpts, persistences.GetMunicipioCountPage(pageOpts.Limit),"municipios"),
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

func GetMunicipioByCode(w http.ResponseWriter, r *http.Request) {

	// Get codigo from the incoming url
	vars := mux.Vars(r)
	codigo, _ := strconv.ParseInt(vars["codigo"],10,0)

	municipio, err := persistences.GetMunicipioByCode(codigo)

	if err != nil {
		commons.ResponseMessageWithJSON(
			w,
			"Nao foi encontrado o municipio informado!",
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
	codigo, _ := strconv.ParseInt(vars["codigo"],10,0)

	// Exclui a MUNICIPIO informada
	_, err := persistences.DeleteMunicipio(codigo)

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
