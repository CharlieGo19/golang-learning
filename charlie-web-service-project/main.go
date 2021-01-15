package main

// Note: a package is kind of like a class with respects to encapsulation.
// Package is just a directory that contains discrete units of code for some sort fo concept.

import (
	"fmt"

	"charlie.web.service/stonks/heroes"
)

func main() {
	var heroOne heroes.Hero = heroes.Hero{
		ID:         2,
		HeroName:   "Geralt of Rivia",
		HeroAlias:  "The Witcher",
		HeroRating: 1000000,
	}
	fmt.Println(heroOne.HeroName, "is a solid", heroOne.HeroRating)
}
