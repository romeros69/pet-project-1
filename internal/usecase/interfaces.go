package usecase

import (
	"context"
	"github.com/gofrs/uuid"
	"pet-project-1/internal/entity"
)

type Book interface {
	GetBooks(context.Context) ([]entity.Book, error)
	GetBookById(context.Context, uuid.UUID) (entity.Book, error)
	CreateBook(context.Context, entity.Book) (uuid.UUID, error)
	DeleteBook(context.Context, uuid.UUID) error
	UpdateBook(context.Context, entity.Book) error
}

type BookRepo interface {
	GetBooks(context.Context) ([]entity.Book, error)
	GetBookById(context.Context, uuid.UUID) (entity.Book, error)
	CreateBook(context.Context, entity.Book) (uuid.UUID, error)
	DeleteBook(context.Context, uuid.UUID) error
	UpdateBook(context.Context, entity.Book) error
}
