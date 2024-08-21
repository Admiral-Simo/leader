package main

import (
	"context"
	"fmt"
	"log"
	"server/store"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbURL := "postgres://postgres:my_password@localhost:5432/sqlctest?sslmode=disable"

	ctx := context.Background()

	dbpool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}

	queries := store.New(dbpool)
	users, err := queries.ListUsers(ctx)
	if err != nil {
	log.Fatal(err)
	}
	printUsers(users)
}

func printMessages(msgs []store.UserMessage) {
	for i, msg := range msgs {
		fmt.Printf("Info: %d %s %s\nMessage: %s.\n", msg.ID, msg.Name, msg.Email, msg.Message)
		if i != len(msgs)-1 {
			fmt.Println()
		}
	}
}

func printUsers(users []store.User) {
	for i, user := range users {
		fmt.Printf("User: %+v\n", user)
		if i != len(users)-1 {
			fmt.Println()
		}
	}
}
