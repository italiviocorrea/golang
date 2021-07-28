package main

import (
	"fmt"
	"github.com/italiviocorrea/golang/grpc/nfcexml/Cassandra"
	"github.com/italiviocorrea/golang/grpc/nfcexml/nfcexmlpb"
	services "github.com/italiviocorrea/golang/grpc/nfcexml/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {

	CassandraSession := Cassandra.Session
	defer CassandraSession.Close()

	lis, err := net.Listen("tcp", os.Getenv("API_SRV_ADDR"))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	nfcexmlpb.RegisterServiceServer(s, &services.Server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	go func() {
		fmt.Println("Iniciando o Servidor...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch
	fmt.Println("Parando o Servidor")
	s.Stop()
	fmt.Println("Fechando o Ouvinte")
	lis.Close()
	fmt.Println("Finalizado o Programa")

}
