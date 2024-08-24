package main

import (
	"server/api"
)

func main() {
	apiInstance := api.NewAPIBuilder().
		WithPort(8080).
		Build()

	apiInstance.StartServer()
}
