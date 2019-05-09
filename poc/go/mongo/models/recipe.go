package models

import (
	"log"

	"github.com/twinj/uuid"
)

// Ingredient is a flavor with a precentage
type Ingredient struct {
	Flavor     *Flavor `json:"flavor" bson:"flavor"`
	Percentage float64 `json:"percentage" bson:"percentage"`
}

// Recipe has a name & a creator and ingredients
type Recipe struct {
	UUID        string        `json:"uuid" bson:"uuid"`
	Name        string        `json:"name" bson:"name"`
	Creator     string        `json:"creator" bson:"creator"` // FIXME: naming: I don't like mixer
	Ingredients []*Ingredient `json:"ingredients" bson:"ingredients"`
}

func (db *DB) recipeCollection() {
	c := db.session.DB("gusta").C("recipe")
	addIndex(c, []string{"name", "creator"})
	addIndex(c, []string{"uuid"})
	db.collections[RecipeCollection] = c
}

func (db *DB) AddRecipe(r *Recipe) {
	r.UUID = uuid.NewV1().String()
	err := db.collections[RecipeCollection].Insert(r)
	if err != nil {
		log.Printf("Failed to add recipe %+v: %s", r, err)
	}
}
