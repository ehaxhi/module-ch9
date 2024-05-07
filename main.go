package main

import (
	"flag"
	"fmt"
	"github.com/ehaxhi/module-ch9/data"
	"github.com/ehaxhi/module-ch9/models"
	"github.com/ehaxhi/module-ch9/printer"
)

func main() {

	fmt.Println("Welcome")
	beachReady := flag.Bool("beach", false, "Display only beach ready")
	skiReady := flag.Bool("ski", false, "Display only ski ready")
	month := flag.Int("month", 0, "Lookup for destination month [1,12]")
	name := flag.String("name", "", "Lookup for destination by name")
	flag.Parse()
	cities, err := models.NewCities(data.NewReader())
	cq, err := models.NewQuery(*beachReady, *skiReady, *month, *name)
	if err != nil {
		fmt.Println("Fatal error, ", err)
		return
	}
	p := printer.New()
	defer p.Cleanup()
	p.CityHeader()

	cs := cities.Filter(cq)

	for _, c := range cs {
		p.CityDetails(c, cq)
	}
}
