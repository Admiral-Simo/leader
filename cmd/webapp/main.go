package main

import (
	"server/api"
	"server/engines"
)

func main() {
	googleEngine := engines.NewGoogleSearchEngine()
	apiInstance := api.NewAPIBuilder().
		WithPort(4000).
		WithSearchEngine(googleEngine).
		Build()

	apiInstance.StartServer()
}
