package service

import (
	"TicTacToe/internal/domain"
	"TicTacToe/internal/repository"
	"context"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, login, password string) (uuid.UUID, error)
	Authenticate(ctx context.Context, login, password string) (uuid.UUID, error)
	ListAllUsers(ctx context.Context) ([]domain.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (domain.User, error)
}

type UserServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) Register(ctx context.Context, login, password string) (uuid.UUID, error) {

	if len(login) < 3 {
		return uuid.Nil, errors.New("login too short")
	}

	if len(password) < 6 {
		return uuid.Nil, errors.New("password too short")
	}

	if _, err := s.repo.FindByLogin(ctx, login); err == nil {
		return uuid.Nil, errors.New("user already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.Nil, err
	}

	user := domain.User{
		ID:       uuid.New(),
		Login:    login,
		PassHash: string(hash),
	}

	if err := s.repo.Save(ctx, user); err != nil {
		return uuid.Nil, err
	}

	return user.ID, nil
}

func (s *UserServiceImpl) Authenticate(ctx context.Context, login, password string) (uuid.UUID, error) {
	user, err := s.repo.FindByLogin(ctx, login)

	if err != nil {
		return uuid.Nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(password)); err != nil {
		return uuid.Nil, errors.New("invalid credentials")
	}

	return user.ID, nil
}

func (s *UserServiceImpl) ListAllUsers(ctx context.Context) ([]domain.User, error) {
	return s.repo.ListUsers(ctx)
}

func (s *UserServiceImpl) GetUserByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	return s.repo.FindByID(ctx, id)
}
