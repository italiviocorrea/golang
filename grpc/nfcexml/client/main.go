package main

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"github.com/italiviocorrea/golang/grpc/nfcexml/nfcexmlpb"
	"google.golang.org/grpc"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	fmt.Println("NFC-e cliente")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)

	if err != nil {
		log.Fatal("Nao foi possivel conectar: %v", err)
	}

	client := nfcexmlpb.NewServiceClient(cc)

	// Excluir um registro

	// ler XML e compacta-lo
	//xml := readFile("/home/icorrea/tmp/myxml.xml")
	//
	//fmt.Println(xml)
	//
	//xmlByte := compressGzip(xml)
	//
	//nfce := &nfcexmlpb.Nfcexml{
	//	Dadosxml: xmlByte.Bytes(),
	//}
	//
	//cResp, err := client.Create(context.Background(), &nfcexmlpb.CreateRequest{Nfcexml: nfce})
	//
	//if err != nil {
	//	log.Fatal("Erro nao esperado %v", err)
	//}
	//fmt.Printf("NFC-e XML gravado com sucesso: %v",cResp)

	// Excluir um registro
	deleteRes, deleteErr := client.Delete(context.Background(), &nfcexmlpb.DeleteRequest{Id: "ab61379c-9b9e-11e9-928d-80fa5b34b9d2"})

	if deleteErr != nil {
		fmt.Printf("Error happened while deleting: %v \n", deleteErr)
	}
	fmt.Printf("NFC-e was deleted: %v \n", deleteRes)

	// Encontrar um registro
	fmt.Println("Encontrar uma NFC-e")

	nfceReq := &nfcexmlpb.FindRequest{Id: "e3ad13ca-9b9f-11e9-9290-80fa5b34b9d2"}
	nfceResp, nfceFindErr := client.Find(context.Background(), nfceReq)

	if nfceFindErr != nil {
		fmt.Printf("Error happened while reading: %v \n", nfceFindErr)
	}
	fmt.Printf("NFC-e encontrada: %v \n", nfceResp)

	fmt.Println(descompressGzip((nfceResp.Nfcexml.Dadosxml)))

	// Encontrar todos os registros
	fmt.Println("Encontrar todos as NFC-es")

	stream, err := client.FindAll(context.Background(), &nfcexmlpb.FindAllRequest{})

	if err != nil {
		fmt.Printf("Erro ao ler todos os paises %v", err)
	}

	for {

		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		// ,descompressGzip(res.GetNfcexml().Dadosxml)
		fmt.Printf("ID: %s\n", res.Nfcexml.Id)
		fmt.Printf("XML: %s\n", descompressGzip(res.Nfcexml.Dadosxml))
		//fmt.Println(descompressGzip((res.Nfcexml.Dadosxml)))
		//fmt.Println(res.Nfcexml)
	}
}

func readFile(name string) string {

	b, err := ioutil.ReadFile(name)

	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func compressGzip(texto string) bytes.Buffer {

	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	_, err := zw.Write([]byte(texto))
	if err != nil {
		log.Fatal(err)
	}

	if err := zw.Close(); err != nil {
		log.Fatal(err)
	}
	return buf
}

func descompressGzip(dadosxml []byte) string {

	readerBuffer := bytes.NewBuffer(dadosxml)

	zr, err := gzip.NewReader(readerBuffer)
	if err != nil {
		log.Fatal(err)
	}

	xml, err := io.Copy(os.Stdout, zr)

	if err != nil {
		log.Fatal(err)
	}

	if err := zr.Close(); err != nil {
		log.Fatal(err)
	}

	return string(xml)
}
