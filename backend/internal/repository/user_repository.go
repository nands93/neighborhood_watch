package repository

import (
	"context"
	"fmt"
	"vizinhanca/internal/database"
	"vizinhanca/internal/model"
)

func CreateUser(ctx context.Context, user *model.User) error {
	result := "INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)"
	_, err := database.DB.Exec(ctx, result, user.Username, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("erro ao inserir: %w", err)
	}

	return nil
}
