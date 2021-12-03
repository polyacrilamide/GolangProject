package main

import (
	api "BooksOrdering/pkg/api"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
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
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Не удалось подключиться к серверу: %v", err)
	}
	client := api.NewOrderClient(conn)

	var i int
	fmt.Println("Выберите действие, которое хотите совершить (введите число): ")
	fmt.Println("1: Регистрация; 2: Авторизация; 3: Просмотр доступных книг")
	fmt.Scanf("%d\n", &i)

	switch i {
	case 1:
		{

		}
	case 2:
		{
			var username, password string
			fmt.Println("Введите логин и пароль: ")
			fmt.Scanf("%s\n", &username)
			fmt.Scanf("%s\n", &password)

			clientHttp := &http.Client{
				Timeout: time.Second * 10,
			}
			req, err := http.NewRequest("POST", "https://localhost:8060/authserver", nil)
			if err != nil {
				fmt.Println("не удалось подключиться к серверу.")
			}
			req.SetBasicAuth(username, password)
			response, err := clientHttp.Do(req)
			if err != nil {
				fmt.Println("Ошибка в запросе http")
				log.Fatal(err)
			}
			defer response.Body.Close()

			fmt.Println("Выберите действие, которое хотите совершить: ")
			fmt.Println("1: Сделать заказ; 2: Посмотреть доступные книги")

			var i_2 int

			fmt.Scan(i_2)

			switch i_2 {
			case 1:
				{
					file := "OrderedBook.json"
					if len(os.Args) > 1 {
						file = os.Args[1]
					}
					order, err := parseFile(file)
					if err != nil {
						log.Fatalf("Не удалось распарсить файл: %v", err)
					}

					res, err := client.MakeOrder(context.Background(), order)
					if err != nil {
						log.Fatalf("Не удалось создать новый заказ: %v", err)
					}
					log.Printf("Заказ добавлен: %t", res.Created)
				}
			case 2:
				{
					getAll, err := client.GetOrder(context.Background(), &api.GetRequest{})
					if err != nil {
						log.Fatalf("Не удалось получить список всех заказов: %v", err)
					}
					for _, ord := range getAll.Orderslist {
						fmt.Printf("Название книги: %v\n", ord.GetBookName())
						fmt.Printf("Номер книги: %v\n", ord.GetBookId())
						fmt.Printf("Логин заказчика: %v\n", ord.GetLogin())
						fmt.Println()
					}
				}
			}
		}
	case 3:
		{
			getAll, err := client.GetOrder(context.Background(), &api.GetRequest{})
			if err != nil {
				log.Fatalf("Не удалось получить список всех заказов: %v", err)
			}
			for _, ord := range getAll.Orderslist {
				fmt.Printf("Название книги: %v\n", ord.GetBookName())
				fmt.Printf("Номер книги: %v\n", ord.GetBookId())
				fmt.Printf("Логин заказчика: %v\n", ord.GetLogin())
				fmt.Println()
			}
		}
	}
}