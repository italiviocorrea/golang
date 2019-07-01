package services

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/italiviocorrea/golang/grpc/nfcexml/model"
	"github.com/italiviocorrea/golang/grpc/nfcexml/nfcexmlpb"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

type Server struct {
}

/**
Handler/Manipulador de requisições para criar/inserir (POST) um país na tabela países.
*/

func (*Server) Create(ctx context.Context, req *nfcexmlpb.CreateRequest) (*nfcexmlpb.CreateResponse, error) {
	fmt.Println("Create XML")
	nfce := req.GetNfcexml()

	data := model.Nfcexml{
		DadosXML: nfce.GetDadosxml(),
	}

	// Valida os dados do país
	if err := data.Validate(); err != nil {

		return &nfcexmlpb.CreateResponse{
			Nfcexml: &nfcexmlpb.Nfcexml{
				Id:       nfce.GetId(),
				Dadosxml: nfce.GetDadosxml(),
			},
			Mensagem: &nfcexmlpb.Mensagem{
				Codigo:   http.StatusBadRequest,
				Mensagem: "Erro na validação dos dados : " + err.Error(),
			},
		}, err

	}
	// Inserir no banco de dados
	// Salva o novo país
	if _, err := data.Create(); err != nil {

		return &nfcexmlpb.CreateResponse{
			Mensagem: &nfcexmlpb.Mensagem{
				Codigo:   http.StatusInternalServerError,
				Mensagem: "Não foi possível gravar o país : " + err.Error(),
			},
		}, err

	}

	// Retorna a resposta
	return &nfcexmlpb.CreateResponse{
		Nfcexml: &nfcexmlpb.Nfcexml{
			Id:       data.ID.String(),
			Dadosxml: data.DadosXML,
		},
		Mensagem: &nfcexmlpb.Mensagem{
			Codigo:   http.StatusCreated,
			Mensagem: "País criado com sucesso! ",
		},
	}, nil

}

// Handler/Manipulador de requisições para excluir/remover (DELETE) um país da tabela países.

func (*Server) Delete(ctx context.Context, req *nfcexmlpb.DeleteRequest) (*nfcexmlpb.DeleteResponse, error) {

	fmt.Println("Excluir nfcexml request")

	id, _ := gocql.ParseUUID(req.Id)
	data := model.Nfcexml{
		ID: id,
	}

	if err := data.Delete(); err != nil {

		return &nfcexmlpb.DeleteResponse{
			Mensagem: &nfcexmlpb.Mensagem{
				Codigo:   http.StatusBadGateway,
				Mensagem: "País não pode ser excluído! " + err.Error(),
			},
		}, err
	}
	// Retorna a resposta
	return &nfcexmlpb.DeleteResponse{
		Mensagem: &nfcexmlpb.Mensagem{
			Codigo:   http.StatusNoContent,
			Mensagem: "Nfcexml excluído com sucesso! ",
		},
	}, nil

}

/**
Handler/Manipulador de requisições para localizar (GET) um país na tabela países.
*/
func (*Server) Find(ctx context.Context, req *nfcexmlpb.FindRequest) (*nfcexmlpb.FindResponse, error) {

	id, _ := gocql.ParseUUID(req.Id)
	data := model.Nfcexml{
		ID: id,
	}

	if err := data.Find(); err != nil {

		return &nfcexmlpb.FindResponse{
			Mensagem: &nfcexmlpb.Mensagem{
				Codigo:   http.StatusNotFound,
				Mensagem: "XML nao encontrado! " + err.Error(),
			},
		}, err
	}

	// Retorna a resposta
	return &nfcexmlpb.FindResponse{
		Nfcexml: &nfcexmlpb.Nfcexml{
			Id:       data.ID.String(),
			Dadosxml: data.DadosXML,
		},
		Mensagem: &nfcexmlpb.Mensagem{
			Codigo:   http.StatusOK,
			Mensagem: "XML encontrado com sucesso! ",
		},
	}, nil

}

/**
	Handler/Manipulador de requisições para listar (GET) os país da tabela países, de forma paginada. A definição da
	paginação (Qual página desejada e quantidade de registro por página) deverá ser passada via query string, veja um
    exemplo na documentação do swagger.
*/
func (*Server) FindAll(req *nfcexmlpb.FindAllRequest, stream nfcexmlpb.Service_FindAllServer) error {

	fmt.Println("Lista XML request")

	var nfcexml model.Nfcexml
	var nfcexmles []model.Nfcexml

	// Pesquisa a lista de nfcexmles no banco de dados.
	nfcexmles = nfcexml.FindAll()

	log.Println("Imprimindo FindALL")
	for _, p := range nfcexmles {
		log.Println(p.DadosXML)

		stream.Send(&nfcexmlpb.FindAllResponse{
			Nfcexml: &nfcexmlpb.Nfcexml{
				Id:       p.ID.String(),
				Dadosxml: p.DadosXML,
			},
		})
	}

	return nil
}
