package routers

import (
	"github.com/gorilla/mux"
	"github.com/italiviocorrea/golang/routers"
)


func InitRoutes() *mux.Router {

	router := mux.NewRouter().StrictSlash(false)

	// Rotas para UF e Municipio
	router = routers.SetSwaggersRoutes(router)
	router = routers.SetHeartBeatsRoutes(router)
	router = SetUFRoutes(router)
	router = SetMunicipioRoutes(router)

	return router
}
