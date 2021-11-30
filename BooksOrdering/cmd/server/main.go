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
	repo := &order.Repository{}
	srv := &order.GRPCServer{IRepository: repo}
	api.RegisterOrderServer(grpcServer, srv)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}