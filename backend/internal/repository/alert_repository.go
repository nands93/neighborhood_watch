package repository

import (
	"context"
	"fmt"
	"vizinhanca/internal/database"
	"vizinhanca/internal/model"
)

func CreateAlert(ctx context.Context, alert *model.Alert) error {
	location := fmt.Sprintf("POINT(%f %f)", alert.Location.Long, alert.Location.Lat)
	result := "INSERT INTO alerts (title, description, category, location, user_id) VALUES ($1, $2, $3, ST_GeogFromText($4), $5)"
	commandTag, err := database.DB.Exec(ctx, result, alert.Title, alert.Description, alert.Category, location, alert.UserID)
	if err != nil {
		return fmt.Errorf("error inserting alert: %w", err)
	}
	if commandTag.RowsAffected() != 1 {
		return fmt.Errorf("no rows were inserted")
	}
	return nil
}
