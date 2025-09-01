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

func GetAlerts(ctx context.Context, lat, lng, radius float64) ([]model.Alert, error) {
	getLocation := `id, title, description, category, user_id, created_at, updated_at, ST_AsText(location) as location 
	FROM alerts
	WHERE ST_DWithin(location, ST_MakePoint($1, $2), $3)
	ORDER BY ST_Distance(location, ST_MakePoint($1, $2)::geography) ASC;
	`

	rows, err := database.DB.Query(ctx, getLocation, lng, lat, radius)
	if err != nil {
		return nil, fmt.Errorf("error querying alerts: %w", err)
	}
	defer rows.Close()

	var alerts []model.Alert

	for rows.Next() {
		var alert model.Alert
		var locationWKT string
		if err := rows.Scan(&alert.ID, &alert.Title, &alert.Description, &alert.Category, &alert.UserID, &alert.CreatedAt, &alert.UpdatedAt, &locationWKT); err != nil {
			return nil, fmt.Errorf("error scanning alert row: %w", err)
		}

		var lng, lat float64
		if _, err := fmt.Sscanf(locationWKT, "POINT(%f %f)", &lng, &lat); err != nil {
			return nil, fmt.Errorf("error parsing location WKT: %w", err)
		}
		alert.Location = model.Point{Lat: lat, Long: lng}

		alerts = append(alerts, alert)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("error iterating over alert rows: %w", rows.Err())
	}

	return alerts, nil
}
