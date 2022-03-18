package taskRouter

import (
	taskController "deepakr-28/simple-golang-backend/controllers"

	"github.com/gorilla/mux"
)

func TasksRouter() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/tasks", taskController.CreateTask).Methods("POST")
	router.HandleFunc("/api/tasks", taskController.GetTasks).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", taskController.GetSingleTask).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", taskController.UpdateTask).Methods("PUT")
	router.HandleFunc("/api/tasks/{id}", taskController.Delete).Methods("DELETE")
	return router
}
