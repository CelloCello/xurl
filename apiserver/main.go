package main

import (
	"fmt"
	"xurl/apiserver/pkg/database"
	"xurl/apiserver/pkg/net"
)

func main() {
	// database
	database.Initialize("data/data.db")

	// network
	g := net.Init("./")

	// start
	err := g.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the service - Error: %v", err))
	}

	// srv := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: g,
	// }

	// go func() {
	// 	// service connections
	// 	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 		log.Fatalf("listen: %s\n", err)
	// 	}
	// }()

	// // Wait for interrupt signal to gracefully shutdown the server with
	// // a timeout of 5 seconds.
	// quit := make(chan os.Signal)
	// // kill (no param) default send syscanll.SIGTERM
	// // kill -2 is syscall.SIGINT
	// // kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// <-quit
	// log.Println("Shutdown Server ...")

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// if err := srv.Shutdown(ctx); err != nil {
	// 	log.Fatal("Server Shutdown:", err)
	// }
	// // catching ctx.Done(). timeout of 5 seconds.
	// select {
	// case <-ctx.Done():
	// 	log.Println("timeout of 5 seconds.")
	// }
	// log.Println("Server exiting")

}
