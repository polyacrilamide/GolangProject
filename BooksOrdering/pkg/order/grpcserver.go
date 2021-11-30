package order

import (
	api "BooksOrdering/pkg/api"
	"context"
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
	return &api.Response{Created: true, NewOrder: order}, nil
}

func (s *GRPCServer) GetOrder (ctx context.Context, req *api.GetRequest) (*api.Response, error){
	orders := s.IRepository.GetAll()
	return &api.Response{NewOrders: orders}, nil
}