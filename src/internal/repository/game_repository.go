package repository

import (
	"TicTacToe/internal/datasource"
	"TicTacToe/internal/domain"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed queries/save_game.sql
var saveGameQuery string

//go:embed queries/load_game.sql
var loadGameQuery string

//go:embed queries/create_table.sql
var createTableQuery string

//go:embed queries/list_games.sql
var listGamesQuery string

//go:embed queries/list_waiting_games.sql
var listWaitingGamesQuery string

//go:embed queries/list_finished_games_by_user.sql
var listFinishedGamesByUserQuery string

//go:embed queries/leaderboard.sql
var leaderboardQuery string

type GameRepositoryImpl struct {
	pool *pgxpool.Pool
}

func NewGameRepositoryImpl(pool *pgxpool.Pool) (*GameRepositoryImpl, error) {
	ctx := context.Background()

	if _, err := pool.Exec(ctx, createTableQuery); err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}

	return &GameRepositoryImpl{pool: pool}, nil
}

func (r *GameRepositoryImpl) Save(game domain.Game) error {
	fmt.Printf("[DEBUG] Сохранение игры с ID: %s\n", game.ID)
	model := datasource.ToDatasourceModel(game)

	boardJSON, _ := json.Marshal(model.GameBoard)
	historyJSON, _ := json.Marshal(model.History)

	_, err := r.pool.Exec(context.Background(), saveGameQuery,
		model.ID, model.PlayerX_ID, model.PlayerO_ID, model.Status,
		model.WinnerID, boardJSON, model.CurrentPlayer,
		historyJSON, model.Bot, model.Botside, model.CreateDate)

	return err
}
func (r *GameRepositoryImpl) Load(id uuid.UUID) (domain.Game, error) {

	var boardJSON, historyJSON []byte
	var model datasource.GameModel

	err := r.pool.QueryRow(context.Background(), loadGameQuery, id).
		Scan(&model.ID, &model.PlayerX_ID, &model.PlayerO_ID, &model.Status,
			&model.WinnerID, &boardJSON, &model.CurrentPlayer,
			&historyJSON, &model.Bot, &model.Botside, &model.CreateDate)

	if err != nil {
		return domain.Game{}, err
	}

	if err := json.Unmarshal(boardJSON, &model.GameBoard); err != nil {
		return domain.Game{}, fmt.Errorf("failed to unmarshal game_board: %w", err)
	}

	if err := json.Unmarshal(historyJSON, &model.History); err != nil {
		return domain.Game{}, fmt.Errorf("failed to unmarshal history: %w", err)
	}

	fmt.Printf("[DEBUG] Загрузка игры с ID: %s\n", id)
	return datasource.ToDomainModel(model), nil
}

func (r *GameRepositoryImpl) ListGames() ([]domain.Game, error) {
	rows, err := r.pool.Query(context.Background(), listGamesQuery)
	if err != nil {
		return nil, fmt.Errorf("query games: %w", err)
	}

	defer rows.Close()

	var games []domain.Game

	for rows.Next() {
		var boardJSON, historyJSON []byte
		var model datasource.GameModel
		var playerXID, playerOID, winnerID uuid.UUID
		var status string

		err := rows.Scan(
			&model.ID,
			&playerXID,
			&playerOID,
			&status,
			&winnerID,
			&boardJSON,
			&model.CurrentPlayer,
			&historyJSON,
			&model.Bot,
			&model.Botside,
			&model.CreateDate,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		model.PlayerX_ID = playerXID
		model.PlayerO_ID = playerOID
		model.Status = status
		model.WinnerID = winnerID

		if err := json.Unmarshal(boardJSON, &model.GameBoard); err != nil {
			return nil, fmt.Errorf("failed to unmarshal game_board: %w", err)
		}
		if err := json.Unmarshal(historyJSON, &model.History); err != nil {
			return nil, fmt.Errorf("failed to unmarshal history: %w", err)
		}

		games = append(games, datasource.ToDomainModel(model))
	}
	return games, rows.Err()
}

func (r *GameRepositoryImpl) ListWaitingGames() ([]domain.Game, error) {
	rows, err := r.pool.Query(context.Background(), listWaitingGamesQuery)
	if err != nil {
		return nil, fmt.Errorf("query waiting games:%w", err)
	}

	defer rows.Close()

	var games []domain.Game
	for rows.Next() {
		var model datasource.GameModel
		var boardJSON, historyJSON []byte

		err := rows.Scan(&model.ID, &model.PlayerX_ID, &model.PlayerO_ID,
			&model.Status, &model.WinnerID, &boardJSON, &model.CurrentPlayer,
			&historyJSON, &model.Bot, &model.Botside, &model.CreateDate)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		if err := json.Unmarshal(boardJSON, &model.GameBoard); err != nil {
			return nil, fmt.Errorf("unmarshal game_board: %w", err)
		}

		if err := json.Unmarshal(historyJSON, &model.History); err != nil {
			return nil, fmt.Errorf("unmarshal history: %w", err)
		}

		games = append(games, datasource.ToDomainModel(model))
	}
	return games, rows.Err()
}

func (r *GameRepositoryImpl) ListFinishedGamesByUser(ctx context.Context, userID uuid.UUID) ([]domain.Game, error) {
	rows, err := r.pool.Query(ctx, listFinishedGamesByUserQuery, userID)
	if err != nil {
		return nil, fmt.Errorf("query finished games: %w", err)
	}
	defer rows.Close()

	var games []domain.Game

	for rows.Next() {

		var boardJSON, historyJSON []byte
		var model datasource.GameModel
		var playerX_ID, playerO_ID, winnerID uuid.UUID
		var status string
		var createDate time.Time

		err := rows.Scan(&model.ID, &playerX_ID, &playerO_ID, &status, &winnerID, &boardJSON,
			&model.CurrentPlayer, &historyJSON, &model.Bot,
			&model.Botside, &createDate)

		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		model.PlayerX_ID = playerX_ID
		model.PlayerO_ID = playerO_ID
		model.Status = status
		model.WinnerID = winnerID
		model.CreateDate = createDate

		if err := json.Unmarshal(boardJSON, &model.GameBoard); err != nil {
			return nil, fmt.Errorf("unmarshal game_board: %w", err)
		}

		if err := json.Unmarshal(historyJSON, &model.History); err != nil {
			return nil, fmt.Errorf("unmarshal history: %w", err)
		}

		games = append(games, datasource.ToDomainModel(model))
	}
	return games, rows.Err()
}

func (r *GameRepositoryImpl) GetLeaderBoard(ctx context.Context, limit int) ([]domain.LeaderBoardEntry, error) {
	rows, err := r.pool.Query(ctx, leaderboardQuery, limit)
	if err != nil {
		return nil, fmt.Errorf("query leaderboard: %w", err)
	}
	defer rows.Close()

	var entries []domain.LeaderBoardEntry
	for rows.Next() {
		var entry domain.LeaderBoardEntry
		err := rows.Scan(&entry.UserID, &entry.Login, &entry.Wins, &entry.TotalGames, &entry.WinRatio)
		if err != nil {
			return nil, fmt.Errorf("scan leaderboard: %w", err)
		}

		entries = append(entries, entry)
	}
	return entries, rows.Err()
}
