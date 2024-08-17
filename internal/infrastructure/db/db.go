package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"chater/internal/config"
)

var (
	pool *pgxpool.Pool
	cfg  *config.Config
)

func ConnectDB() error {
	cfg = config.LoadConfig()
	databaseURL := cfg.Database.Host
	if databaseURL == "" {
		return fmt.Errorf("DATABASE_URL not set")
	}

	var err error
	pool, err = pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		return err
	}

	return nil
}

func CloseDB() {
	if pool != nil {
		pool.Close()
	}
}

func CreateUser(username, hashedPassword string) error {
	_, err := pool.Exec(context.Background(), "INSERT INTO users (username, password) VALUES ($1, $2)", username, hashedPassword)
	return err
}

func GetUserPassword(username string) (string, error) {
	var password string
	err := pool.QueryRow(context.Background(), "SELECT password FROM users WHERE username=$1", username).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

func UserExists(username string) bool {
	var exists bool
	err := pool.QueryRow(context.Background(), "SELECT exists (SELECT 1 FROM users WHERE username=$1)", username).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}
