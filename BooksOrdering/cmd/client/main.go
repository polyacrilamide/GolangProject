package main

import (
	api "BooksOrdering/pkg/api"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"os"
	
)

func parseFile(file string) (*api.NewOrder, error) {
	var order *api.NewOrder
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &order)
	if err != nil {
		return nil, err
	}
	return order, err
}

func main() {
	file := "OrderedBook.json"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}
	order, err := parseFile(file)
	if err != nil {
		log.Fatalf("Fail to parse file: %v", err)
	}
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Fail to connect: %v", err)
	}
	client := api.NewOrderClient(conn)
	res, err := client.MakeOrder(context.Background(),order)
	if err != nil {
		log.Fatalf("Fail to create new order: %v", err)
	}
	log.Printf("Заказ добавлен: %t", res.Created)

	getAll, err := client.GetOrder(context.Background(), &api.GetRequest{})
	if err != nil {
		log.Fatalf("Failed to get all orders: %v", err)
	}
	for _, ord := range getAll.NewOrders{
		fmt.Printf("Название книги: %v\n", ord.GetName())
		fmt.Printf("Номер книги: %v\n", ord.GetBookId())
		fmt.Printf("ФИО заказчика: %v\n", ord.GetFIO())
		fmt.Printf("Личный номер заказчика: %v\n", ord.GetId())
		fmt.Printf("Пароль заказчика: %v\n", ord.GetPassword())
	}
}