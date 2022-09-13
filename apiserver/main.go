package main

import (
	"fmt"
	"xurl/apiserver/pkg/database"
	"xurl/apiserver/pkg/net"
)

func main() {
	// database
	database.Initialize("data.db")

	// network
	g := net.Init("./")

	// start
	err := g.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the service - Error: %v", err))
	}
}
