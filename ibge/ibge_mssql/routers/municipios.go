package routers

import (
	"github.com/italiviocorrea/golang/commons"
	"github.com/italiviocorrea/golang/ibge/ibge_mssql/resources"
	"github.com/gorilla/mux"
	"log"
)

func SetMunicipioRoutes(router *mux.Router) *mux.Router {

	router.HandleFunc(commons.AppConfig.Context+"/municipios", resources.CreateMunicipio).Methods("POST")
	router.HandleFunc(commons.AppConfig.Context+"/municipios/{codigo}", resources.UpdateMunicipio).Methods("PUT")
	router.HandleFunc(commons.AppConfig.Context+"/municipios", resources.GetMunicipios).Methods("GET")
	router.HandleFunc(commons.AppConfig.Context+"/municipios/{codigo}", resources.GetMunicipioByCode).Methods("GET")
	router.HandleFunc(commons.AppConfig.Context+"/municipios/{codigo}", resources.DeleteMunicipio).Methods("DELETE")
	log.Println("Criando as rotas ...")

	return router

}
