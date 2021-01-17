package main

import (
	"fmt"
)

// just like imports, you can set const blocks too.

const (
	one   int    = 1
	two   int    = 2
	three string = "Three"
	four  int    = iota
	five  int    = 2 << iota
	// the value 2 [0010] is bit shifted (because iota here is now 4) 4 times [00100000] -> therefore 2^5 = 32
	// note iota increments by 1 in a constant block and starts at 0 unless defined otherwise. i.e. iota + 6 ...
	// ... or the example above.
)

func main() {
	var value int = 420
	fmt.Println(value)
	var f float32 = 3.14
	fmt.Println(f)
	// float you have to define if you need 32 or 64 bit.
	firstName := "Charlie"
	fmt.Println(firstName)
	// pointer - init a string, will reserve a string at memory location x.
	var alias *string = new(string)
	// dereference the pointer to assign a value to the location,
	// and not to try and assign the value to the address value.
	*alias = "Amaise"
	fmt.Println(alias)
	// create a string variable
	var hero string = "Clarke Kent"
	// print address of hero.
	fmt.Println(hero, "is located at:", &hero)
	// set constant - they must be declared when they're initialised.
	// As they're evaluated at compiletime, therefore you can not assign...
	// ...via a function as they're evaluated at runtime.
	const pi float32 = 3.1415
	fmt.Println(pi)
	// print from const block
	fmt.Println(one, two, three, four, five)
	// init a fixed array (arrays are static in size)
	var anArray [3]string = [3]string{"Big", "Hairy", "Witcher"}
	fmt.Println("Element 2 is:", anArray[1])
	fmt.Println("The full array reads:", anArray)
	// init a slice, a dynamic array ~ you can build a slice from an array
	// note, this method actually uses a pointer, so a chance in the array ...
	// ... will change the values in the slice and vice versa.
	// NOTE! when appending to slice, it creates a copy and no longer points to original.
	// : means from beggining of array to end & is $startIndex:$endIndex
	var aSlice []string = anArray[:]
	aSlice = append(aSlice, "Appended Nuts")
	fmt.Println("This is a slice:", aSlice)
	// maps and keys - map[key]val
	var aMap map[string]int = map[string]int{"PenisSize": 6}
	fmt.Println("Benis size in millimeters:", aMap["PenisSize"])
	delete(aMap, "PenisSize")
	fmt.Println("Did I remove my shame? ", aMap)
	// Structs have two step initialisation process.
	type witcher struct {
		ID          int
		firstName   string
		hottyRating int8
	}
	var geralt witcher = witcher{
		ID:          1,
		firstName:   "Geralt of Rivia",
		hottyRating: 10,
	}
	fmt.Printf("ID: %d, Name: %s, HottyRating: %d\n", geralt.ID, geralt.firstName, geralt.hottyRating)
}
