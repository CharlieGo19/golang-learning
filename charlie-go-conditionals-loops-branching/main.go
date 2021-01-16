package main

func main() {
	// typical for loop
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
}
