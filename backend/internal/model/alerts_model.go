package model

import "time"

type Point struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Alert struct {
	ID          int64     `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Category    string    `json:"category" db:"category"`
	Location    Point     `json:"location" db:"location"`
	UserID      int64     `json:"user_id" db:"user_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
