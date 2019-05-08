package main

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Vendor struct {
	UUID string `json:"uuid" bson:"uuid"`
	Name string `json:"name" bson:"name"`
	Abbr string `json:"abbr" bson:"abbr"`
}
type Flavor struct {
	UUID    string  `json:"uuid" bson:"uuid"`
	Name    string  `json:"name" bson:"name"`
	Vendor  Vendor  `json:"vendor" bson:"vendor"`
	Density float64 `json:"density" bson:"density"`
}

func getVendor(c *mgo.Collection, name string) *Vendor {
	var result Vendor
	err := c.Find(bson.M{"name": name}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	return &result
}

func main() {
	log.Printf("Starting up")
	session, err := mgo.Dial("mongodb://mongodb:27017")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	vc := session.DB("gusta").C("vendor")
	idx := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err = vc.EnsureIndex(idx)
	if err != nil {
		log.Fatal(err)
	}
	fc := session.DB("gusta").C("flavor")
	idx = mgo.Index{
		Key:        []string{"name", "vendor.uuid"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err = fc.EnsureIndex(idx)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	v := []Vendor{
		{
			Name: "Capella",
			Abbr: "CAP",
			UUID: "CAP",
		},
		{
			Name: "The Perfumers Apprentice",
			Abbr: "TPA",
			UUID: "TPA",
		},
	}
	vc.Insert(v[0])
	vc.Insert(v[1])
	log.Printf("Vendor=%+v", getVendor(vc, "Capella"))

	f := &Flavor{
		Name:    "Vanilla Custrad",
		Vendor:  v[0],
		Density: 1.012,
	}
	fc.Insert(f)
}
