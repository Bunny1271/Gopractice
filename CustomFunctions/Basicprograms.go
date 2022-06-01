package Customfunctions

import "fmt"

func Basicprograms() {
	fmt.Println("Hello")

	// To find area and circumference of circle

	var area float64
	var cir float64
	var radius float64
	const pi = 3.14

	fmt.Println("Enter the radius of circle:")
	fmt.Scan(&radius)

	area = pi * radius * radius
	fmt.Println("Area of the circle is:", area)

	cir = 2 * pi * radius
	fmt.Println("Circumferences of a circle is:", cir)

	// To find factorial of a number

	var num int
	var factorial int = 1
	fmt.Println("Enter the number to factorial:")
	fmt.Scan(&num)
	if num < 0 {
		fmt.Println("Cannot factorial a negative number, so Enter a positive number")
	} else {
		for i := 1; i <= num; i++ {
			factorial = factorial * i
		}
		fmt.Println("Factorial is ", factorial)
	}

}
