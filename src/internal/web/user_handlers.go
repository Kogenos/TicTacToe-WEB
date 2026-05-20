package web

import (
	"TicTacToe/internal/service"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type UserHandler struct {
	userService service.UserService
	authService service.AuthService
}

func NewUserHandler(userService service.UserService, authService service.AuthService) *UserHandler {
	return &UserHandler{
		userService: userService,
		authService: authService,
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req SignUpRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	userID, err := h.userService.Register(r.Context(), req.Login, req.Password)
	if err != nil {
		sendError(w, err.Error(), http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"user_id": userID.String()})
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req JwtRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	accessToken, refreshToken, err := h.authService.Authenticate(r.Context(), req.Login, req.Password)
	if err != nil {
		sendError(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(JwtResponse{
		Type:         "Bearer",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (h *UserHandler) RefreshAccessToken(w http.ResponseWriter, r *http.Request) {
	var req RefreshJwtRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	newAccessToken, oldRefreshToken, err := h.authService.RefreshAccessToken(r.Context(), req.RefreshToken)
	if err != nil {
		sendError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(JwtResponse{
		Type:         "Bearer",
		AccessToken:  newAccessToken,
		RefreshToken: oldRefreshToken,
	})

}

func (h *UserHandler) RefreshRefreshToken(w http.ResponseWriter, r *http.Request) {
	var req RefreshJwtRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	newAccessToken, newRefreshToken, err := h.authService.RefreshRefreshToken(r.Context(), req.RefreshToken)
	if err != nil {
		sendError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(JwtResponse{
		Type:         "Bearer",
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	})

}

func (h *UserHandler) GetUserByAccessToken(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		sendError(w, "authorization required", http.StatusUnauthorized)
		return
	}

	const prefix = "Bearer "

	if !strings.HasPrefix(authHeader, prefix) {
		sendError(w, "invalid auth scheme", http.StatusBadRequest)
		return
	}

	tokenString := authHeader[len(prefix):]
	user, err := h.authService.GetUserByAccessToken(r.Context(), tokenString)

	if err != nil {
		sendError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":    user.ID,
		"login": user.Login,
	})

}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.ListAllUsers(r.Context())
	if err != nil {
		sendError(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	type UserResponse struct {
		ID    uuid.UUID `json:"id"`
		Login string    `json:"login"`
	}

	resp := make([]UserResponse, len(users))
	for i, u := range users {
		resp[i] = UserResponse{ID: u.ID, Login: u.Login}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := r.PathValue("id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		sendError(w, "Invalid user UUID", http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUserByID(r.Context(), userID)
	if err != nil {
		sendError(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":    user.ID,
		"login": user.Login,
	})
}
