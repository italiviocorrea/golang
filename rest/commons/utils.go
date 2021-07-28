package commons

import (
	"encoding/json"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
	"strconv"
)

type (

 	RequestValidation interface {
		Validate() error
	}

	notFound struct {
		Message    string `json:"message"`
		HTTPStatus int    `json:"status"`
	}

	notFoundResource struct {
		Data notFound `json:"data"`
	}

	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HTTPStatus int    `json:"status"`
	}

	errorResource struct {
		Data appError `json:"data"`
	}

	Configuration struct {
		Server   string `default:"0.0.0.0:8080"`
		DBHost   string `default:"localhost"`
		DBPort   int    `default:"1433"`
		DBUser   string `default:"sa"`
		DBPwd    string `default:"senha#123"`
		Database string `default:"dbibgeapi"`
		Context  string `default:"/ibge/v3"`
		LogLevel int    `default:"4"`
	}

	Pagination struct {
		Has_first string `json:"has_first,omitempty"`
		Has_next string `json:"has_next,omitempty"`
		Has_prev string `json:"has_prev,omitempty"`
		Has_last string `json:"has_last,omitempty"`
		Total    int    `json:"total,omitempty"`
	}

	StatusResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message,omitempty"`
	}

	Link struct {
		Name   string `json:"name"`
		Method string `json:"method,omitempty"`
		Href   string `json:"href"`
		Rel    string `json:"rel,omitempty"`
	}

	PageOpts struct {
		Offset int
		Limit  int
	}
)


// DisplayAppError provides app specific error in JSON
func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {
	errObj := appError{
		Error:      handlerError.Error(),
		Message:    message,
		HTTPStatus: code,
	}
	//log.Printf("AppError]: %s\n", handlerError)
	//Error.Printf("AppError]: %s\n", handlerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		w.Write(j)
	}
}

func ResponseMessageWithJSON(w http.ResponseWriter, message string, code int) {
	notFoundObj := notFound{
		Message:    message,
		HTTPStatus: code,
	}
	//log.Printf("AppError]: %s\n", handlerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(notFoundResource{Data: notFoundObj}); err == nil {
		w.Write(j)
	}
}

func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}

// AppConfig holds the configuration values from config.json file
var AppConfig Configuration

// Initialize AppConfig
func initConfig(varEnvPrefix string) {
	loadAppEnv(varEnvPrefix)
}

// Reads configuration da env vars.
func loadAppEnv(varEnvPrefix string) {

	AppConfig = Configuration{}
	err := envconfig.Process(varEnvPrefix, &AppConfig)
	if err != nil {
		log.Fatalf("[loadAppEnv]: %s\n", err)
	}
	log.Println(AppConfig)
}

func GetPagination(r *http.Request) PageOpts {
	// define os valores padroes de paginacao
	pageOpts := PageOpts{}
	pageOpts.Limit = 20
	pageOpts.Offset = 1

	// pega os query params de paginacao, caso tenham sido informados
	offset := r.URL.Query().Get("page_num")
	limit := r.URL.Query().Get("page_size")

	if offset != "" {
		page_numI, err := strconv.ParseInt(offset, 10, 0)
		if err == nil {
			pageOpts.Offset = int(page_numI)
		}
	}

	if limit != "" {
		page_sizeI, err := strconv.ParseInt(limit, 10, 0)
		if err == nil {
			pageOpts.Limit = int(page_sizeI)
		}
	}
	return pageOpts
}

func GetLinkPagination(pageOpts PageOpts, countPage int, resources string) Pagination {
	page := Pagination{}
	if countPage > 1 {
		page.Has_next = "GET "+resources+"?page_num=" + strconv.Itoa(pageOpts.Offset+1) + "&page_size=" + strconv.Itoa(pageOpts.Limit)
	}
	if pageOpts.Offset > 1 {
		page.Has_prev = "GET "+resources+"?page_num=" + strconv.Itoa(pageOpts.Offset-1) + "&page_size=" + strconv.Itoa(pageOpts.Limit)
	}

	if pageOpts.Offset > 2 {
		page.Has_first = "GET "+resources+"?page_num=" + strconv.Itoa(1) + "&page_size=" + strconv.Itoa(pageOpts.Limit)
	}
	if countPage > 2 {
		page.Has_last = "GET "+resources+"?page_num=" + strconv.Itoa(countPage) + "&page_size=" + strconv.Itoa(pageOpts.Limit)
	}

	return page
}
