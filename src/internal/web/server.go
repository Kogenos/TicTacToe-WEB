package web

import (
	"context"
	"log"
	"net/http"
	"os"
)

type Server struct {
	handler          *GameHandler
	userHandler      *UserHandler
	JWTauthenticator *JWTAuthenticator
	port             string
	server           *http.Server
}

func NewServer(handler *GameHandler, userHandler *UserHandler, jwtAuthenticator *JWTAuthenticator) *Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Server{
		handler:          handler,
		userHandler:      userHandler,
		JWTauthenticator: jwtAuthenticator,
		port:             port,
	}
}

func (s *Server) SetupRoutes() {

	// Открытые эндпоинты
	http.HandleFunc("POST /register", s.userHandler.Register)
	http.HandleFunc("POST /login", s.userHandler.Login)
	http.HandleFunc("POST /refresh-access", s.userHandler.RefreshAccessToken)
	http.HandleFunc("POST /refresh-refresh", s.userHandler.RefreshRefreshToken)

	// Защищенные эндпоинты
	http.HandleFunc("POST /game", s.JWTauthenticator.Middleware(s.handler.CreateGame))
	http.HandleFunc("POST /game/{id}/move", s.JWTauthenticator.Middleware(s.handler.HandleMove))
	http.HandleFunc("POST /game/{id}/join", s.JWTauthenticator.Middleware(s.handler.JoinGame))
	http.HandleFunc("GET /game/{id}", s.JWTauthenticator.Middleware(s.handler.GetGame))
	http.HandleFunc("GET /games/available", s.JWTauthenticator.Middleware(s.handler.ListAvailableGames))
	http.HandleFunc("GET /games", s.JWTauthenticator.Middleware(s.handler.ListGames))
	http.HandleFunc("GET /users", s.JWTauthenticator.Middleware(s.userHandler.ListUsers))
	http.HandleFunc("GET /user/{id}", s.JWTauthenticator.Middleware(s.userHandler.GetUserByID))
	http.HandleFunc("GET /user/me", s.JWTauthenticator.Middleware(s.userHandler.GetUserByAccessToken))
	http.HandleFunc("GET /games/finished", s.JWTauthenticator.Middleware(s.handler.GetFinishedGames))
	http.HandleFunc("GET /leaderboard", s.JWTauthenticator.Middleware(s.handler.GetLeaderBoard))
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr: ":" + s.port,
	}
	log.Printf("Server starting on port %s", s.port)
	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) {
	log.Printf("Shutting down server")
	s.server.Shutdown(ctx)
}
