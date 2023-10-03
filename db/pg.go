package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DBPool *pgxpool.Pool

func InitPostgress() {
	databaseURL := "postgres://postgres:rhlDBPass@rhldb.c1szjdndqicm.ap-south-1.rds.amazonaws.com:5432/userdb"
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		panic(err)
	}
	config.MaxConns = 20
	// Initialize the connection pool and assign it to the global variable DBPool
	DBPool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		panic(err)
	}
}
