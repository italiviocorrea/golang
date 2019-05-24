package routers

import (
	"github.com/gorilla/mux"
	"github.com/italiviocorrea/golang/resources"
)

func SetHeartBeatsRoutes(router *mux.Router) *mux.Router {

	router.HandleFunc("/", resources.Heartbeat)

	return router
}

