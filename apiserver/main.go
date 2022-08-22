package main

import (
	"fmt"
	"short_url/apiserver/pkg/database"
	"short_url/apiserver/pkg/router"
)

func main() {
	// database
	database.Initialize("data.db")

	// routing
	r := router.Init()

	// start
	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the service - Error: %v", err))
	}
}