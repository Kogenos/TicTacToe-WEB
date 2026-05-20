package repository

import (
	"TicTacToe/internal/domain"
	"context"

	"github.com/google/uuid"
)

type GameRepository interface {
	Save(game domain.Game) error
	Load(id uuid.UUID) (domain.Game, error)
	ListGames() ([]domain.Game, error)
	ListWaitingGames() ([]domain.Game, error)
	ListFinishedGamesByUser(ctx context.Context, userID uuid.UUID) ([]domain.Game, error)
	GetLeaderBoard(ctx context.Context, limit int) ([]domain.LeaderBoardEntry, error)
}
