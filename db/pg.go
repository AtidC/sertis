package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	config "github.com/spf13/viper"
)

var pool *pgxpool.Pool

func InitPostgresPool() *pgxpool.Pool {

	_pool, err := pgxpool.Connect(context.Background(), config.GetString("db.postgres.connection.string"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	pool = _pool

	return pool
}

func GetPostgresPool() *pgxpool.Pool {
	if pool != nil {
		return pool
	}

	return InitPostgresPool()
}
