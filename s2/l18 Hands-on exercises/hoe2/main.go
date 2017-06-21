package main

import (
	"os"
	"text/template"
)

/*
 * # Requirements:
 * - Create a data structure to pass to a template which
 * 	- Contains information about California hotels including Name, Address, City, Zip, Region
 * 	- Region can be: Southern, Central, Northern
 * 	- Can hold an unlimited number of hotels
 */

const (
	RegionSouthern = "Southern"
	RegionCentral  = "Central"
	RegionNorthern = "Northern"
)

var tpl *template.Template

type Hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := []Hotel{
		{
			"Hotel California",
			"Evil road 666",
			"San Francisco",
			"90210",
			RegionCentral,
		},
		{
			"Californication",
			"Mulholland Drive 135",
			"Hollywood",
			"13579",
			RegionSouthern,
		},
		{
			"California Dreammin'",
			"Acustic Av. 246",
			"Los Angeles",
			"987456",
			RegionNorthern,
		},
	}

	tpl.Execute(os.Stdout, hotels)
}
