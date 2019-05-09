package models

import (
	"log"

	"github.com/globalsign/mgo"
)

// CollectionType providing access to the collection map
type CollectionType int

const (
	VendorCollection CollectionType = iota
	FlavorCollection CollectionType = iota
	RecipeCollection CollectionType = iota
)

// DB wrapper for a MongoDB session
type DB struct {
	session     *mgo.Session
	collections map[CollectionType]*mgo.Collection
}

// Connect to the MongoDB at url
func Connect(url string) *DB {
	session, err := mgo.Dial("mongodb://mongodb:27017")
	if err != nil {
		log.Fatal(err)
	}
	db := &DB{
		session:     session,
		collections: make(map[CollectionType]*mgo.Collection, 3),
	}
	db.vendorCollection()
	db.flavorCollection()
	db.recipeCollection()

	return db
}

func addIndex(c *mgo.Collection, names []string) {
	idx := mgo.Index{
		Key:        names,
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := c.EnsureIndex(idx)
	if err != nil {
		log.Fatal(err)
	}
}

// Close the connection
func (db *DB) Close() {
	db.session.Close()
}
