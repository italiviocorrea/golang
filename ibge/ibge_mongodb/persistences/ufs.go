package persistences

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"com/ItalivioCorrea/ibge/ibge_mongodb/models"
	"com/ItalivioCorrea/commons"
	"strconv"
)

// UfStore provides CRUD operations against the collection "Ufs".
type UfStore struct {
	C *mgo.Collection
}

// Create inserts the value of struct Uf into collection.
func (store UfStore) Create(b *models.Uf) error {
	// Assign a new bson.ObjectId
	b.ID = bson.NewObjectId()
	err := store.C.Insert(b)
	return err
}

// Modifica um documento da colecao.
func (store UfStore) Update(b models.Uf) error {
	// atualizacao parcial no MongoDB
	err := store.C.Update(bson.M{"codigo": b.Codigo},
		bson.M{"$set": bson.M{
			"nome":     b.Nome,
			"sigla":    b.Sigla,
		}})
	return err
}

// Exclui um documento da colecao.
func (store UfStore) Delete(codigo int) error {
	err := store.C.Remove(bson.M{"codigo": codigo})
	return err
}

// Retorna todos os documentos da colecao.
func (store UfStore) GetAll(page_num int, page_size int) []models.UFResponse {
	// calcular a paginacao
	skips := page_size * (page_num - 1)

	// fazer a busca no banco de dados
	var b []models.UFResponse
	iter := store.C.Find(nil).Skip(skips).Limit(page_size).Sort("codigo").Iter()

	result := models.UFResponse{}
	// Uf adaptado para listar somente os campos a serem exibidos
	Uf := models.UFResponse{}
	for iter.Next(&result) {
		// adiciona os links
		links := []commons.Link{commons.Link{Name:"self",Method:"GET",Href:"/ibge/v3/Ufs/"+strconv.Itoa(result.Codigo)}}
		Uf.ID = result.ID
		Uf.Codigo = result.Codigo
		Uf.Nome = result.Nome
		Uf.Sigla = result.Sigla
		Uf.Links = links
		// anexa o documento a lista
		b = append(b, Uf)
	}
	return b
}

// Calcula o n.o total de paginas com base no tamanho da pagina
func (store UfStore) GetCountPage(page_size int) int {

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

//func (store UfStore) GetByID(id string) (models.UFResponse, error) {
//
//	var b models.UFResponse
//
//	log.Println("Pegando UF ..."+id)
//
//	err := store.C.FindId(bson.ObjectIdHex(id)).One(&b)
//
//	// adiciona os links
//	links := []commons.Link{commons.Link{Name:"update",Method:"PUT",Href:"/ibge/v3/ufs/"+b.ID.Hex()},
//		commons.Link{Name:"remove",Method:"DELETE",Href:"/ibge/v3/ufs/"+b.ID.Hex()}}
//
//	b.Links = links
//
//	// retorna a UF ou o erro
//	return b, err
//}


// Retorna um documento da colecao atraves do codigo da Uf
func (store UfStore) GetByCode(codigo int) (models.UFResponse, error) {

	// fazer a busca no banco de dados
	var b models.UFResponse
	err := store.C.Find(bson.M{"codigo": codigo}).One(&b)
	// adiciona os links
	links := []commons.Link{	commons.Link{Name:"self",Method:"GET",Href:"/ibge/v3/Ufs/"+strconv.Itoa(b.Codigo)},
	 						commons.Link{Name:"update",Method:"PUT",Href:"/ibge/v3/Ufs/"+strconv.Itoa(b.Codigo)},
							commons.Link{Name:"remove",Method:"DELETE",Href:"/ibge/v3/Ufs/"+strconv.Itoa(b.Codigo)}}
	b.Links = links
	// retorna a resposta
	return b, err
}
