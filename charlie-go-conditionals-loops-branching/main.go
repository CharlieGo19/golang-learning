package main

import (
	"fmt"
)

func main() {
	// typical for loop
	// can also init i outside of loop and do for ; i < 5; i++
	for i := 0; i < 5; i++ {
		if i == 2 {
			// ignore 2.
			continue
		} else if i == 3 {
			println("You're one cheeky C U Next Tuesday M8!")
		} else {
			println(i, "is not three.")
		}
	}
	var value int8 = 0
	// if conditional was removed this would be infinite loop, aka, while(true) loop :( RIP
	for {
		println("WHY DOESN'T GO USE WHILE LOOP? SAD!", value)
		if value == 10 {
			break
		} else {
			value++
		}
	}
	// looping through a slice, note _ is a special write only var, it's for a variable you do not need.
	// in this case _ is the index, as I have no intention of printing the actual size, just the reaction.
	var benisSize map[int]string = map[int]string{1: "Smoll...", 4: "H'ok", 9: "Stonks!"}
	for _, react := range benisSize {
		println("You benis size is", react)
		// switch statements have implicit breaking, so no need to type break on every case
		// if you do want it to carry on you can use the fallthrough statement.
		switch react {
		case "Stonks!":
			println("Oh my, you are a pretty thang... No homo.")
		default:
			println("You leave much to be desired...")
		}
	}
	// a panic (like an exception) is an oh snap! button, they should be use sparingly, usually if you cannot handle, then you
	// would return an error and let the calling function handle it.
	// this is not printed out, but would be returned, and the calling function could if it wanted to print it.
	fmt.Errorf("Can return formated errors, like what the value, %v, of value was.", value)
	panic("Couldn't find the benis!")
	// note: for testing webapis, look at postman
}
