package web

import (
	"TicTacToe/internal/domain"
	"TicTacToe/internal/repository"
	"TicTacToe/internal/serviceWithRepos"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

type GameHandler struct {
	gameService *serviceWithRepos.GameServiceWithRepo
	repo        repository.GameRepository
}

func NewGameHandler(gameService *serviceWithRepos.GameServiceWithRepo, repo repository.GameRepository) *GameHandler {
	return &GameHandler{
		gameService: gameService,
		repo:        repo,
	}
}

func (h *GameHandler) CreateGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := r.Context().Value(UserIDKey).(uuid.UUID)
	if !ok {
		sendError(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var req struct {
		Bot     bool   `json:"bot"`
		Botside string `json:"botside"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	var botSide domain.Cell
	if req.Bot {
		switch req.Botside {
		case "X":
			botSide = domain.CellX
		case "O":
			botSide = domain.CellO
		default:
			sendError(w, "botside must be X or O", http.StatusBadRequest)
			return
		}
	} else {
		botSide = domain.CellEmpty
	}

	game, err := h.gameService.CreateGame(r.Context(), userID, req.Bot, botSide)
	if err != nil {
		sendError(w, "Failed to create game", http.StatusInternalServerError)
		return
	}

	responseDTO := ToGameDTO(game)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseDTO)
}

func (h *GameHandler) GetGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	gameIDStr := r.PathValue("id")
	gameID, err := uuid.Parse(gameIDStr)
	if err != nil {
		sendError(w, "Invalid game UUID", http.StatusBadRequest)
		return
	}

	game, err := h.repo.Load(gameID)
	if err != nil {
		sendError(w, "Game not found", http.StatusNotFound)
		return
	}

	responseDTO := ToGameDTO(game)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseDTO)
}

func (h *GameHandler) HandleMove(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	gameIDStr := r.PathValue("id")
	gameID, err := uuid.Parse(gameIDStr)
	if err != nil {
		sendError(w, "Invalid game UUID", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value(UserIDKey).(uuid.UUID)
	if !ok {
		sendError(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var moveDTO MoveDTO
	if err := json.NewDecoder(r.Body).Decode(&moveDTO); err != nil {
		sendError(w, "Invalid Json body", http.StatusBadRequest)
		return
	}

	updatedGame, err := h.gameService.MakeMovePlayer(r.Context(), gameID, userID, moveDTO.Row, moveDTO.Col)
	if err != nil {
		sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if updatedGame.Bot && updatedGame.Status == domain.StatusPlaying {
		if updatedGame.CurrentPlayer == updatedGame.Botside {
			updatedGame, err = h.gameService.MakeMoveBot(r.Context(), gameID)
			if err != nil {
				sendError(w, "Bot move failed", http.StatusInternalServerError)
				return
			}
		}
	}

	responseDTO := ToGameDTO(updatedGame)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseDTO)
}

func sendError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}

func (h *GameHandler) ListGames(w http.ResponseWriter, r *http.Request) {
	games, err := h.repo.ListGames()

	if err != nil {
		sendError(w, "Failed to fetch games", http.StatusInternalServerError)
		return
	}

	dtos := make([]GameDTO, len(games))

	for i, g := range games {
		dtos[i] = ToGameDTO(g)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dtos)
}

func (h *GameHandler) JoinGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	gameIDStr := r.PathValue("id")
	gameID, err := uuid.Parse(gameIDStr)
	if err != nil {
		sendError(w, "Invalid game UUID", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value(UserIDKey).(uuid.UUID)
	if !ok {
		sendError(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	updatedGame, err := h.gameService.JoinGame(r.Context(), gameID, userID)
	if err != nil {
		sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseDTO := ToGameDTO(updatedGame)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseDTO)
}

func (h *GameHandler) ListAvailableGames(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	games, err := h.gameService.GetAvailableGames(r.Context())
	if err != nil {
		sendError(w, "failed to fetch games", http.StatusInternalServerError)
		return
	}

	dtos := make([]GameDTO, len(games))

	for i, g := range games {
		dtos[i] = ToGameDTO(g)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dtos)
}

func (h *GameHandler) GetFinishedGames(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := r.Context().Value(UserIDKey).(uuid.UUID)
	if !ok {
		sendError(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	games, err := h.gameService.GetFinishedGamesByUser(r.Context(), userID)
	if err != nil {
		sendError(w, "Failed to fetch finished games", http.StatusInternalServerError)
		return
	}

	dtos := make([]GameDTO, len(games))
	for i, g := range games {
		dtos[i] = ToGameDTO(g)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dtos)
}

func (h *GameHandler) GetLeaderBoard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	limitStr := r.URL.Query().Get("limit")
	if limitStr == "" {
		limitStr = "10"
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		sendError(w, "invalid limit parameter", http.StatusBadRequest)
		return
	}

	if limit > 100 {
		limit = 100
	}

	entries, err := h.gameService.GetLeaderBoard(r.Context(), limit)
	if err != nil {
		sendError(w, "Failed to fetch leaderboard", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entries)
}
