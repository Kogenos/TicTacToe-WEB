package repository

import (
	"TicTacToe/internal/domain"
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User) error
	FindByLogin(ctx context.Context, login string) (domain.User, error)
	ListUsers(ctx context.Context) ([]domain.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (domain.User, error)
}
