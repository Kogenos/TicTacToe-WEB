package domain

import (
	"time"

	"github.com/google/uuid"
)

type GameStatus string

const (
	StatusWaiting  GameStatus = "waiting"
	StatusPlaying  GameStatus = "playing"
	StatusFinished GameStatus = "finished"
)

type Game struct {
	ID            uuid.UUID
	PlayerX_ID    uuid.UUID
	PlayerO_ID    uuid.UUID
	Status        GameStatus
	WinnerID      uuid.UUID
	GameBoard     Board
	CurrentPlayer Cell
	History       []Move
	Bot           bool
	Botside       Cell
	CreateDate    time.Time
}

func NewGame(playerX_ID uuid.UUID) *Game {
	return &Game{
		ID:            uuid.New(),
		PlayerX_ID:    playerX_ID,
		PlayerO_ID:    uuid.Nil,
		Status:        StatusWaiting,
		WinnerID:      uuid.Nil,
		GameBoard:     NewBoard(),
		CurrentPlayer: CellX,
		History:       []Move{},
		Bot:           false,
		Botside:       CellEmpty,
		CreateDate:    time.Now(),
	}
}

func NewGameWithBot(playerX_ID uuid.UUID, botSide Cell) *Game {
	return &Game{
		ID:            uuid.New(),
		PlayerX_ID:    playerX_ID,
		PlayerO_ID:    uuid.Nil,
		Status:        StatusPlaying,
		WinnerID:      uuid.Nil,
		GameBoard:     NewBoard(),
		CurrentPlayer: CellX,
		History:       []Move{},
		Bot:           true,
		Botside:       botSide,
		CreateDate:    time.Now(),
	}
}
