package serviceWithRepos

import (
	"TicTacToe/internal/domain"
	"TicTacToe/internal/repository"
	"TicTacToe/internal/service"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type GameServiceWithRepo struct {
	repo  repository.GameRepository
	logic *service.MinimaxGameService
}

func NewGameServiceWithRepo(repo repository.GameRepository, logic *service.MinimaxGameService) *GameServiceWithRepo {
	return &GameServiceWithRepo{
		repo:  repo,
		logic: logic,
	}
}

func (s *GameServiceWithRepo) MakeMove(game domain.Game, row, col int) (domain.Game, error) {
	newGame, err := s.logic.MakeMove(game, row, col)
	if err != nil {
		return newGame, err
	}

	if saveErr := s.repo.Save(newGame); saveErr != nil {
		return newGame, saveErr
	}
	return newGame, nil
}

func (s *GameServiceWithRepo) MakeMovePlayer(ctx context.Context, gameID, playerID uuid.UUID, row, col int) (domain.Game, error) {

	game, err := s.repo.Load(gameID)
	if err != nil {
		return domain.Game{}, err
	}

	if game.Status != domain.StatusPlaying {
		return domain.Game{}, errors.New("game not in progress")
	}

	expPlayer := game.CurrentPlayer
	var allowedID uuid.UUID
	if expPlayer == domain.CellX {
		allowedID = game.PlayerX_ID
	} else {
		allowedID = game.PlayerO_ID
	}

	if playerID != allowedID {
		return domain.Game{}, errors.New("not your turn")
	}

	newGame, err := s.logic.MakeMove(game, row, col)
	if err != nil {
		return newGame, err
	}

	finished, winner, _ := s.logic.GameFinished(newGame)
	if finished {
		newGame.Status = domain.StatusFinished
		if winner == domain.CellX {
			newGame.WinnerID = newGame.PlayerX_ID
		} else if winner == domain.CellO {
			newGame.WinnerID = newGame.PlayerO_ID

		} else {
			newGame.WinnerID = uuid.Nil
		}
	}

	if err := s.repo.Save(newGame); err != nil {
		return domain.Game{}, err
	}
	return newGame, nil
}

func (s *GameServiceWithRepo) MakeMoveBot(ctx context.Context, gameID uuid.UUID) (domain.Game, error) {
	game, err := s.repo.Load(gameID)
	if err != nil {
		return domain.Game{}, err
	}

	if !game.Bot {
		return domain.Game{}, errors.New("not a bot game")
	}

	if game.Status != domain.StatusPlaying {
		return domain.Game{}, errors.New("game not in progress")
	}

	if game.CurrentPlayer != game.Botside {
		return domain.Game{}, errors.New("not bots turn")
	}

	row, col, err := s.logic.GetNextMove(game)
	if err != nil {
		return domain.Game{}, err
	}

	newGame, err := s.logic.MakeMove(game, row, col)
	if err != nil {
		return domain.Game{}, err
	}

	finished, winner, _ := s.logic.GameFinished(newGame)
	if finished {
		newGame.Status = domain.StatusFinished
		if winner == domain.CellX {
			newGame.WinnerID = newGame.PlayerX_ID
		} else if winner == domain.CellO {
			newGame.WinnerID = newGame.PlayerO_ID

		} else {
			newGame.WinnerID = uuid.Nil
		}

	}

	if err := s.repo.Save(newGame); err != nil {
		return domain.Game{}, err
	}
	return newGame, nil
}

func (s *GameServiceWithRepo) GameFinished(game domain.Game) (bool, domain.Cell, error) {
	return s.logic.GameFinished(game)
}

func (s *GameServiceWithRepo) ValidateGame(game domain.Game) error {
	return s.logic.ValidateGame(game)
}

func (s *GameServiceWithRepo) GetNextMove(game domain.Game) (int, int, error) {
	return s.logic.GetNextMove(game)
}

func (s *GameServiceWithRepo) CreateGame(ctx context.Context, creatorID uuid.UUID, isBot bool, botside domain.Cell) (domain.Game, error) {
	var game *domain.Game
	if isBot {
		game = domain.NewGameWithBot(creatorID, botside)
	} else {
		game = domain.NewGame(creatorID)
	}

	if isBot && botside == domain.CellX {
		row, col, err := s.logic.GetNextMove(*game)
		if err != nil {
			return domain.Game{}, fmt.Errorf("bot cannot make first move: %w", err)
		}
		newGame, err := s.logic.MakeMove(*game, row, col)
		if err != nil {
			return domain.Game{}, fmt.Errorf("bot first move failed: %w", err)
		}
		game = &newGame
	}

	if err := s.repo.Save(*game); err != nil {
		return domain.Game{}, err
	}
	return *game, nil
}

func (s *GameServiceWithRepo) JoinGame(ctx context.Context, gameID, playerID uuid.UUID) (domain.Game, error) {
	game, err := s.repo.Load(gameID)
	if err != nil {
		return domain.Game{}, err
	}

	if game.Status != domain.StatusWaiting {
		return domain.Game{}, errors.New("game is not waiting for players")
	}

	if game.Bot {
		return domain.Game{}, errors.New("cannot join bot game")
	}

	if game.PlayerX_ID == playerID || game.PlayerO_ID == playerID {
		return domain.Game{}, errors.New("player is already in game")
	}

	if game.PlayerO_ID == uuid.Nil {
		game.PlayerO_ID = playerID
	} else {
		return domain.Game{}, errors.New("game already has two players")
	}

	game.Status = domain.StatusPlaying
	if err := s.repo.Save(game); err != nil {
		return domain.Game{}, err
	}

	return game, nil
}

func (s *GameServiceWithRepo) GetAvailableGames(ctx context.Context) ([]domain.Game, error) {
	return s.repo.ListWaitingGames()
}

func (s *GameServiceWithRepo) GetFinishedGamesByUser(ctx context.Context, userID uuid.UUID) ([]domain.Game, error) {
	return s.repo.ListFinishedGamesByUser(ctx, userID)
}

func (s *GameServiceWithRepo) GetLeaderBoard(ctx context.Context, limit int) ([]domain.LeaderBoardEntry, error) {
	return s.repo.GetLeaderBoard(ctx, limit)
}
