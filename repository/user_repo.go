// Repository will store any Database handler.
// Querying, or Creating/ Inserting into any database will stored here

package dao

import (
	"log"

	. "github.com/user/go-test-rest/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "users"
)

// Establish a connection to database
func (m *UserDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of movies
func (m *UserDAO) FindAll() ([]User, error) {
	var user []User
	err := db.C(COLLECTION).Find(bson.M{}).All(&user)
	return user, err
}

// Find a movie by its id
func (m *UserDAO) FindById(id string) (User, error) {
	var user User
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// Insert a movie into database
func (m *UserDAO) Insert(user User) error {
	err := db.C(COLLECTION).Insert(&user)
	return err
}

// Delete an existing movie
func (m *UserDAO) Delete(user User) error {
	err := db.C(COLLECTION).Remove(&user)
	return err
}

// Update an existing movie
func (m *UserDAO) Update(user User) error {
	err := db.C(COLLECTION).UpdateId(user.ID, &user)
	return err
}
