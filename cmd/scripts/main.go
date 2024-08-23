package main

import (
	"context"
	"fmt"
	"log"
	"server/store"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbURL := "postgres://postgres:strongpassword@localhost:5433/leader?sslmode=disable"

	ctx := context.Background()

	dbpool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}

	queries := store.New(dbpool)
	histories, _ := queries.ListSearchHistory(ctx)
	fmt.Println(histories)
}
