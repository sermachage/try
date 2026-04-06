// entry of program
package main

//importing packages
import (
	"fmt"
	"math"
)

// main function
func main() {
	//taking the two lengths of the triangle as input
	var a, b float64
	fmt.Print("Enter the lengths of the two sides of the triangle: ")
	fmt.Scan(&a, &b)

	//calculating the length of the hypotenuse using Pythagorean theorem
	hypotenuse := math.Sqrt(a*a + b*b)

	//printing the result
	fmt.Printf("hypotenuse is: %.2f\n", hypotenuse)
}
