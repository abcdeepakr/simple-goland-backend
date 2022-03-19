package taskRouter

import (
	taskController "deepakr-28/simple-golang-backend/controllers/tasksController"
	userController "deepakr-28/simple-golang-backend/controllers/usersController"

	"github.com/gorilla/mux"
)

func TasksRouter() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/", taskController.Root).Methods("GET")
	router.HandleFunc("/api/tasks", taskController.CreateTask).Methods("POST")
	router.HandleFunc("/api/tasks", taskController.GetTasks).Methods("GET")
	router.HandleFunc("/api/getData", taskController.ExternalApi).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", taskController.GetSingleTask).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", taskController.UpdateTask).Methods("PUT")
	router.HandleFunc("/api/tasks/{id}", taskController.Delete).Methods("DELETE")
	router.HandleFunc("/api/users", userController.CreateUser).Methods("POST")
	router.HandleFunc("/api/users", userController.GetUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", userController.GetSingleUser).Methods("GET")
	router.HandleFunc("/api/users/{id}", userController.UpdateSingleUser).Methods("PUT")
	router.HandleFunc("/api/users/{id}", userController.DeleteSingleUser).Methods("DELETE")
	return router
}
