package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
)

type (
	/**
	Define a estrutura da classe Países.
	*/
	Paises struct {
		Codigo  int32   `json:"codigo" binding:"required" description:"codigo do pais" sql:"type:int" gorm:"primary_key"`
		Nome    string  `json:"nome"   binding:"required" description:"nome do pais" gorm:"not null;type:varchar(32)"`
	}

)

/*
	Implementa as regras de validação dos dados de países, antes da inclusão
	ou atualização no banco de dados.
*/


/**
Define o nome da tabela no banco de dados.
*/
func (t *Paises) TableName() string {
	return "paises"
}

/**
Valida os dados do pais. Existe uma outra forma, pode ser melhorado.
*/
func (t Paises) Validate() error {

	//var validNome = regexp.MustCompile(`[0-9]{2}`)

	if t.Codigo <= 100 || t.Codigo > 10000 {
		return errors.New("Código do País inválido! Deve ser maior que 100 e menor que 10000.")
	}

	if len(t.Nome) < 3 {
		return errors.New("Nome do País inválido! Nome muito curto.")
	}

	return nil
}

/**
	Garante que o fim da vigencia fique igual a NULL, caso nao seja informado. Este método será executado
    antes do método Create. Necessário para evite ficar com o valor default 0001-01-01.
*/
func (t *Paises) BeforeCreate(scope *gorm.Scope) error {
	//if t.FimVigencia != nil && t.FimVigencia.IsZero() {
	//	scope.SetColumn("FimVigencia", nil)
	//}
	return nil
}

/**
Cria o pais informado.
*/
func (t *Paises) Create() (int64, error) {

	var err error

	if err = db.Create(&t).Error; err != nil {
		log.Printf("Erro ao inserir país : " + err.Error())
		return -1, err
	}

	return db.RowsAffected, nil
}

/**
Grava modificações nos dados do pais.
*/
func (t *Paises) Save() error {

	var err error

	if err = db.Save(&t).Error; err != nil {
		log.Printf("Erro ao inserir país : " + err.Error())
		return err
	}

	return nil
}

/**
Remove fisicamente o pais do banco de dados.
*/
func (t Paises) Delete() error {

	var err error

	if err = db.Where("codigo = ?",t.Codigo).Delete(Paises{}).Error; err != nil {
		log.Printf("Erro ao remover o país : " + err.Error())
		return err
	}

	return nil
}

/**
Localiza e retorna um pais a partir do seu codigo
*/
func (t *Paises) Find() error {

	var err error

	if err = db.First(&t, "codigo = ?", t.Codigo).Error; err != nil {
		log.Println("Erro ao pesquisar pais :" + err.Error())
		return err
	}

	return nil

}

/**
Retorna uma lista de registro de forma paginada.
*/
func (t Paises) FindAll(page_num int32, page_size int32) ([]Paises, error) {

	// Declara a variável que receberá o resultado da pesquisa
	var paises []Paises

	// Declara a variável que receberá o resultado da pesquisa com os links, e que será devolvida como resposta.
	var paisResp []Paises

	// Declaração da variável que poderá receber o erro na pesquisa (caso ocorra).
	var err error

	// Calcula o offset (deslocamento) ou qtd de registros a ser pulado.
	var offset = (page_num - 1) * page_size

	// Consulta o banco de dados
	if err = db.Order("codigo").Limit(page_size).Offset(offset).Find(&paises).Error; err != nil {
		log.Println("Erro ao pesquisar pais :" + err.Error())
		return nil, err
	}

	// Fixa os links para cada pais retornado após a pesquisa, e armazenado na lista paisResp
	for _, pais := range paises {
		paisResp = append(paisResp, pais)
	}

	// Retorna a lista de paises com os links.
	return paisResp, nil

}

/**
Calcula o numero de paginas de acordo com o numero de registros na tabela.
*/
func (t Paises) CountPage(limit int) int {

	// variável para armazenar a qtd total de registro da tabela.
	var count int

	// Conta no banco de dados a quantidade de registro
	db.Table(t.TableName()).Count(&count)

	// Calcula a quantidade de páginas com base na quantidade registro e tamanho da página (limit).
	if count > limit {
		count = count / limit
		var remainder = count % limit
		if (remainder > 0) {
			count += 1
		}
	} else {
		count = 1
	}

	log.Printf("Total de Page: " + strconv.Itoa(count))

	// retorna a quantidade de ṕáginas
	return count

}

