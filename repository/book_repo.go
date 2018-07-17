// Repository will store any Database handler.
// Querying, or Creating/ Inserting into any database will stored here

package repository

import (
	"log"

	. "github.com/user/go-test-rest/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BookRepoAccess struct {
	Server   string
	Database string
}

var dbBook *mgo.Database

const (
	BookCollection = "book"
)

// Establish a connection to database
func (m *BookRepoAccess) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	dbBook = session.DB(m.Database)
}

// Find list of movies
func (m *BookRepoAccess) FindAll() ([]Book, error) {
	var book []Book
	err := dbBook.C(BookCollection).Find(bson.M{}).All(&book)
	return book, err
}

// Find a movie by its id
func (m *BookRepoAccess) FindById(id string) (Book, error) {
	var book Book
	err := dbBook.C(BookCollection).FindId(bson.ObjectIdHex(id)).One(&book)
	return book, err
}

// Insert a movie into database
func (m *BookRepoAccess) Insert(book Book) error {
	err := dbBook.C(BookCollection).Insert(&book)
	return err
}

// Delete an existing movie
func (m *BookRepoAccess) Delete(book Book) error {
	err := dbBook.C(BookCollection).Remove(&book)
	return err
}

// Update an existing movie
func (m *BookRepoAccess) Update(book Book) error {
	err := dbBook.C(BookCollection).UpdateId(book.ID, &book)
	return err
}
