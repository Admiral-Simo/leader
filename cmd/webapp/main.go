package main

import (
	"log"
	"server/api"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	apiInstance := api.NewAPIBuilder().
		WithPort(8080).
		Build()

	apiInstance.StartServer()
}
