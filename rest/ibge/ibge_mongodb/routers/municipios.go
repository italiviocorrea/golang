package routers

import (
	"github.com/italiviocorrea/golang/commons"
	"github.com/gorilla/mux"
	"log"
	"github.com/italiviocorrea/golang/ibge/ibge_mongodb/resources"
)

// Define as rotas para UF
func SetMunicipiosRoutes(router *mux.Router) *mux.Router {

	//ufRouter := mux.NewRouter()
	router.HandleFunc(commons.AppConfig.Context+"/municipios", resources.CreateMunicipio).Methods("POST")
	router.HandleFunc(commons.AppConfig.Context+"/municipios/{codigo}", resources.UpdateMunicipio).Methods("PUT")
	router.HandleFunc(commons.AppConfig.Context+"/municipios", resources.GetMunicipios).Methods("GET")
	router.HandleFunc(commons.AppConfig.Context+"/municipios/{codigo}", resources.GetMunicipioByID).Methods("GET")
	router.HandleFunc(commons.AppConfig.Context+"/municipios/{codigo}", resources.DeleteMunicipio).Methods("DELETE")
	log.Println("Criando as rotas para municipios...")

	//router.PathPrefix("/municipios").Handler(common.AuthorizeRequest(ufRouter))
	return router

}
