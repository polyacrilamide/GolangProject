package bookfinder

import (
	api "BooksOrdering/pkg/api"
	"context"
	"errors"
	"log"
)

type BookRepository interface {
	Find (book *api.BookInfo) (bool, error)
}

type Repository struct {
	Books []*api.BookInfo
}

func (repo *Repository) Find (book *api.BookInfo) (bool, error){
	for _, bookID := range repo.Books {
		if book.BookId == bookID.BookId {
			log.Println("Book finder: Запрашиваемая книга найдена.")
			return true, nil
		}
	}
	log.Printf("Запрашиваемой книги с номером %s не существует.\n", book.BookId)
	log.Println("Поиск по названию:")
	for _, bookName := range repo.Books{
		if book.BookName == bookName.BookName{
			log.Println("Book finder: Запрашиваемая книга найдена.")
			return true, nil
		}
	}
	log.Println("Запрашиваемой книги не существует.")
	return false, errors.New("книги не существует")
}

type GRPCServer struct {
	BookRepository BookRepository
}

func (s *GRPCServer) FindTheBook (ctx context.Context, book *api.BookInfo) (*api.Answer, error){
	finder, err := s.BookRepository.Find(book)
	if err != nil {
		return nil, err
	}
	return &api.Answer{Availability: finder}, nil
}