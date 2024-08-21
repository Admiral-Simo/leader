package main

import (
	"server/api"
)

func main() {
	apiInstance := api.NewAPIBuilder().
		WithPort(80).
		Build()

	apiInstance.StartServer()
}
