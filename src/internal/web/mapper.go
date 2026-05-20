package web

import "TicTacToe/internal/domain"

func ToDomainGame(dto GameDTO) domain.Game {

	var board domain.Board
	for i := 0; i < domain.Width; i++ {
		for j := 0; j < domain.Height; j++ {
			board[i][j] = domain.Cell(dto.GameBoard[i][j])
		}
	}

	history := make([]domain.Move, len(dto.History))
	for i, m := range dto.History {
		history[i] = domain.Move{
			Row:    m.Row,
			Col:    m.Col,
			Player: domain.Cell(m.Player),
		}
	}

	return domain.Game{
		ID:            dto.ID,
		GameBoard:     board,
		CurrentPlayer: domain.Cell(dto.CurrentPlayer),
		History:       history,
		Bot:           dto.Bot,
		Botside:       domain.Cell(dto.Botside),
	}
}

func ToGameDTO(game domain.Game) GameDTO {

	var board [domain.Width][domain.Height]int
	for i := 0; i < domain.Width; i++ {
		for j := 0; j < domain.Height; j++ {
			board[i][j] = int(game.GameBoard[i][j])
		}
	}

	history := make([]MoveDTO, len(game.History))
	for i, m := range game.History {
		history[i] = MoveDTO{
			Row:    m.Row,
			Col:    m.Col,
			Player: int(m.Player),
		}
	}

	return GameDTO{
		ID:            game.ID,
		PlayerX_ID:    game.PlayerX_ID,
		PlayerO_ID:    game.PlayerO_ID,
		Status:        string(game.Status),
		WinnerID:      game.WinnerID,
		GameBoard:     board,
		CurrentPlayer: int(game.CurrentPlayer),
		History:       history,
		Bot:           game.Bot,
		Botside:       int(game.Botside),
		CreateDate:    game.CreateDate,
	}
}
