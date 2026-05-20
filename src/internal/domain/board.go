package domain

const (
	Width  = 3
	Height = 3
)

type Board [Width][Height]Cell

func NewBoard() Board {

	var board Board

	for i := range board {
		for j := range board[i] {
			board[i][j] = CellEmpty
		}
	}
	return board
}

func (b *Board) SetCell(i, j int, cell Cell) {
	b[i][j] = cell
}

func (b *Board) GetCell(i, j int) Cell {
	return b[i][j]
}

func (b *Board) IsEmptyCell(i, j int) bool {
	return b[i][j] == CellEmpty
}
