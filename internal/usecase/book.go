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
