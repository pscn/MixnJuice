package models

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/twinj/uuid"
)

// Vendor with name and abbreviation
type Vendor struct {
	UUID string `json:"uuid" bson:"uuid"`
	Name string `json:"name" bson:"name"`
	Abbr string `json:"abbr" bson:"abbr"`
}

func (db *DB) vendorCollection() {
	c := db.session.DB("gusta").C("vendor")
	addIndex(c, []string{"name"})
	addIndex(c, []string{"abbr"})
	addIndex(c, []string{"uuid"})

	db.collections[VendorCollection] = c
}

func (db *DB) AddVendor(v *Vendor) {
	v.UUID = uuid.NewV1().String()
	err := db.collections[VendorCollection].Insert(v)
	if err != nil {
		log.Printf("Failed to add vendor %+v: %s", v, err)
	}
}

func (db *DB) GetVendor(name string) *Vendor {
	var result Vendor
	err := db.collections[VendorCollection].Find(bson.M{"name": name}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	return &result
}

func (db *DB) UpdateVendor(v *Vendor) {
	// strange, I'd expect this to work.  what am I missing?
	change := mgo.Change{
		// only allow to change name and abbr
		Update:    bson.M{"name": v.Name, "abbr": v.Abbr},
		ReturnNew: true,
	}
	_, err := db.collections[VendorCollection].Find(
		bson.M{"uuid": v.UUID},
	).Apply(change, v)
	if err != nil {
		log.Printf("Failed to update vendor %+v: %s", v, err)
	}
}
