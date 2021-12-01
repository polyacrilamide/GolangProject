package order

import (
	api "BooksOrdering/pkg/api"
	"context"
	"google.golang.org/grpc"
	"log"
)

type IRepository interface {
	Create(order *api.NewOrder) (*api.NewOrder, error)
	GetAll() []*api.NewOrder
}

type Repository struct {
	orders []*api.NewOrder
}

func (repo *Repository) Create(bookOrder *api.NewOrder) (*api.NewOrder, error) {
	updated := append(repo.orders, bookOrder)
	repo.orders = updated
	return bookOrder, nil
}

func (repo *Repository) GetAll() []*api.NewOrder{
	return repo.orders
}

type GRPCServer struct {
	IRepository IRepository
}

func (s *GRPCServer) MakeOrder(ctx context.Context, req *api.NewOrder) (*api.Response, error) {
	order, err := s.IRepository.Create(req)
	if err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(":8070", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Order: Не удалось подключиться к серверу Book Finder: %v", err)
	}

	client := api.NewBookSearchClient(conn)
	res, err := client.FindTheBook(context.Background(), &api.BookInfo{BookId: req.BookId, BookName: req.Name})
	if err != nil {
		log.Fatalf("Order: Не удалось запросить поиск книги: %v", err)
	}
	log.Printf("Запрос на поиск книги отправлен. Результат: %t", res.Availability)
	return &api.Response{Created: true, NewOrder: order}, nil
}

func (s *GRPCServer) GetOrder (ctx context.Context, req *api.GetRequest) (*api.Response, error){
	orders := s.IRepository.GetAll()
	return &api.Response{NewOrders: orders}, nil
}