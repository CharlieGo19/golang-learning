package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	// Simple addition functions.
	fmt.Printf("Result for 64-bit: %.12f (0.123456789123+8, at 12 dp.)\n", add64(0.123456789123, 8))
	fmt.Printf("Result for 32-bit: %.12f (0.123456789123+8, at 12 dp.)\n", add32(0.123456789123, 8))

	// Division, one with error one without (multiple return demo).
	divisionOne, divisionOneRemainder, err := divide(1.2345, 0)
	if err != nil {
		fmt.Printf("Damn, you fell for the trap card: %s\n", err)
	} else {
		fmt.Printf("Lol, if this is hit then hot damnnnnn, the answer is: %f, with a remainder of: %f.\n", divisionOne, divisionOneRemainder)
	}
	// Note: do not ignore errors!!! I'm doing it here for demonstration purposes only.
	divisionTwo, divisionTwoRemainder, _ := divide(1.2345, 6.7891)
	fmt.Printf("Answer for 1.2345 divided by 6.7891 is %f with a remainder of %f.\n", divisionTwo, divisionTwoRemainder)

	// The below is essentially as slice, as demonstrated by instance two, we get the same result!
	total1 := sum("Geralt", 1, 2, 3, 4, 5)
	fmt.Printf("The sum of your values is: %v\n", total1)
	// Below is the same as the above...
	var numbers []int = []int{1, 2, 3, 4, 5}
	total2 := sum("Clarke", numbers...)
	fmt.Printf("The sum of your values is: %v\n", total2)
}

// Two simple functions, just demonstrating how to define a function.
// I also want to demonstrate to you the potential for rounding errors,
// When not using the correct float, note the same input but different output.
func add64(parameterOne, parameterTwo float64) float64 {
	return parameterOne + parameterTwo
}

func add32(parameterOne, parameterTwo float32) float32 {
	return parameterOne + parameterTwo
}

// Demonstrating multiple return values.
func divide(parameterOne, parameterTwo float64) (float64, float64, error) {
	if parameterTwo == 0 {
		return math.NaN(), math.NaN(), errors.New("you cannot divide by zero, go back to school!")
	}
	return parameterOne / parameterTwo, math.Mod(parameterOne, parameterTwo), nil
}

// Demonstrating variadic parameters, note can only be last arguement.
// values ...int is basically creating a slice of ints, but in reverse.
func sum(name string, values ...int) int {
	fmt.Printf("Hello, %s, I will add those values up for you...\n", name)
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

// To demonstrate control flow.
// Name your return variable, then have a blank return, this allows the entire function to be evaluated.
func divideControl(parameterOne, parameterTwo float64) (result float64, remainder float64, err error) {
	if parameterTwo == 0 {
		err = errors.New("you cannot divide by zero, go back to school!")
	}
	result = parameterOne / parameterTwo
	remainder = math.Mod(parameterOne, parameterTwo)
	return
}
