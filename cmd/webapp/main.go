package main

import (
	"server/api"
)

func main() {
	apiInstance := api.NewAPIBuilder().
		WithPort(4000).
		Build()

	apiInstance.StartServer()
}
