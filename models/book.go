package models

import "gopkg.in/mgo.v2/bson"

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Book struct {
	ID     bson.ObjectId `bson:"_id" json:"id"`
	Name   string        `bson:"name" json:"name"`
	Author string        `bson:"author" json:"author"`
	Stock  string        `bson:"stock" json:"stock"`
}
