package persistences

import "com/ItalivioCorrea/ibge/models"

type UfPersistence interface {
	CreateUF(uf models.Uf) (int64, error)
	UpdateUF(uf models.Uf) (int64, error)
	DeleteUF(codigo int) (int64, error)
	GetAllUF(page_num int, page_size int) []models.UFResponse
	GetUFByCode(codigo int) (models.UFResponse, error)
	GetUFCountPage(page_size int) int
}

type MunicipioPersistence interface {
	CreateMunicipio(municipio models.Municipio) (int64, error)
	UpdateUF(municipio models.Municipio) (int64, error)
	DeleteMunicipio(codigo int64) (int64, error)
	GetAllMunicipio(page_num int, page_size int) []models.MunicipioResponse
	GetMunicipioByCode(codigo int64) (models.MunicipioResponse, error)
	GetMunicipioCountPage(page_size int) int
}
