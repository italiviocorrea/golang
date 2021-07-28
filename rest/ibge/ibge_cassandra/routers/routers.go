package routers

import (
	"github.com/gorilla/mux"
	"github.com/italiviocorrea/golang/routers"
)

// InitRoutes registra todas as rotas da aplicacao.
func InitRoutes() *mux.Router {

	router := mux.NewRouter().StrictSlash(false)

	// Rotas para a entidade UF e Municipio
	router = routers.SetSwaggersRoutes(router)
	router = routers.SetHeartBeatsRoutes(router)
	router = SetUFRoutes(router)
	router = SetMunicipioRoutes(router)

	return router
}
