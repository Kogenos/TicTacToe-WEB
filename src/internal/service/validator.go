package service

import (
	"TicTacToe/internal/domain"
	"errors"
)

func (gs *MinimaxGameService) ValidateGame(game domain.Game) error {

	prevBoard := domain.NewBoard()
	for _, move := range game.History {
		if prevBoard[move.Row][move.Col] != domain.CellEmpty {
			return errors.New("invalid history: duplicate move")
		}
		prevBoard[move.Row][move.Col] = move.Player
	}

	if !boardsEqual(prevBoard, game.GameBoard) {
		return errors.New("game board does not match history: previous moves have been altered")
	}

	return nil
}

func boardsEqual(prev, current domain.Board) bool {
	for i := 0; i < domain.Width; i++ {
		for j := 0; j < domain.Height; j++ {
			if prev[i][j] != current[i][j] {
				return false
			}
		}
	}
	return true
}
