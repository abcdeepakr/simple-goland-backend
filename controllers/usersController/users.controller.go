package userController

import (
	model "deepakr-28/simple-golang-backend/models"
	"encoding/json"
	"strconv"

	"github.com/gorilla/mux"

	"net/http"
)

var users []model.User

func createUser(userBody model.User) []model.User {

	users = append(users, userBody)
	return users
}

func getAllUsers() []model.User {

	return users

}

func getSingleUser(userId int) model.User {

	var userFound model.User

	for _, user := range users {
		if user.UserId == userId {
			userFound = user
		}
	}

	return userFound
}

func updateUserDetails(userId int, userBody model.User) []model.User {
	for index, user := range users {
		if user.UserId == userId {
			users = append(users[:index], users[index+1:]...)
			users = append(users, userBody)
		}
	}
	return users
}

func deleteUser(userId int) []model.User {
	for index, user := range users {
		if user.UserId == userId {
			users = append(users[:index], users[index+1:]...)
		}
	}
	return users
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var newUser model.User

	_ = json.NewDecoder(r.Body).Decode(&newUser)

	result := createUser(newUser)

	json.NewEncoder(w).Encode(result)

}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Allow-Control-Allow-Methods", "POST")

	// var newUser model.User

	// _ = json.NewDecoder(r.Body).Decode(&newUser)

	result := getAllUsers()

	json.NewEncoder(w).Encode(result)
}
func GetSingleUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)

	userId, _ := strconv.Atoi(params["id"])

	result := getSingleUser(userId)

	json.NewEncoder(w).Encode(result)

}
func UpdateSingleUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	userId, _ := strconv.Atoi(params["id"])

	var updatedUser model.User
	_ = json.NewDecoder(r.Body).Decode(&updatedUser)

	result := updateUserDetails(userId, updatedUser)
	json.NewEncoder(w).Encode(result)

}
func DeleteSingleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	userId, _ := strconv.Atoi(params["id"])

	// var updatedUser model.User
	// _ = json.NewDecoder(r.Body).Decode(&updatedUser)

	result := deleteUser(userId)
	json.NewEncoder(w).Encode(result)
}
