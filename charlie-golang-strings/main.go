package main

import (
	"fmt"
	"strings"
	"regexp"
)

func main() {
	// Strings are just a slice of bytes
	// When you modify a string in go, it makes a copy of it, then points to the copy.
	// Therefore, Strings are read only.

	hexString := "\x41\x44\x41\x4D\x20\x49\x53\x20\x41\x57\x45\x53\x4F\x4D\x45"
	fmt.Println(hexString)

	// You can print the string as a quoted byte sequence.
	// Which helps debug non-printable characters!

	for i := 0; i < len(hexString); i++ {
		fmt.Printf("%q\n", hexString[i])
	}

	// For example:

	crazyString := "x\bd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	
	fmt.Println("Insane output:", crazyString)

	for i := 0; i < len(crazyString); i++ {
		fmt.Printf("%q\n", crazyString[i])
	}

	// If you take an index of a string, it returns the byte value.
	// To get the character, you select a range.
	
	fmt.Println("An index of a string (byte value):", hexString[6])
	fmt.Println("The range of characters of a string:", hexString[5:6])

	// Notes, when optimising, look at strings.Compare as opposed to ==.
	// Strings.Compare could be quicker sometimes.

	// Below is an example of string splitting.

	var splitMe string = "I am a line. Not of snow."
	splitHim := strings.Split(splitMe, ". ")
	for i := 0; i < len(splitHim); i++ {
		fmt.Println(splitHim[i])
	}

	// Find and replace example.
	// The last arg is the number of times you want it replacing.
	// -1 means replace all.

	var revisedString string = strings.Replace(splitMe, "Not of snow.", "Ok, I'm a line of cocaine, are you happy?!", -1)
	fmt.Println(revisedString)

	// When doing a regex search, you compile the expression when initialising.

	var regexStringToSearch string = "Charlie is super cool."
	r, _ := regexp.Compile(`C(\w[a-z]+)e\b`)
	
	fmt.Println("Did we find Charlie?")
	fmt.Println(r.MatchString(regexStringToSearch))
}