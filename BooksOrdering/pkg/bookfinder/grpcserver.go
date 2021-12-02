package bookfinder

import (
	api "BooksOrdering/pkg/api"
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Find (book *api.BookInfo) (bool, error){
	db, err := sql.Open("mysql", "root:golangpass@tcp(localhost: 3306)/booksdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from Books where bookId = ?", book.BookId)
	if err != nil {
		panic(err)
	}

	var bookName, bookId string

	for rows.Next() {
		if err := rows.Scan(bookName, bookId); err == nil{
			log.Println("Book finder: Запрашиваемая книга найдена.")
			return true, nil
		}
	}
	rows, err = db.Query("select * from Books where bookName = ?", book.BookName)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(bookName, bookId); err == nil{
			log.Println("Book finder: Запрашиваемая книга найдена.")
			return true, nil
		}
	}
	log.Println("Запрашиваемой книги не существует.")
	return false, nil
}

type GRPCServer struct {}

func (s *GRPCServer) FindTheBook (ctx context.Context, book *api.BookInfo) (*api.Answer, error){
	finder, err := Find(book)
	if err != nil {
		return nil, err
	}
	return &api.Answer{Availability: finder}, nil
}