package routers

import (
	"github.com/gorilla/mux"
	"com/ItalivioCorrea/resources"
)

func SetHeartBeatsRoutes(router *mux.Router) *mux.Router {

	router.HandleFunc("/", resources.Heartbeat)

	return router
}

