package order

import (
	api "BooksOrdering/pkg/api"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"log"
)

//MySQL - port: 3306
//root: root
//pass: golangpass

type OrdersList struct {
	Id int
	Name string
	BookId string
	Login string
}

type GRPCServer struct {
}

func (s *GRPCServer) MakeOrder(ctx context.Context, req *api.NewOrder) (*api.Response, error) {
	db, err := sql.Open("mysql", "root:golangpass@tcp(localhost: 3306)/ordersdb")
	if err != nil {
		fmt.Println("Не удается открыть БД")
		panic(err)
	}
	defer db.Close()
	result, err := db.Exec("insert into ordersdb.Orders (bookName, bookId, login) values (?, ?, ?)", req.Name, req.BookId, req.Login)
	if err != nil{
		panic(err)
	}
	fmt.Println("Заказ успешно добавлен.")
	fmt.Println(result.LastInsertId())

	conn, err := grpc.Dial(":8070", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Order: Не удалось подключиться к серверу Book Finder: %v", err)
	}

	client := api.NewBookSearchClient(conn)
	res, err := client.FindTheBook(context.Background(), &api.BookInfo{BookId: req.BookId, BookName: req.Name})
	if err != nil {
		log.Fatalf("Order: Ошибка поиска книги: %v", err)
	}
	log.Printf("Запрос на поиск книги отправлен. Результат: %t", res.Availability)
	return &api.Response{Created: true, NewOrder: req}, nil
}



func (s *GRPCServer) GetOrder (ctx context.Context, req *api.GetRequest) (*api.OrdersListResponse, error){
	db, err := sql.Open("mysql", "root:golangpass@tcp(localhost: 3306)/ordersdb")
	if err != nil {
		fmt.Println("Ошибка открытия в GetOrder")
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from Orders")
	if err != nil {
		fmt.Println("ошибка чтения")
		panic(err)
	}
	defer rows.Close()

	var orders []*api.OrdersList
	var o api.OrdersList

	for rows.Next() {
		o2 := new(OrdersList)
		if err := rows.Scan(&o2.Id, &o2.Name, &o2.BookId, &o2.Login); err != nil {
			fmt.Println("Не удается прочитать строку")
			continue
		}
		o.BookName = o2.Name
		o.BookId = o2.BookId
		o.Login = o2.Login
		orders = append(orders, &o)
	}
	if err := rows.Err(); err != nil {
		log.Fatalln(err)
	}
	return &api.OrdersListResponse{Orderslist: orders}, nil
}