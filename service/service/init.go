package service

import (
	"github.com/gin-gonic/gin"
	"github.com/Yhlas9/book_store/config"
	"github.com/Yhlas9/book_store/service/repo"
	"github.com/Yhlas9/book_store/service/structure"
	"log"
)

type Service struct {
	Repo *repo.Repo
}

func NewService(config config.AppConfig) *Service {
	repository := repo.New(config)

	return &Service{
		Repo: repository,
	}
}

func (s *Service) GetBooks(ctx *gin.Context) []structure.Book {
	books, err := s.Repo.GetBooks(ctx)
	if err != nil {
		log.Printf("Error getting books: Error:%v", err)
		return []structure.Book{}
	}

	return books
}

func (s *Service) StoreBook(ctx *gin.Context, book structure.Book) error {
	if err := s.Repo.StoreBook(ctx, book); err != nil {
		log.Printf("Detected error when storeBook Error:%v", err)
		return err
	}

	return nil
}

func (s *Service) DeleteBook(ctx *gin.Context, id int) error {
	if err := s.Repo.DeleteBook(ctx, id); err != nil {
		log.Printf("Detected error when deleteBook Error:%v", err)
		return err
	}

	return nil
}

func (s *Service) UpdateBook(ctx *gin.Context, id int, book structure.Book) error {
	if err := s.Repo.UpdateBook(ctx, id, book); err != nil {
		log.Printf("Detected error when updateBook Error:%v", err)
		return err
	}

	return nil
}

func (s *Service) GetBook(ctx *gin.Context, id int) (structure.Book, error) {
	book, err := s.Repo.GetBook(ctx, id)
	if err != nil {
		log.Printf("Detected error when getBook Error:%v", err)
		return structure.Book{}, err
	}

	return book, nil
}
