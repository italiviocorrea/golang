package persistences

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/italiviocorrea/golang/commons"
	"github.com/italiviocorrea/golang/ibge/ibge_mongodb/models"
	"strconv"
)

// MunicipioStore provides CRUD operations against the collection "Municipios".
type MunicipioStore struct {
	C *mgo.Collection
}

// Create inserts the value of struct Municipio into collection.
func (store MunicipioStore) Create(b *models.Municipio) error {
	// Assign a new bson.ObjectId
	b.ID = bson.NewObjectId()
	err := store.C.Insert(b)
	return err
}

// Modifica um documento da colecao.
func (store MunicipioStore) Update(b models.Municipio) error {
	// atualizacao parcial no MongoDB
	err := store.C.Update(bson.M{"codigo": b.Codigo},
		bson.M{"$set": bson.M{
			"nome":     b.Nome,
			"uf": 	    b.Uf,
		}})
	return err
}

// Exclui um documento da colecao.
func (store MunicipioStore) Delete(codigo int64) error {
	err := store.C.Remove(bson.M{"codigo": codigo})
	return err
}

// Retorna todos os documentos da colecao.
func (store MunicipioStore) GetAll(page_num int, page_size int) []models.MunicipiosResponse {
	// calcular a paginacao
	skips := page_size * (page_num - 1)

	// fazer a busca no banco de dados
	var b []models.MunicipiosResponse
	iter := store.C.Find(nil).Skip(skips).Limit(page_size).Sort("codigo").Iter()

	result := models.MunicipioResponse{}
	// municipio adaptado para listar somente os campos a serem exibidos
	municipio := models.MunicipiosResponse{}
	for iter.Next(&result) {
		// adiciona os links
		links := []commons.Link{commons.Link{Name:"self",Method:"GET",Href:"/ibge/v3/municipios/"+strconv.FormatInt(result.Codigo,10)}}
		municipio.ID = result.ID
		municipio.Codigo = result.Codigo
		municipio.Nome = result.Nome
		municipio.Uf = result.Uf.Sigla
		municipio.Links = links
		// anexa o documento a lista
		b = append(b, municipio)
	}
	return b
}

// Calcula o n.o total de paginas com base no tamanho da pagina
func (store MunicipioStore) GetCountPage(page_size int) int {

	// fazer a busca no banco de dados
	total, err := store.C.Find(nil).Count()
	if err != nil {
		total = 1
	}
	if total > page_size {
		total = total / page_size
	}else{
		total = 1
	}
	return total
}

//func (store MunicipioStore) GetByID(id string) (models.MunicipioResponse, error) {
//	var b models.MunicipioResponse
//	log.Println("Pegando municipio ..."+id)
//	err := store.C.FindId(id).One(&b)
//	// adiciona os links
//	links := []commons.Link{commons.Link{Name:"update",Method:"PUT",Href:"/ibge/v3/municipios/"+strconv.FormatInt(b.Codigo,10)},
//		commons.Link{Name:"remove",Method:"DELETE",Href:"/ibge/v3/municipios/"+strconv.FormatInt(b.Codigo,10)}}
//	b.Links = links
//	// retorna a Municipio ou o erro
//	return b, err
//}


// Retorna um documento da colecao atraves do codigo da Municipio
func (store MunicipioStore) GetByCode(codigo int64) (models.MunicipioResponse, error) {

	// fazer a busca no banco de dados
	var b models.MunicipioResponse
	err := store.C.Find(bson.M{"codigo": codigo}).One(&b)
	// adiciona os links
	links := []commons.Link{	commons.Link{Name:"self",Method:"GET",Href:"/ibge/v3/municipios/"+strconv.FormatInt(codigo,10)},
	 						commons.Link{Name:"update",Method:"PUT",Href:"/ibge/v3/municipios/"+strconv.FormatInt(codigo,10)},
							commons.Link{Name:"remove",Method:"DELETE",Href:"/ibge/v3/municipios/"+strconv.FormatInt(codigo,10)}}
	b.Links = links
	// retorna a resposta
	return b, err
}
