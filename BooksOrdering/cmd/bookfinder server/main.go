package main

import (
	api "BooksOrdering/pkg/api"
	"BooksOrdering/pkg/bookfinder"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main (){
	grpcServer := grpc.NewServer()
	srv := &bookfinder.GRPCServer{}
	api.RegisterBookSearchServer(grpcServer, srv)

	listener, err := net.Listen("tcp", ":8070")
	if err != nil {
		log.Fatalf("Book Finder: Ошибка прослушивания: %v", err)
	}
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Book Finder: Ошибка запуска сервера: %v", err)
	}
}
