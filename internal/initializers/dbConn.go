package initializers

import (
	"context"
	"log"

	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectToDB(ctx context.Context) (*pgxpool.Pool, error){
	DB_URL := os.Getenv("DB_URL")
	pool, err := pgxpool.New(ctx, DB_URL)
	if err != nil {
		log.Println("Failed to log in to the db", err.Error())
		return nil, err
	}
	log.Println("Connected to the db")
	return pool, nil 
}
