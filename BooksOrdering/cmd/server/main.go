package main

import (
	api "BooksOrdering/pkg/api"
	"BooksOrdering/pkg/order"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main (){
	grpcServer := grpc.NewServer()
	srv := &order.GRPCServer{}
	api.RegisterOrderServer(grpcServer, srv)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Order: Ошибка прослушивания: %v", err)
	}
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Order: Ошибка запуска сервера: %v", err)
	}
}