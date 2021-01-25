package main

import (
	"fmt"
	"reflect"
)

type hero struct {
	heroID    int
	heroName  string
	heroTitle string
}

type antiHero struct {
	antiHeroID       int
	antiHeroName     string
	antiHeroTitle    string
	antiHeroHomeCity string
}

func main() {
	var newHero hero = hero{0, "Geralt", "of Rivia"}
	var newAntiHero antiHero = antiHero{0, "Daredevil", "Devil of Hell's Kitchen", "New York City"}

	fmt.Printf("Our hero is %s, %s.\n", newHero.heroName, newHero.heroTitle)
	fmt.Printf("Our antihero is %s, %s, from %s.\n", newAntiHero.antiHeroName, newAntiHero.antiHeroTitle, newAntiHero.antiHeroHomeCity)

	// Imagine you have lots of code or a dynamically created type (i.e. varName := x)
	// We can use reflect to identify suc things.
	// Note: When using reflection in this way, you can get some weird results!

	fmt.Printf("The type that defines %s, %s is %v\n", newHero.heroName, newHero.heroTitle, reflect.TypeOf(newHero))
	// We can also get the values of the struct if we don't know how it is structured.
	fmt.Printf("The value is %v\n", reflect.ValueOf(newHero))
	// Additionally we can determine the kind the type is.
	fmt.Printf("The kind of type that defines %s, %s is a %v\n", newHero.heroName, newHero.heroTitle, reflect.TypeOf(newHero).Kind())

	/* Now to demonstrate how using reflection could be useful:
	   Imagine we have two similar structs that can be put into db,
	   rather than creating two functions to add distinct entities we
	   could use reflection to identify its a struct, then do a switch.
	*/
	addSomeHero(newHero)
	addSomeHero(newAntiHero)

	// When working with reflection you need to be explicit, if not, it may panic.
	// For example, if you have a slice of heroes, the name of the type will be blank.
	// Reflection is to be used mainly with interfaces (if I've understoof it correctly).

	sliceType := reflect.TypeOf([]hero{})
	// type, length, capacity
	newHeroList := reflect.MakeSlice(sliceType, 0, 0)
	newHeroList = reflect.Append(newHeroList, reflect.ValueOf(hero{0, "Superman", "Man of Steel"}))
	newHeroList = reflect.Append(newHeroList, reflect.ValueOf(hero{1, "Geralt", "of Rivia"}))

	fmt.Printf("List of heroes: %v\n", newHeroList)

}

func addSomeHero(obj interface{}) bool {
	// Have we passed in a struct?
	if reflect.TypeOf(obj).Kind() == reflect.Struct {
		// If oui, then lets get the value
		value := reflect.ValueOf(obj)
		// Let's check the name, as we should know this information
		switch reflect.TypeOf(obj).Name() {
		case "hero":
			// Do something with the value
			var sqlString string = "Some SQL prepare statement for a HERO."
			fmt.Printf("SQL: %s\n", sqlString)
			fmt.Printf("Added %v\n", value.Field(1))
		case "antiHero":
			// Do something with the value
			var sqlString string = "Some SQL prepare statement for an ANTIHERO."
			fmt.Printf("SQL: %s\n", sqlString)
			fmt.Printf("Added %v\n", value.Field(1))
		}
		// We should probably do a default that says get tuckered.
		// But we have the if for that at the moment.
		return true
	} else {
		return false
	}
}
