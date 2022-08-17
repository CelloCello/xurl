package main

import (
	"short_url/apiserver/pkg/database"
	"short_url/apiserver/pkg/router"
)

func main() {
	// database
	database.Initialize("data.db")

	// routing
	r := router.Init()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
