package web

import (
	"TicTacToe/internal/domain"
	"time"

	"github.com/google/uuid"
)

type GameDTO struct {
	ID            uuid.UUID                        `json:"id"`
	PlayerX_ID    uuid.UUID                        `json:"player_x_id"`
	PlayerO_ID    uuid.UUID                        `json:"player_o_id"`
	Status        string                           `json:"status"`
	WinnerID      uuid.UUID                        `json:"winner_id"`
	GameBoard     [domain.Width][domain.Height]int `json:"game_board"`
	CurrentPlayer int                              `json:"current_player"`
	History       []MoveDTO                        `json:"history"`
	Bot           bool                             `json:"bot"`
	Botside       int                              `json:"botside"`
	CreateDate    time.Time                        `json:"create_date"`
}

type MoveDTO struct {
	Row    int `json:"row"`
	Col    int `json:"col"`
	Player int `json:"player"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SignUpRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type JwtRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type JwtResponse struct {
	Type         string `json:"type"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshJwtRequest struct {
	RefreshToken string `json:"refresh_token"`
}
