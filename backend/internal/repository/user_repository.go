package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"vizinhanca/internal/database"
	"vizinhanca/internal/model"

	"github.com/jackc/pgx/v5"
)

func CreateUser(ctx context.Context, user *model.User) error {
	result := "INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)"
	commandTag, err := database.DB.Exec(ctx, result, user.Username, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("erro ao inserir: %w", err)
	}
	rows := commandTag.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("no rows were inserted")
	}
	return nil
}

func GetUserAuth(ctx context.Context, username string) (*model.User, error) {
	result := "SELECT id, username, password_hash FROM users WHERE username = $1"
	var user model.User
	err := database.DB.QueryRow(ctx, result, username).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("failed to query user: %w", err)
	}
	user.Username = username
	return &user, nil
}

func GetUserPublic(ctx context.Context, username string) (*model.User, error) {
	user := &model.User{Username: username}
	result := "SELECT usernamel FROM users WHERE username = $1"
	err := database.DB.QueryRow(ctx, result, user.Username).Scan(&user.Username, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("failed to query user: %w", err)
	}
	return user, nil
}
