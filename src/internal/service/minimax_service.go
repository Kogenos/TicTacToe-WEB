package service

import (
	"TicTacToe/internal/domain"
	"errors"
	"math"
)

const (
	WinScore  = 10
	LoseScore = -10
	DrawScore = 0
)

func (gs *MinimaxGameService) GetNextMove(game domain.Game) (int, int, error) {

	if finished, _, _ := gs.GameFinished(game); finished {
		return -1, -1, errors.New("game already finished")
	}

	if game.CurrentPlayer != game.Botside {
		return -1, -1, errors.New("not bots turn")
	}

	bestScore := math.MinInt
	bestRow, bestCol := -1, -1

	for i := 0; i < domain.Width; i++ {
		for j := 0; j < domain.Height; j++ {
			if game.GameBoard[i][j] == domain.CellEmpty {
				newBoard := game.GameBoard
				newBoard[i][j] = game.CurrentPlayer
				score := minimax(newBoard, switchPlayer(game.CurrentPlayer), game.Botside, 0)

				if score > bestScore {
					bestScore = score
					bestRow, bestCol = i, j
				}
			}
		}
	}

	if bestRow == -1 {
		return -1, -1, errors.New("no available moves")
	}
	return bestRow, bestCol, nil
}

func switchPlayer(player domain.Cell) domain.Cell {
	if player == domain.CellX {
		return domain.CellO
	}
	return domain.CellX

}

func minimax(board domain.Board, currentPlayer, botSide domain.Cell, depth int) int {
	winner := checkWinner(board)

	if winner == botSide {
		return WinScore - depth
	}

	if winner != domain.CellEmpty && winner != botSide {
		return LoseScore + depth
	}

	if isBoardFull(board) {
		return DrawScore
	}

	if currentPlayer == botSide {
		bestScore := math.MinInt
		for i := 0; i < domain.Width; i++ {
			for j := 0; j < domain.Height; j++ {
				if board[i][j] == domain.CellEmpty {
					newBoard := board
					newBoard[i][j] = currentPlayer
					score := minimax(newBoard, switchPlayer(currentPlayer), botSide, depth+1)
					if score > bestScore {
						bestScore = score
					}
				}
			}
		}
		return bestScore
	} else {
		bestScore := math.MaxInt
		for i := 0; i < domain.Width; i++ {
			for j := 0; j < domain.Height; j++ {
				if board[i][j] == domain.CellEmpty {
					newBoard := board
					newBoard[i][j] = currentPlayer
					score := minimax(newBoard, switchPlayer(currentPlayer), botSide, depth+1)
					if score < bestScore {
						bestScore = score
					}
				}
			}
		}
		return bestScore
	}
}
