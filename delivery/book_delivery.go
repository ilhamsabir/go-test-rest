package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/user/go-test-rest/config"
	"gopkg.in/mgo.v2/bson"

	. "github.com/user/go-test-rest/models"
	. "github.com/user/go-test-rest/repository"
	utils "github.com/user/go-test-rest/utils"
)

var Bconfig = Config{}
var BrepoAccess = BookRepoAccess{}

// GET list of Book-Data
func GetAllBook(w http.ResponseWriter, r *http.Request) {
	book, err := BrepoAccess.FindAll()
	if err != nil {
		utils.SendJSONResponse(w, 1, "Error", err.Error(), http.StatusInternalServerError)
		return
	}
	utils.SendJSONResponse(w, 0, "Success", book, http.StatusOK)
}

// GET a Book-Data by its ID
func FindBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book, err := BrepoAccess.FindById(params["id"])
	if err != nil {
		utils.SendJSONResponse(w, 1, "Error", "Invalid book ID", http.StatusBadRequest)
		return
	}
	utils.SendJSONResponse(w, 0, "Success", book, http.StatusOK)
}

// POST a new Book-Data
func CreateBookData(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		utils.SendJSONResponse(w, 1, "Error", "Invalid Payload", http.StatusBadRequest)
		return
	}
	book.ID = bson.NewObjectId()
	if err := BrepoAccess.Insert(book); err != nil {
		utils.SendJSONResponse(w, 1, "Error", err.Error(), http.StatusInternalServerError)
		return
	}
	utils.SendJSONResponse(w, 0, "Success", book, http.StatusOK)
}

// PUT update an existing Book-Data
func UpdateBookData(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		utils.SendJSONResponse(w, 1, "Error", "Invalid Payload", http.StatusBadRequest)
		return
	}
	if err := BrepoAccess.Update(book); err != nil {
		utils.SendJSONResponse(w, 1, "Error", err.Error(), http.StatusInternalServerError)
		return
	}
	utils.SendJSONResponse(w, 0, "Success", map[string]string{"result": "success"}, http.StatusOK)
}

// DELETE an existing Book-Data
func DeleteBookData(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		utils.SendJSONResponse(w, 1, "Error", "Invalid Payload", http.StatusBadRequest)
		return
	}
	if err := BrepoAccess.Delete(book); err != nil {
		utils.SendJSONResponse(w, 1, "Error", err.Error(), http.StatusInternalServerError)
		return
	}
	utils.SendJSONResponse(w, 0, "Success", map[string]string{"result": "success"}, http.StatusOK)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	Bconfig.Read()

	BrepoAccess.Server = Bconfig.Server
	BrepoAccess.Database = Bconfig.Database
	BrepoAccess.Connect()
}
