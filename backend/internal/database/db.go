package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func ConnectToDB() error {
	fmt.Println("=== Starting database connection ===")

	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, reading from environment variables")
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	var err error
	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		fmt.Printf("Connection attempt %d/%d...\n", i+1, maxRetries)

		DB, err = pgxpool.New(context.Background(), dsn)
		if err == nil {
			if err := DB.Ping(context.Background()); err == nil {
				fmt.Println("Successfully connected to the database!")
				return nil
			} else {
				fmt.Printf("Ping failed: %v\n", err)
				if DB != nil {
					DB.Close()
				}
			}
		} else {
			fmt.Printf("Pool creation failed: %v\n", err)
		}

		if i < maxRetries-1 {
			fmt.Println("Waiting 3 seconds before retry...")
			time.Sleep(3 * time.Second)
		}
	}

	return fmt.Errorf("failed to connect to database after %d attempts: %v", maxRetries, err)
}
