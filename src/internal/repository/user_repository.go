package repository

import (
	"TicTacToe/internal/domain"
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed queries/create_users_table.sql
var createUsersTableQuery string

//go:embed queries/save_user.sql
var saveUserQuery string

//go:embed queries/find_user_by_login.sql
var findUserByLogin string

//go:embed queries/list_users.sql
var listUsersQuery string

//go:embed queries/find_user_by_id.sql
var findUserByIDQuery string

type UserRepositoryImpl struct {
	pool *pgxpool.Pool
}

func NewUserRepositoryImpl(pool *pgxpool.Pool) (*UserRepositoryImpl, error) {
	ctx := context.Background()

	if _, err := pool.Exec(ctx, createUsersTableQuery); err != nil {
		return nil, fmt.Errorf("failed to create users table: %w", err)
	}

	if _, err := pool.Exec(ctx, createBlacklistTableQuery); err != nil {
		return nil, fmt.Errorf("failed to create blacklist table: %w", err)
	}

	return &UserRepositoryImpl{pool: pool}, nil
}

func (r *UserRepositoryImpl) Save(ctx context.Context, user domain.User) error {
	_, err := r.pool.Exec(ctx, saveUserQuery, user.ID, user.Login, user.PassHash)
	return err
}

func (r *UserRepositoryImpl) FindByLogin(ctx context.Context, login string) (domain.User, error) {
	var user domain.User
	err := r.pool.QueryRow(ctx, findUserByLogin, login).Scan(&user.ID, &user.Login, &user.PassHash)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.User{}, errors.New("user not found")
	}
	return user, err
}

func (r *UserRepositoryImpl) ListUsers(ctx context.Context) ([]domain.User, error) {
	rows, err := r.pool.Query(ctx, listUsersQuery)
	if err != nil {
		return nil, fmt.Errorf("query users:%w", err)
	}
	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var u domain.User

		if err := rows.Scan(&u.ID, &u.Login, &u.PassHash); err != nil {
			return nil, fmt.Errorf("scan user:%w", err)
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

func (r *UserRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	var user domain.User
	err := r.pool.QueryRow(ctx, findUserByIDQuery, id).Scan(&user.ID, &user.Login, &user.PassHash)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.User{}, errors.New("user not found")
	}
	return user, err
}
