package service

import (
	"TicTacToe/internal/domain"
	"TicTacToe/internal/repository"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	Authenticate(ctx context.Context, login, password string) (accessToken, refreshToken string, err error)
	RefreshAccessToken(ctx context.Context, refreshToken string) (newAccessToken, oldRefreshToken string, err error)
	RefreshRefreshToken(ctx context.Context, refreshToken string) (newAccessToken, newRefreshToken string, err error)
	GetUserByAccessToken(ctx context.Context, accessToken string) (domain.User, error)
}

type AuthServiceImpl struct {
	userService   UserService
	jwtProvider   *JwtProvider
	blacklistRepo repository.BlacklistRepository
}

func NewAuthService(userService UserService, jwtProvider *JwtProvider, blacklistRepo repository.BlacklistRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		userService:   userService,
		jwtProvider:   jwtProvider,
		blacklistRepo: blacklistRepo,
	}
}

func (s *AuthServiceImpl) Authenticate(ctx context.Context, login, password string) (string, string, error) {
	userID, err := s.userService.Authenticate(ctx, login, password)
	if err != nil {
		return "", "", err
	}

	accessToken, err := s.jwtProvider.GenerateAccessToken(userID)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.jwtProvider.GenerateRefreshToken(userID)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthServiceImpl) RefreshAccessToken(ctx context.Context, refreshToken string) (string, string, error) {
	token, err := s.jwtProvider.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", "", errors.New("invalid refresh token")
	}

	tokenHash := s.jwtProvider.HashToken(refreshToken)
	blacklisted, err := s.blacklistRepo.IsBlacklisted(ctx, tokenHash)
	if err != nil {
		return "", "", err
	}
	if blacklisted {
		return "", "", errors.New("refresh token already used")
	}

	userID, err := s.jwtProvider.GetUserIDFromToken(token)
	if err != nil {
		return "", "", err
	}

	if _, err := s.userService.GetUserByID(ctx, userID); err != nil {
		return "", "", errors.New("user not found")
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	exp, ok := claims["exp"].(float64)
	if ok {
		expiresAt := time.Unix(int64(exp), 0)
		if err := s.blacklistRepo.Add(ctx, tokenHash, expiresAt); err != nil {
			return "", "", err
		}
	}

	newAccessToken, err := s.jwtProvider.GenerateAccessToken(userID)
	if err != nil {
		return "", "", err
	}

	return newAccessToken, refreshToken, nil
}

func (s *AuthServiceImpl) RefreshRefreshToken(ctx context.Context, refreshToken string) (string, string, error) {
	token, err := s.jwtProvider.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", "", errors.New("invalid refresh token")
	}

	tokenHash := s.jwtProvider.HashToken(refreshToken)
	blacklisted, err := s.blacklistRepo.IsBlacklisted(ctx, tokenHash)
	if err != nil {
		return "", "", err
	}
	if blacklisted {
		return "", "", errors.New("refresh token already used")
	}

	userID, err := s.jwtProvider.GetUserIDFromToken(token)
	if err != nil {
		return "", "", err
	}

	if _, err := s.userService.GetUserByID(ctx, userID); err != nil {
		return "", "", errors.New("user not found")
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	exp, ok := claims["exp"].(float64)
	if ok {
		expiresAt := time.Unix(int64(exp), 0)
		if err := s.blacklistRepo.Add(ctx, tokenHash, expiresAt); err != nil {
			return "", "", err
		}
	}

	newAccessToken, err := s.jwtProvider.GenerateAccessToken(userID)
	if err != nil {
		return "", "", err
	}

	newRefreshToken, err := s.jwtProvider.GenerateRefreshToken(userID)
	if err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshToken, nil
}

func (s *AuthServiceImpl) GetUserByAccessToken(ctx context.Context, accessToken string) (domain.User, error) {
	token, err := s.jwtProvider.ValidateAccessToken(accessToken)
	if err != nil {
		return domain.User{}, errors.New("invalid access token")
	}

	userID, err := s.jwtProvider.GetUserIDFromToken(token)
	if err != nil {
		return domain.User{}, err
	}

	return s.userService.GetUserByID(ctx, userID)

}
