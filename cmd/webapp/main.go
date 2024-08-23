package main

import (
	"server/api"
	"time"
)

func main() {
    time.Sleep(time.Second * 3)
	apiInstance := api.NewAPIBuilder().
		WithPort(8080).
		Build()

	apiInstance.StartServer()
}
