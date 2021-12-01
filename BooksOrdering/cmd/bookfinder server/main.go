package main

import (
	api "BooksOrdering/pkg/api"
	"BooksOrdering/pkg/bookfinder"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main (){
	someBooks := []*api.BookInfo{
		{BookId: "123456", BookName: "book1"},
		{BookId: "456789", BookName: "book2"},
		{BookId: "789123", BookName: "book3"},
		{BookId: "123789", BookName: "book4"},
	}
	repo := &bookfinder.Repository{Books: someBooks}
	grpcServer := grpc.NewServer()
	srv := &bookfinder.GRPCServer{BookRepository: repo}
	api.RegisterBookSearchServer(grpcServer, srv)

	listener, err := net.Listen("tcp", ":8070")
	if err != nil {
		log.Fatalf("Book Finder: Ошибка прослушивания: %v", err)
	}
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Book Finder: Ошибка запуска сервера: %v", err)
	}
}
