package psql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/Yhlas9/book_store/config"
	"log"
)

func Init(config config.AppConfig) *pgxpool.Pool {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.Psql.Username, config.Psql.Password, config.Psql.Host, config.Psql.Port, config.Psql.Database)

	dbPool, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}

	err = dbPool.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	return dbPool
}
