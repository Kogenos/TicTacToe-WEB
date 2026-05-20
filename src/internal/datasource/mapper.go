package datasource

import "TicTacToe/internal/domain"

func ToDomainModel(model GameModel) domain.Game {

	var board domain.Board
	for i := 0; i < domain.Width; i++ {
		for j := 0; j < domain.Height; j++ {
			board[i][j] = domain.Cell(model.GameBoard[i][j])
		}
	}

	history := make([]domain.Move, len(model.History))
	for i, m := range model.History {
		history[i] = domain.Move{
			Row:    m.Row,
			Col:    m.Col,
			Player: domain.Cell(m.Player),
		}
	}

	return domain.Game{
		ID:            model.ID,
		PlayerX_ID:    model.PlayerX_ID,
		PlayerO_ID:    model.PlayerO_ID,
		Status:        domain.GameStatus(model.Status),
		WinnerID:      model.WinnerID,
		GameBoard:     board,
		CurrentPlayer: domain.Cell(model.CurrentPlayer),
		History:       history,
		Bot:           model.Bot,
		Botside:       domain.Cell(model.Botside),
		CreateDate:    model.CreateDate,
	}
}

func ToDatasourceModel(game domain.Game) GameModel {

	var board [domain.Width][domain.Height]int
	for i := 0; i < domain.Width; i++ {
		for j := 0; j < domain.Height; j++ {
			board[i][j] = int(game.GameBoard[i][j])
		}
	}

	history := make([]MoveModel, len(game.History))
	for i, m := range game.History {
		history[i] = MoveModel{
			Row:    m.Row,
			Col:    m.Col,
			Player: int(m.Player),
		}
	}

	return GameModel{
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
