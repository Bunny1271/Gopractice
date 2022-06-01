package Customfunctions

import (
	"fmt"
	"unsafe"
)

/*
	The following are the basic types available in Go

	bool
	Numeric Types
	int8, int16, int32, int64, int
	uint8, uint16, uint32, uint64, uint
	float32, float64
	complex64, complex128
	byte  	--	byte is an alias of uint8
	rune	--  rune is an alias of int32
	string
*/

/*
	Signed integers
	int8: represents 8 bit signed integers
	size: 8 bits
	range: -128 to 127

	int16: represents 16 bit signed integers
	size: 16 bits
	range: -32768 to 32767

	int32: represents 32 bit signed integers
	size: 32 bits
	range: -2147483648 to 2147483647

	int64: represents 64 bit signed integers
	size: 64 bits
	range: -9223372036854775808 to 9223372036854775807

	int: represents 32 or 64 bit integers depending on the underlying
	platform. You should generally be using int to represent integers
	unless there is a need to use a specific sized integer.
	size: 32 bits in 32 bit systems and 64 bit in 64 bit systems.
	range: -2147483648 to 2147483647 in 32 bit systems
	and -9223372036854775808 to 9223372036854775807 in 64 bit systems
*/

/*
	Unsigned integers
	uint8: represents 8 bit unsigned integers
	size: 8 bits
	range: 0 to 255

	uint16: represents 16 bit unsigned integers
	size: 16 bits
	range: 0 to 65535

	uint32: represents 32 bit unsigned integers
	size: 32 bits
	range: 0 to 4294967295

	uint64: represents 64 bit unsigned integers
	size: 64 bits
	range: 0 to 18446744073709551615

	uint : represents 32 or 64 bit unsigned integers depending on the underlying platform.
	size : 32 bits in 32 bit systems and 64 bits in 64 bit systems.
	range : 0 to 4294967295 in 32 bit systems and 0 to 18446744073709551615 in 64 bit systems

*/

func SampleTypes() {

	//Boolean type return either true or false
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b // Both a and b must be true
	fmt.Println("c:", c)
	d := a || b // either a or b is true
	fmt.Println("d:", d)

	//Integer
	var a1 int = 89
	b1 := 95
	fmt.Println("value of a is", a1, "and b is", b1)

	var a2 int = 89
	b2 := 95
	fmt.Println("value of a is", a2, "and b is", b2)
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a2))   //type and size of a
	fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b2)) //type and size of b

	// Floating point types
	// float32: 32 bit floating point numbers
	// float64: 64 bit floating point numbers

	a3, b3 := 5.67, 8.97
	fmt.Printf("type of a %T b %T\n", a3, b3)
	sum := a3 + b3
	diff := a3 - b3
	fmt.Println("sum", sum, "diff", diff)
	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no1-no2)

	// Complex types
	// complex64: complex numbers which have float32 real and imaginary parts
	// complex128: complex numbers with float64 real and imaginary parts

	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)

	//String type

	first := "Naveen"
	last := "Ramanathan"
	name := first + " " + last
	fmt.Println("My name is", name)

	// Type Conversions

	i := 55   //int
	j := 67.8 //float64
	// sum := i + j //int + float64 not allowed
	k := i + int(j) //j is converted to int
	fmt.Println(k)

}
