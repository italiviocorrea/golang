package resources

import "net/http"

type UFResource interface {
	CreateUf(w http.ResponseWriter, r *http.Request)
	UpdateUf(w http.ResponseWriter, r *http.Request)
	DeleteUf(w http.ResponseWriter, r *http.Request)
	GetUfByCode(w http.ResponseWriter, r *http.Request)
	GetUfs(w http.ResponseWriter, r *http.Request)
}

type MunicipioResource interface {
	CreateMunicipio(w http.ResponseWriter, r *http.Request)
	UpdateMunicipio(w http.ResponseWriter, r *http.Request)
	DeleteMunicipio(w http.ResponseWriter, r *http.Request)
	GetMunicipioByCode(w http.ResponseWriter, r *http.Request)
	GetMunicipios(w http.ResponseWriter, r *http.Request)
}
