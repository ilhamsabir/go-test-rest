package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/user/go-test-rest/config"
	. "github.com/user/go-test-rest/models"
	. "github.com/user/go-test-rest/repository"
	utils "github.com/user/go-test-rest/utils"
	"gopkg.in/mgo.v2/bson"
)

var config = Config{}
var repoAccess = UserRepoAccess{}

// GET list of movies
func AllUserEndPoint(w http.ResponseWriter, r *http.Request) {
	users, err := repoAccess.FindAll()
	if err != nil {
		utils.SendJSONResponse(w, 1, "Error", err.Error(), http.StatusInternalServerError)
		return
	}
	utils.SendJSONResponse(w, 0, "Success", users, http.StatusOK)
}

// GET a movie by its ID
func FindUserEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := repoAccess.FindById(params["id"])
	if err != nil {
		utils.SendJSONResponse(w, 1, "Error", "Invalid user ID", http.StatusBadRequest)
		return
	}
	utils.SendJSONResponse(w, 0, "Success", user, http.StatusOK)
}

// POST a new movie
func CreateUserEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.SendJSONResponse(w, 1, "Error", "Invalid request payload", http.StatusBadRequest)
		return
	}
	user.ID = bson.NewObjectId()
	if err := repoAccess.Insert(user); err != nil {
		utils.SendJSONResponse(w, 1, "Error", err.Error(), http.StatusInternalServerError)

		return
	}
	utils.SendJSONResponse(w, 0, "Success", user, http.StatusOK)
}

// PUT update an existing movie
func UpdateUserEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.SendJSONResponse(w, 1, "Error", "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := repoAccess.Update(user); err != nil {
		utils.SendJSONResponse(w, 1, "Error", err.Error(), http.StatusInternalServerError)
		return
	}
	utils.SendJSONResponse(w, 0, "Success", map[string]string{"result": "success"}, http.StatusOK)
}

// DELETE an existing movie
func DeleteUserEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.SendJSONResponse(w, 1, "Error", "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := repoAccess.Delete(user); err != nil {
		utils.SendJSONResponse(w, 1, "Error", err.Error(), http.StatusInternalServerError)
		return
	}
	utils.SendJSONResponse(w, 0, "Success", map[string]string{"result": "success"}, http.StatusOK)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	repoAccess.Server = config.Server
	repoAccess.Database = config.Database
	repoAccess.Connect()
}
