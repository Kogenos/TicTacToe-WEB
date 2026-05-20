package service

import (
	"TicTacToe/internal/domain"
)

type GameService interface {
	GetNextMove(game domain.Game) (int, int, error)
	ValidateGame(game domain.Game) error
	GameFinished(game domain.Game) (bool, domain.Cell, error)
	MakeMove(game domain.Game, row, col int) (domain.Game, error)
}
