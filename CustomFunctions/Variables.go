package Customfunctions

import (
	"fmt"
	"math"
)

func SampleVariables() {

	// Declaring a single variable
	// var name type (syntax)

	var age int // variable declaration with datatype
	fmt.Println("My age is:", age)
	age = 29 //assignment
	fmt.Println("My age is:", age)
	age = 30 //assignment
	fmt.Println("My age is:", age)

	// Declaring a variable with an initial value
	// var name type = initialvalue (syntax)

	// variable declaration with initial value
	var age1 int = 25
	fmt.Println("My age is:", age1)

	// Type inference
	// var name = initialvalue (syntax)

	var age2 = 23 // type will be inferred
	fmt.Println("My age is:", age2)

	// Multiple variable declaration
	// var name1,name2 type = initialvalue1, initialvalue2 (syntax)

	var width, height int = 100, 50 //declaring multiple variables
	var width1, height1 = 10, 5     //int is dropped
	fmt.Println("Width is:", width, "Height is:", height)
	fmt.Println("Width1 is:", width1, "Height1 is:", height1)

	/*

		if the initial value is not specified for width2 and height2,
		they will have 0 assigned as their initial value.

	*/

	var width2, height2 int
	fmt.Println("width2 is", width2, "height2 is", height2)
	width2 = 100
	height2 = 50
	fmt.Println("new width2 is", width2, "new height2 is", height2)

	/* Declaring variables belonging to different types ina single statement
	var(
		name1 = initialvalue1
		name2 = initialvalue2
	) (syntax)
	*/

	var (
		name    = "naveen"
		age3    = 18
		height3 int
	)
	fmt.Println("my name is", name, ", age is", age3, "and height is", height3)

	// * Short hand declaration
	// name := initialvalue (syntax)

	count := 10
	fmt.Println("Count =", count)

	// Declare multiple variables in a single line using short hand syntax.

	name1, age4 := "john", 20 //Short hand declaration
	fmt.Println("My name is", name1, "age is", age4)

	/*
		Short hand declaration requires initial values for all variables
		on the left-hand side of the assignment.
		The following program will print an error assignment mismatch: 2 variables but 1 values.
		This is because age has not been assigned a value.
	*/

	// name2, age5 := "naveen" //error
	// fmt.Println("my name is", name2, "age is", age5)

	/*
		Short hand syntax can only be used when at least one of the variables
		on the left side of := is newly declared.
		Consider the following program,
	*/

	a, b := 20, 30 // declare variables a and b
	fmt.Println("a is", a, "b is", b)
	b, c := 40, 50 // b is already declared but c is new
	fmt.Println("b is", b, "c is", c)
	b, c = 80, 90 // assign new values to already declared variables b and c
	fmt.Println("changed b is", b, "c is", c)
	/* In the above program, in line no.94,
	   b has already been declared but c is newly
	   declared and hence it works and outputs
	*/

	// Where if we run the programm below
	/*
			a, b := 20, 30 //a and b declared
		    fmt.Println("a is", a, "b is", b)
		    a, b := 40, 50 //error, no new variables

		// It will print error
		This is because both the variables a and b
		have already been declared and there are no new
		variables in the left side of := in line no.103
	*/

	/*Variables can also be assigned values which are
	computed during run time. Consider the following
	program,
	*/

	a1, b1 := 145.8, 543.8
	c1 := math.Min(a1, b1)
	fmt.Println("Minimum value is", c1)
}
