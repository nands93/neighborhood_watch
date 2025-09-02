package repository

import (
	"context"
	"fmt"
	"vizinhanca/internal/database"
	"vizinhanca/internal/model"
)

func CreateAlert(ctx context.Context, alert *model.Alert) error {
	location := fmt.Sprintf("POINT(%f %f)", alert.Location.Lng, alert.Location.Lat)
	result := `
        INSERT INTO alerts (title, description, category, location, user_id) 
        VALUES ($1, $2, $3, ST_GeogFromText($4), $5)
        RETURNING id, created_at, updated_at, ST_AsText(location)`

	var locationWKT string
	err := database.DB.QueryRow(ctx, result,
		alert.Title,
		alert.Description,
		alert.Category,
		location,
		alert.UserID,
	).Scan(&alert.ID, &alert.CreatedAt, &alert.UpdatedAt, &locationWKT)
	if err != nil {
		return fmt.Errorf("error inserting alert and returning id/timestamps: %w", err)
	}

	if _, err := fmt.Sscanf(locationWKT, "POINT(%f %f)", &alert.Location.Lng, &alert.Location.Lat); err != nil {
		return fmt.Errorf("error parsing returned location: %w", err)
	}
	return nil
}

func GetAlerts(ctx context.Context, lat, lng, radius float64) ([]model.Alert, error) {
	getLocation := `SELECT id, title, description, category, user_id, created_at, updated_at, ST_AsText(location) as location 
	FROM alerts
	WHERE ST_DWithin(location, ST_MakePoint($1, $2)::geography, $3)
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

		if _, err := fmt.Sscanf(locationWKT, "POINT(%f %f)", &alert.Location.Lng, &alert.Location.Lat); err != nil {
			return nil, fmt.Errorf("error parsing location WKT: %w", err)
		}
		alerts = append(alerts, alert)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating alert rows: %w", err)
	}

	return alerts, nil
}
