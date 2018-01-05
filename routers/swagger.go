package routers

import (
	"github.com/gorilla/mux"
	"github.com/italiviocorrea/golang/commons"
	"net/http"
	"flag"
)


//var staticContent = flag.String("staticPath", "./swagger-ui/", "Path to folder with Swagger UI")

func SetSwaggersRoutes(router *mux.Router) *mux.Router {

	var dir string

	flag.StringVar(&dir, "dir", "./swagger-ui/", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	router.PathPrefix(commons.AppConfig.Context + "/swagger-ui/").Handler(http.StripPrefix(commons.AppConfig.Context+"/swagger-ui/", http.FileServer(http.Dir(dir))))

	return router
}
