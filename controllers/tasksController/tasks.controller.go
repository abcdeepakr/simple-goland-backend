package taskController

import (
	model "deepakr-28/simple-golang-backend/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// POST create Task
var tasks []model.Task

func createTask(task model.Task) []model.Task {
	tasks = append(tasks, task)
	return tasks
}

// GET get Task

func getTask(taskId int) model.Task {

	var result model.Task
	for _, task := range tasks {
		if task.Id == taskId {
			result = task
		}
	}
	return result
}

// GET get Tasks

func getTasks() []model.Task {
	return tasks
}

// PUT update Task

func updateTask(taskId int, taskBody model.Task) model.Task {

	for _, task := range tasks {
		if task.Id == taskId {
			task = taskBody
		}
	}
	return taskBody
}

// DELETE delete Task
func deleteTask(taskId int) []model.Task {

	for index, task := range tasks {
		if task.Id == taskId {
			tasks = append(tasks[:index], tasks[index+1:]...)
		}
	}
	return tasks
}

// METHODS DEFINED IN THE ROUTER

func CreateTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var task model.Task

	_ = json.NewDecoder(r.Body).Decode(&task) // create a decoder, and pass the request body which then decodes the body in the structure

	// the request body has been decoded in the structure we defined in the model
	// the same can now be passed to the helper function createTask, which will create a task for us

	createTask(task)

	json.NewEncoder(w).Encode(task) // this will be returned in the api response body

}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tasks := getTasks()
	json.NewEncoder(w).Encode(tasks) // this will be returned in the api response body

}

func GetSingleTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	taskId, _ := strconv.Atoi(params["id"])
	task := getTask(taskId)

	json.NewEncoder(w).Encode(task) // this will be returned in the api response body

}

func UpdateTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	var newTask model.Task
	_ = json.NewDecoder(r.Body).Decode(&newTask)
	taskId, _ := strconv.Atoi(params["id"])

	task := updateTask(taskId, newTask)
	json.NewEncoder(w).Encode(task) // this will be returned in the api response body

}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	taskId, _ := strconv.Atoi(params["id"])

	task := deleteTask(taskId)
	json.NewEncoder(w).Encode(task)
}

type Post struct {
	UserID int32  `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func ExternalApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		log.Fatal(err)

	}

	defer resp.Body.Close()
	var posts []Post

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(body, &posts)
	for _, post := range posts {
		// fmt.Printf("\n%+v\n", posts[i])
		fmt.Println("TITLE : ", post.Title)
	}
	json.NewEncoder(w).Encode(posts)
}
func Root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		log.Fatal(err)

	}

	defer resp.Body.Close()
	var posts []Post

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(body, &posts)
	for _, post := range posts {
		// fmt.Printf("\n%+v\n", posts[i])
		fmt.Println("TITLE : ", post.Title)
	}
	json.NewEncoder(w).Encode(posts)
}
