package endpoints

import (
	"../models"
	"net/http"
	"os"
	"strconv"
)

type (

	PageOpts struct {
		Offset int
		Limit  int
	}
)

func getPagination(r *http.Request) PageOpts {

	// define os valores padroes de paginacao
	pageOpts := PageOpts{}
	pageOpts.Limit = 20
	pageOpts.Offset = 1

	// pega os query params de paginacao, caso tenham sido informados
	offset := r.URL.Query().Get("pagina")
	limit := r.URL.Query().Get("limite")

	if offset != "" {
		paginaI, err := strconv.ParseInt(offset, 10, 0)
		if err == nil {
			pageOpts.Offset = int(paginaI)
		}
	}

	if limit != "" {
		limiteI, err := strconv.ParseInt(limit, 10, 0)
		if err == nil {
			pageOpts.Limit = int(limiteI)
		}
	}
	return pageOpts
}

func getLinkPagination(pageOpts PageOpts, countPage int, resources string) models.Paginacao {

	page := models.Paginacao{}

	if countPage > 1 {
		page.Proxima = "GET "+resources+"?pagina=" + strconv.Itoa(pageOpts.Offset+1) + "&limite=" + strconv.Itoa(pageOpts.Limit)
	}
	if pageOpts.Offset > 1 {
		page.Anterior = "GET "+resources+"?pagina=" + strconv.Itoa(pageOpts.Offset-1) + "&limite=" + strconv.Itoa(pageOpts.Limit)
	}

	if pageOpts.Offset > 2 {
		page.Primeira = "GET "+resources+"?pagina=" + strconv.Itoa(1) + "&limite=" + strconv.Itoa(pageOpts.Limit)
	}
	if countPage > 2 {
		page.Ultima = "GET "+resources+"?pagina=" + strconv.Itoa(countPage) + "&limite=" + strconv.Itoa(pageOpts.Limit)
	}

	page.TotalPaginas = countPage

	return page
}

func setVarEnv() {

	if os.Getenv("API_EVT_CONTEXT") == "" {
		os.Setenv("API_EVT_CONTEXT", "/evt/v1")
	}

	if os.Getenv("API_MON_CONTEXT") == "" {
		os.Setenv("API_MON_CONTEXT", "/mon/v1/paises")
	}

	if os.Getenv("API_DOC_CONTEXT") == "" {
		os.Setenv("API_DOC_CONTEXT", "/docs/v1/paises/")
	}

	if os.Getenv("API_SRV_CONTEXT") == "" {
		os.Setenv("API_SRV_CONTEXT", "/api/v1")
	}

	if os.Getenv("API_RSP_VERSION") == "" {
		os.Setenv("API_RSP_VERSION", "v3")
	}

	if os.Getenv("API_SRV_LOGLEVEL") == "" {
		os.Setenv("API_SRV_LOGLEVEL", "4")
	}

}

