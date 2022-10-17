package usecase

import (
	"context"
	"pet-project-1/internal/entity"
)

type Book interface {
	GetBooks(context.Context) ([]entity.Book, error)
}

type BookRepo interface {
	GetBooks(context.Context) ([]entity.Book, error)
}
