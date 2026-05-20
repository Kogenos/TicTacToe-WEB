package service

import (
	"TicTacToe/internal/domain"
	"errors"
)

type MinimaxGameService struct{}

func NewMinimaxGameService() *MinimaxGameService {
	return &MinimaxGameService{}
}

func (gs *MinimaxGameService) MakeMove(game domain.Game, row, col int) (domain.Game, error) {

	if row < 0 || row > domain.Width-1 || col < 0 || col > domain.Height-1 {
		return game, errors.New("coordinates out of range")
	}

	finished, _, _ := gs.GameFinished(game)
	if finished {
		return game, errors.New("game already finished")
	}

	if game.GameBoard[row][col] != domain.CellEmpty {
		return game, errors.New("cell is occupied")
	}

	if game.CurrentPlayer != domain.CellX && game.CurrentPlayer != domain.CellO {
		return game, errors.New("invalid current player")
	}

	game.GameBoard[row][col] = game.CurrentPlayer
	game.History = append(game.History, domain.Move{Row: row, Col: col, Player: game.CurrentPlayer})

	if game.CurrentPlayer == domain.CellX {
		game.CurrentPlayer = domain.CellO
	} else {
		game.CurrentPlayer = domain.CellX
	}

	return game, nil
}
