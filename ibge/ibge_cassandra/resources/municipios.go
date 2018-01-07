package resources

import (
	"github.com/italiviocorrea/golang/commons"
	"github.com/italiviocorrea/golang/ibge/models"
	"github.com/gorilla/mux"
	"strconv"
	"net/http"
	"log"
	"encoding/json"
	"github.com/italiviocorrea/golang/ibge/ibge_cassandra/persistences"
)

func CreateMunicipio(w http.ResponseWriter, r *http.Request) {

	log.Println("Criando uma MUNICIPIO ...")

	var dataResource models.Municipio

	// Decodifica a entrada JSON para a entidade municipio
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

	// Cria o municipio
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

	munResp, err := persistences.GetMunicipioByCode(dataResource)

	j, err := json.Marshal(
		models.MunicipioResource{
			Data: munResp,
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

	// Get codigo from the incoming url
	vars := mux.Vars(r)

	codigo, _ := strconv.ParseInt(vars["codigo"],10,0)

	var dataResource models.Municipio

	// Decodifica a entrada JSON para a entidade municipio
	err := json.NewDecoder(r.Body).Decode(&dataResource)

	if err != nil {
		commons.DisplayAppError(
			w,
			err,
			"Erro na decodificacao!",
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


	// Atualiza os dados do municipio
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
	j, err := json.Marshal(
		models.MunicipiosResource{
			Data: municipios,
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

	municipio := models.Municipio{Codigo:codigo}

	municipios, err := persistences.GetMunicipioByCode(municipio)

	if err != nil {
		commons.ResponseMessageWithJSON(
			w,
			"Nao foi encontrado o municipio informado!",
			http.StatusNotFound,
		)

		return
	}

	j, err := json.Marshal(
		models.MunicipioResource{
			Data: municipios,
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

	// Exclui o municipio informado
	municipio := models.Municipio{Codigo:codigo}

	_, err := persistences.DeleteMunicipio(municipio)

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
