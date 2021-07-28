package routers

import (
	"github.com/italiviocorrea/golang/commons"
	"github.com/italiviocorrea/golang/ibge/ibge_cassandra/resources"
	"github.com/gorilla/mux"
	"log"
)

// Define as rotas para UF
func SetUFRoutes(router *mux.Router) *mux.Router {

	//ufRouter := mux.NewRouter()
	router.HandleFunc(commons.AppConfig.Context+"/ufs", resources.CreateUf).Methods("POST")
	router.HandleFunc(commons.AppConfig.Context+"/ufs/{codigo}", resources.UpdateUf).Methods("PUT")
	router.HandleFunc(commons.AppConfig.Context+"/ufs", resources.GetUfs).Methods("GET")
	router.HandleFunc(commons.AppConfig.Context+"/ufs/{codigo}", resources.GetUfByCode).Methods("GET")
	router.HandleFunc(commons.AppConfig.Context+"/ufs/{codigo}", resources.DeleteUf).Methods("DELETE")
	log.Println("Criando as rotas ...")

	//router.PathPrefix("/ufs").Handler(common.AuthorizeRequest(ufRouter))
	return router

}
