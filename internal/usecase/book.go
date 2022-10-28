package usecase

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
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

func (b *BookUseCase) GetBookById(ctx context.Context, id uuid.UUID) (entity.Book, error) {
	return b.repo.GetBookById(ctx, id)
}

func (b *BookUseCase) CreateBook(ctx context.Context, book entity.Book) (uuid.UUID, error) {
	return b.repo.CreateBook(ctx, book)
}

func (b *BookUseCase) DeleteBook(ctx context.Context, id uuid.UUID) error {
	_, err := b.repo.GetBookById(ctx, id)
	if err != nil {
		return fmt.Errorf("there is no book with this id")
	}
	return b.repo.DeleteBook(ctx, id)
}

func (b *BookUseCase) UpdateBook(ctx context.Context, book entity.Book) error {
	_, err := b.repo.GetBookById(ctx, book.ID)
	if err != nil {
		return fmt.Errorf("there is no book with this id")
	}
	return b.repo.UpdateBook(ctx, book)
}
