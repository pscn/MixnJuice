package models

import (
	"log"

	"github.com/globalsign/mgo/bson"
	"github.com/twinj/uuid"
)

type Flavor struct {
	UUID    string  `json:"uuid" bson:"uuid"`
	Name    string  `json:"name" bson:"name"`
	Vendor  *Vendor `json:"vendor" bson:"vendor"`
	Density float64 `json:"density" bson:"density"`
}

func (db *DB) flavorCollection() {
	c := db.session.DB("gusta").C("flavor")
	addIndex(c, []string{"name", "vendor.uuid"})
	addIndex(c, []string{"uuid"})
	db.collections[FlavorCollection] = c
}

func (db *DB) AddFlavor(f *Flavor) {
	f.UUID = uuid.NewV1().String()
	err := db.collections[FlavorCollection].Insert(f)
	if err != nil {
		log.Printf("Failed to add flavor %+v: %s", f, err)
	}
}

func (db *DB) GetFlavor(name string) *Flavor {
	var result Flavor
	err := db.collections[FlavorCollection].Find(bson.M{"name": name}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	return &result
}
