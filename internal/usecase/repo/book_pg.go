package repo

import (
	"context"
	"fmt"
	"pet-project-1/internal/entity"
	"pet-project-1/internal/usecase"
	"pet-project-1/pkg/postgres"
)

type BookRepo struct {
	pg *postgres.Postgres
}

func NewBookRepo(pg *postgres.Postgres) *BookRepo {
	return &BookRepo{
		pg: pg,
	}
}

var _ usecase.BookRepo = (*BookRepo)(nil)

func (b *BookRepo) GetBooks(ctx context.Context) ([]entity.Book, error) {
	query := `SELECT * FROM book`

	rows, err := b.pg.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()

	var books []entity.Book

	for rows.Next() {
		var book entity.Book
		err = rows.Scan(&book.ID,
			&book.Tittle,
			&book.Author)
		if err != nil {
			return nil, fmt.Errorf("error in parsing book: %w", err)
		}
		books = append(books, book)
	}
	return books, nil
}

func (b *BookRepo) CreateBook(ctx context.Context, book entity.Book) error {
	query := `INSERT INTO book (tittle, author) VALUES ($1, $2)`

	rows, err := b.pg.Pool.Query(ctx, query, book.Tittle, book.Author)

	if err != nil {
		return fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()

	return nil
}

func (b *BookRepo) DeleteBook(ctx context.Context, ID string) error {
	query := `DELETE FROM book WHERE id = $1`

	rows, err := b.pg.Pool.Query(ctx, query, ID)

	if err != nil {
		return fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()

	return nil
}

func (b *BookRepo) UpdateBook(ctx context.Context, book entity.Book) error {
	query := `UPDATE book SET tittle=$1, author=$2 WHERE id=$3`

	rows, err := b.pg.Pool.Query(ctx, query, book.Tittle, book.Author, book.ID)

	if err != nil {
		return fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()

	return nil
}
