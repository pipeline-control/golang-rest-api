package routes

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/pipeline-control/golang-rest-api/models"
)

var users []models.User

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")

	if r.Body == nil {
		json.NewEncoder(rw).Encode("Please send some data")
		return
	}

	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	if user.IsEmpty() {
		json.NewEncoder(rw).Encode("No data inside JSON")
		return
	}

	rand.Seed(time.Now().UnixNano())
	user.Id = strconv.Itoa(rand.Intn(100))
	users = append(users, user)
	json.NewEncoder(rw).Encode(users)
}

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")
	json.NewEncoder(rw).Encode(users)
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")

	params := r.URL.Query()
	for _, user := range users {
		if user.Id == params["id"][0] {
			json.NewEncoder(rw).Encode(user)
			return
		}
	}
	json.NewEncoder(rw).Encode("No User found with given id")
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")

	if r.Body == nil {
		json.NewEncoder(rw).Encode("Please send some data")
		return
	}

	var newUser models.User
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	if newUser.IsEmpty() {
		json.NewEncoder(rw).Encode("No data inside JSON")
		return
	}

	params := r.URL.Query()
	for _, user := range users {
		if user.Id == params["id"][0] {
			user = newUser
			json.NewEncoder(rw).Encode(user)
			return
		}
	}

	json.NewEncoder(rw).Encode("No User found with given id")
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-type", "application/json")

	params := r.URL.Query()
	for index, user := range users {
		if user.Id == params["id"][0] {
			users = append(users[:index], users[index+1:]...)
			json.NewEncoder(rw).Encode("User with given id deleted successfully")
			return
		}
	}

	json.NewEncoder(rw).Encode("No User found with given id")
}
