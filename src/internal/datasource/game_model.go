package datasource

import (
	"time"

	"github.com/google/uuid"
)

type GameModel struct {
	ID            uuid.UUID   `db:"id" json:"id"`
	PlayerX_ID    uuid.UUID   `db:"player_x_id" json:"player_x_id"`
	PlayerO_ID    uuid.UUID   `db:"player_o_id" json:"player_o_id"`
	Status        string      `db:"status" json:"status"`
	WinnerID      uuid.UUID   `db:"winner_id" json:"winner_id"`
	GameBoard     [3][3]int   `db:"game_board" json:"game_board"`
	CurrentPlayer int         `db:"current_player" json:"current_player"`
	History       []MoveModel `db:"history" json:"history"`
	Bot           bool        `db:"bot" json:"bot"`
	Botside       int         `db:"bot_side" json:"bot_side"`
	CreateDate    time.Time   `db:"create_date" json:"create_date"`
}

type MoveModel struct {
	Row    int `json:"row"`
	Col    int `json:"col"`
	Player int `json:"player"`
}
