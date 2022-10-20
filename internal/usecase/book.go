package usecase

import (
	"context"
	"pet-project-1/internal/entity"
)

type BookUseCase struct {
	repo BookRepo
}

func NewBookUseCase(r BookRepo) *BookUseCase {
	return &BookUseCase{repo: r}
}

var _ Book = (*BookUseCase)(nil)

func (b *BookUseCase) GetBooks(ctx context.Context) ([]entity.Book, error) {
	return b.repo.GetBooks(ctx)
}

func (b *BookUseCase) CreateBook(ctx context.Context, book entity.Book) error {
	return b.repo.CreateBook(ctx, book)
}

func (b *BookUseCase) DeleteBook(ctx context.Context, ID string) error {
	return b.repo.DeleteBook(ctx, ID)
}
