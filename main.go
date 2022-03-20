package main

import (
	router "deepakr-28/simple-golang-backend/router"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000" // set default port if not specified
	}

	newRouter := router.TasksRouter()
	log.Fatal(http.ListenAndServe(":"+port, newRouter))
	fmt.Println("LISTING ON PORT 4000")
}
