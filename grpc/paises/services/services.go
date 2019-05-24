package services

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"sefaz.ms.gov.br/cotin/paises/model"
	"sefaz.ms.gov.br/cotin/paises/paisespb"
)

type Server struct {

}

/**
	Handler/Manipulador de requisições para criar/inserir (POST) um país na tabela países.
*/

func (*Server) CriarPais(ctx context.Context, req *paisespb.CriarPaisRequest) (*paisespb.CriarPaisResponse, error) {
	fmt.Println("Criar pais request")
	pais := req.GetPais()

	data := model.Paises{
		Codigo: pais.GetCodigo(),
		Nome: pais.GetNome(),
	}

	// Valida os dados do país
	if err := data.Validate(); err != nil {

		return &paisespb.CriarPaisResponse{
			Pais: &paisespb.Pais{
				Codigo: pais.GetCodigo(),
				Nome: pais.GetNome(),
				InicioVigencia: pais.GetInicioVigencia(),
				FimVigencia: pais.FimVigencia,
			},
			Mensagem: &paisespb.Mensagem{
				Codigo: http.StatusBadRequest,
				Mensagem: "Erro na validação dos dados : " + err.Error(),
			},
		}, err

	}
	// Inserir no banco de dados
	// Salva o novo país
	if _, err := data.Create(); err != nil {

		return &paisespb.CriarPaisResponse{
			Pais: &paisespb.Pais{
				Codigo: pais.GetCodigo(),
				Nome: pais.GetNome(),
				InicioVigencia: pais.GetInicioVigencia(),
				FimVigencia: pais.FimVigencia,
			},
			Mensagem: &paisespb.Mensagem{
				Codigo: http.StatusInternalServerError,
				Mensagem: "Não foi possível gravar o país : " + err.Error(),
			},
		},  err

	}

	// Retorna a resposta
	return &paisespb.CriarPaisResponse{
		Pais: &paisespb.Pais{
			Codigo: pais.GetCodigo(),
			Nome: pais.GetNome(),
			InicioVigencia: pais.GetInicioVigencia(),
			FimVigencia: pais.FimVigencia,
		},
		Mensagem: &paisespb.Mensagem{
			Codigo: http.StatusCreated,
			Mensagem: "País criado com sucesso! ",
		},
	}, nil

}

/**
	Handler/Manipulador de requisições para atualizar (PUT) um país na tabela países. Atualiza todos os campos.
	Pode ser melhorado para modificar somente os campos modificados.
*/
func (*Server) ModificarPais(ctx context.Context, req *paisespb.ModificarPaisRequest)  (*paisespb.ModificarPaisResponse, error) {

	fmt.Println("Modificar pais request")
	pais := req.GetPais()

	data := model.Paises{
		Codigo: pais.GetCodigo(),
		Nome: pais.GetNome(),
	}

	// Valida os dados do país
	if err := data.Validate(); err != nil {

		return &paisespb.ModificarPaisResponse{
			Pais: &paisespb.Pais{
				Codigo: pais.GetCodigo(),
				Nome: pais.GetNome(),
			},
			Mensagem: &paisespb.Mensagem{
				Codigo: http.StatusBadRequest,
				Mensagem: "Erro na validação dos dados : " + err.Error(),
			},
		},err

	}

	// Salva as modificações do pais
	if err := data.Save(); err != nil {

		return &paisespb.ModificarPaisResponse{
			Pais: &paisespb.Pais{
				Codigo: pais.GetCodigo(),
				Nome: pais.GetNome(),
			},
			Mensagem: &paisespb.Mensagem{
				Codigo: http.StatusInternalServerError,
				Mensagem: "Não foi possível atualizar o país : " + err.Error(),
			},
		},err

	}

	// Retorna a resposta
	return &paisespb.ModificarPaisResponse{
		Pais: &paisespb.Pais{
			Codigo: pais.GetCodigo(),
			Nome: pais.GetNome(),
		},
		Mensagem: &paisespb.Mensagem{
			Codigo: http.StatusNoContent,
			Mensagem: "País atualizado com sucesso! ",
		},
	},nil

}

// Handler/Manipulador de requisições para excluir/remover (DELETE) um país da tabela países.

func (*Server) ExcluirPais(ctx context.Context, req *paisespb.ExcluirPaisRequest) (*paisespb.ExcluirPaisResponse, error)  {

	fmt.Println("Excluir pais request")

	data := model.Paises{
		Codigo: req.GetCodigo(),
	}

	if err := data.Delete(); err != nil {

		return &paisespb.ExcluirPaisResponse{
			Mensagem: &paisespb.Mensagem{
				Codigo: http.StatusBadGateway,
				Mensagem: "País não pode ser excluído! "+err.Error(),
			},
		},err
	}
	// Retorna a resposta
	return &paisespb.ExcluirPaisResponse{
		Mensagem: &paisespb.Mensagem{
			Codigo: http.StatusNoContent,
			Mensagem: "Pais excluído com sucesso! ",
		},
	},nil

}

/**
	Handler/Manipulador de requisições para localizar (GET) um país na tabela países.
*/
func (*Server) BuscarPais(ctx context.Context, req *paisespb.BuscarPaisRequest)  (*paisespb.BuscarPaisResponse, error) {


	data := model.Paises{
		Codigo: req.GetCodigo(),
	}

	if err := data.Find(); err != nil {

		return &paisespb.BuscarPaisResponse{
			Mensagem: &paisespb.Mensagem{
				Codigo: http.StatusNotFound,
				Mensagem: "Pais nao encontrado! " + err.Error(),
			},
		}, err
	}

	// Retorna a resposta
	return &paisespb.BuscarPaisResponse{
		Pais: &paisespb.Pais{
			Codigo: data.Codigo,
			Nome: data.Nome,
		},
		Mensagem: &paisespb.Mensagem{
			Codigo: http.StatusOK,
			Mensagem: "Pais encontrado com sucesso! ",
		},
	},nil


}

/**
	Handler/Manipulador de requisições para listar (GET) os país da tabela países, de forma paginada. A definição da
	paginação (Qual página desejada e quantidade de registro por página) deverá ser passada via query string, veja um
    exemplo na documentação do swagger.
*/
func (*Server) ListarTodosPaises(req *paisespb.BuscarTodosPaisesRequest, stream paisespb.PaisService_ListarTodosPaisesServer) error {

	fmt.Println("Lista paises request")

	var pais model.Paises
	var paises []model.Paises

	var offset int32 = 1
	var limit int32 = 20

	if req.PageNumber > 0 {
		offset = req.PageNumber
	}

	if  req.PageSize > 0 {
		limit = req.PageSize
	}

	// Pesquisa a lista de paises no banco de dados.
	paises,err := pais.FindAll(offset,limit)

	if err != nil {

		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}

	for _,p := range paises {

		stream.Send(&paisespb.BuscarTodosPaisesResponse{
			Pais: &paisespb.Pais{
				Codigo: p.Codigo,
				Nome: p.Nome,
			},
		})
	}

	return nil
}
