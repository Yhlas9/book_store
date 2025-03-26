package repo

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/Yhlas9/book_store/config"
	"github.com/Yhlas9/book_store/pkg/database/psql"
	"github.com/Yhlas9/book_store/service/structure"
)

type Repo struct {
	DB *pgxpool.Pool
}

func New(config config.AppConfig) *Repo {
	return &Repo{
		DB: psql.Init(config),
	}
}

func (r *Repo) GetBooks(ctx context.Context) ([]structure.Book, error) {
	var books []structure.Book

	query, err := r.DB.Query(ctx, "SELECT id, title, author FROM books")
	if err != nil {
		return nil, err
	}

	for query.Next() {
		var book structure.Book
		if err = query.Scan(&book.Id, &book.Title, &book.Author); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (r *Repo) StoreBook(ctx context.Context, book structure.Book) error {
	if _, err := r.DB.Exec(ctx, "INSERT INTO books(title, author) values ($1,$2)", book.Title, book.Author); err != nil {
		return err
	}

	return nil
}

func (r *Repo) DeleteBook(ctx context.Context, id int) error {
	if _, err := r.DB.Exec(ctx, "DELETE FROM books WHERE id = $1", id); err != nil {
		return err
	}

	return nil
}

func (r *Repo) UpdateBook(ctx context.Context, id int, book structure.Book) error {
	if _, err := r.DB.Exec(ctx, "UPDATE  books SET title= $1,author =$2  WHERE id = $3 ", book.Title, book.Author, id); err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetBook(ctx context.Context, id int) (structure.Book, error) {
	var book structure.Book
	if err := r.DB.QueryRow(ctx, "SELECT id, title, author FROM books WHERE id = $1", id).Scan(&book.Id, &book.Title, &book.Author); err != nil {
		return book, err
	}

	return book, nil
}
