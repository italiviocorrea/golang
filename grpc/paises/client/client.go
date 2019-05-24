package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"sefaz.ms.gov.br/cotin/paises/paisespb"
)

func main()  {


	fmt.Println("Pais Client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := paisespb.NewPaisServiceClient(cc)

	// excluir
	deleteRes, deleteErr := c.ExcluirPais(context.Background(), &paisespb.ExcluirPaisRequest{Codigo: 9999})

	if deleteErr != nil {
		fmt.Printf("Error happened while deleting: %v \n", deleteErr)
	}
	fmt.Printf("Pais was deleted: %v \n", deleteRes)

	// Criar um pais
	fmt.Println("Criando um pais")

	pais := &paisespb.Pais{
		Codigo: 9999,
		Nome: "Teste de gravacao",
	}

	paisResposta, err3 := c.CriarPais(context.Background(), &paisespb.CriarPaisRequest{Pais: pais})
	if err3 != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Pais has been created: %v", paisResposta)

	// encontrar um pais
	fmt.Println("Encontrando um país")

	_, err = c.BuscarPais(context.Background(),&paisespb.BuscarPaisRequest{Codigo: 9999})

	if err != nil {
		fmt.Printf("Error happened while reading: %v \n", err)
	}

	buscarPaisReq := &paisespb.BuscarPaisRequest{Codigo: 9999}
	buscarPaisRes, buscarPaisErr := c.BuscarPais(context.Background(),buscarPaisReq)

	if buscarPaisErr != nil {
		fmt.Printf("Error happened while reading: %v \n", buscarPaisErr)
	}

	fmt.Printf("Pais encontrado: %v \n", buscarPaisRes)


	// atualizar paises
	// Criar um pais
	fmt.Println("Modificando um pais")

	paisUpdate := &paisespb.Pais{
		Codigo: 9999,
		Nome: "Modificando o Pais",
	}

	paisResposta1, err4 := c.ModificarPais(context.Background(), &paisespb.ModificarPaisRequest{Pais: paisUpdate})
	if err4 != nil {
		log.Fatalf("Unexpected error: %v", err4)
	}
	fmt.Printf("Pais was updated: %v", paisResposta1)


	// Listar paises
	fmt.Println("Listando Paises")

	stream, err := c.ListarTodosPaises(context.Background(), &paisespb.BuscarTodosPaisesRequest{PageNumber: 3, PageSize: 100})

	if err != nil {
		log.Fatalf("error while calling ListarTodosPaises RPC: %v", err)
	}
	
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(res.GetPais())
	}



}

