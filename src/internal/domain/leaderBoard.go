package domain

import "github.com/google/uuid"

type LeaderBoardEntry struct {
	UserID     uuid.UUID `json:"user_id"`
	Login      string    `json:"login"`
	Wins       int       `json:"wins"`
	TotalGames int       `json:"total_games"`
	WinRatio   float64   `json:"win_ratio"`
}
