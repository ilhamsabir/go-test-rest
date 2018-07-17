// Repository will store any Database handler.
// Querying, or Creating/ Inserting into any database will stored here

package repository

import (
	"log"

	. "github.com/user/go-test-rest/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserRepoAccess struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "users"
)

// Establish a connection to database
func (m *UserRepoAccess) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of movies
func (m *UserRepoAccess) FindAll() ([]User, error) {
	var user []User
	err := db.C(COLLECTION).Find(bson.M{}).All(&user)
	return user, err
}

// Find a movie by its id
func (m *UserRepoAccess) FindById(id string) (User, error) {
	var user User
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// Insert a movie into database
func (m *UserRepoAccess) Insert(user User) error {
	err := db.C(COLLECTION).Insert(&user)
	return err
}

// Delete an existing movie
func (m *UserRepoAccess) Delete(user User) error {
	err := db.C(COLLECTION).Remove(&user)
	return err
}

// Update an existing movie
func (m *UserRepoAccess) Update(user User) error {
	err := db.C(COLLECTION).UpdateId(user.ID, &user)
	return err
}
