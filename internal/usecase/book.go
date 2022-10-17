package usecase

import (
	"context"
	"pet-project-1/internal/entity"
)

type BookService struct {
	repo BookRepo
}

func newBookService(r BookRepo) *BookService {
	return &BookService{repo: r}
}

var _ Book = (*BookService)(nil)

func (b *BookService) GetBooks(ctx context.Context) ([]entity.Book, error) {
	return b.repo.GetBooks(ctx)
}
