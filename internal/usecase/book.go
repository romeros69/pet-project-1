package usecase

import (
	"context"
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

func (b *BookUseCase) CreateBook(ctx context.Context, book entity.Book) (uuid.UUID, error) {
	return b.repo.CreateBook(ctx, book)
}

func (b *BookUseCase) DeleteBook(ctx context.Context, ID string) error {
	return b.repo.DeleteBook(ctx, ID)
}

func (b *BookUseCase) UpdateBook(ctx context.Context, book entity.Book) error {
	return b.repo.UpdateBook(ctx, book)
}
