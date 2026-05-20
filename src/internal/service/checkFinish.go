package service

import "TicTacToe/internal/domain"

func (gs *MinimaxGameService) GameFinished(game domain.Game) (bool, domain.Cell, error) {

	winner := checkWinner(game.GameBoard)
	if winner != domain.CellEmpty {
		return true, winner, nil
	}

	if isBoardFull(game.GameBoard) {
		return true, domain.CellEmpty, nil
	}

	return false, domain.CellEmpty, nil

}

func checkWinner(board domain.Board) domain.Cell {

	for i := 0; i < domain.Width; i++ {
		if board[i][0] != domain.CellEmpty && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return board[i][0]
		}
	}

	for j := 0; j < domain.Height; j++ {
		if board[0][j] != domain.CellEmpty && board[0][j] == board[1][j] && board[1][j] == board[2][j] {
			return board[0][j]
		}
	}

	if board[0][0] != domain.CellEmpty && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return board[0][0]
	}

	if board[0][2] != domain.CellEmpty && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return board[0][2]
	}

	return domain.CellEmpty
}

func isBoardFull(board domain.Board) bool {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == domain.CellEmpty {
				return false
			}
		}
	}
	return true
}
