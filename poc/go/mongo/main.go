package main

import (
	"log"

	"github.com/pscn/MixnJuice/poc/go/mongo/models"
)

func main() {
	log.Printf("Starting up")
	connection := models.Connect("mongodb://mongodb:27017")
	defer connection.Close()
	connection.AddVendor(&models.Vendor{
		Name: "Capella",
		Abbr: "CAP",
		UUID: "CAP",
	})
	connection.AddVendor(&models.Vendor{
		Name: "The Perfumers Apprentice",
		Abbr: "TPA",
		UUID: "TPA",
	})
	v := connection.GetVendor("Capella")
	log.Printf("Vendor=%+v", v)
	v.Name = "Cappela Inc."
	log.Printf("Vendor=%+v", v)
	//connection.UpdateVendor(v)

	f := &models.Flavor{
		Name:    "Vanilla Custrad",
		Vendor:  v,
		Density: 1.012,
	}
	log.Printf("PRE Flavor=%+v", f)
	log.Printf("PRE Vendor=%+v", f.Vendor)
	connection.AddFlavor(f)
	f = connection.GetFlavor("Vanilla Custrad")
	log.Printf("POST Flavor=%+v", f)
	log.Printf("POST Vendor=%+v", f.Vendor)
}
