package main

import (
	router "deepakr-28/simple-golang-backend/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	newRouter := router.TasksRouter()
	log.Fatal(http.ListenAndServe(":4000", newRouter))
	fmt.Println("LISTING ON PORT 4000")
}