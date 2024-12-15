// package name, the same for every file in the folder (main here just because)
// normally it will look for main func in main package to start running a program
package main

// importing packages, many with parentheses or single without
import (
	"fmt"
	"unicode/utf8"
)

// Go is statically and strongly typed,
// declare or infare, cannot add num to string

// Go is compiled
// BLAZINGLY FAST
// concurrency (Goroutines)
// simple

// MODULES AND PACKAGES

// Packages => folders with go files
// Modules => collection of packages

// to run it(build and run): `go run <file>`
// to build it: `go build <file>`
func main() {
	// in Go we can declare a variable by using:
	// var <var_name> type = <value>
	var num int = 2
	var numInfer = 4

	// or we can declare it dynamically(it will infer the type) by using:
	// <var_name> := <value>
	numShortInfer := 2
	stringInfer := "Hello, world!"

	// we can have const variables also,
	// const <var_name> <type> = <value>
	// we can omit the type and it will infer it
	const constNum int = 4

	// cannot change type in later assignment
	// stringDynamic = 2 it doesn't work

	fmt.Println("Declared with type:", num)
	fmt.Println("Var declaration infer:", numInfer)
	fmt.Println("Infer int declaration:", numShortInfer)
	fmt.Println("Infer string declaration:", stringInfer)
	fmt.Println("Constant:", constNum)

	// Go is strongly types
	// cannot add two different types

	// it doesn't work
	// somVar := num + stringInfer

	// basic data types in Go:
	// bool
	// float32 float64
	// int int16 int32 int64
	// rune (char?)
	// string
	// uint uint8 uint16 uint32 uint64

	// Type casting:
	// <type>(<variable>)
	var floatFromInt float64 = float64(num) + 0.24
	fmt.Println("Float from int:", floatFromInt)

	// Strings
	// with double quotes we can't create new lines etc.
	doubleQuotesString := "Some string"

	// with backticks we can do anything
	rawString := `Some raw string\nsom
	e .       d`

	fmt.Println("With double quotes:", doubleQuotesString)
	fmt.Println("Raw string", rawString)

	// Getting length of string
	fmt.Println("length using `len():`", len(doubleQuotesString))
	// It is not the number of chars, it is the number of bytes

	// when we are dealing with some fancy chars use this:
	fmt.Println("Length using utf.RuneCountInString()", utf8.RuneCountInString(doubleQuotesString))

	// Runes
	var someRune rune = 'a'
	fmt.Println("Rune 'a':", someRune) // 97

	// variable declarations without value

	var defaultInt int
	fmt.Println("Default int value:", defaultInt) // it will be 0
	// when value is not declared, it will give the default value
	// for all int, float, uint, rune => 0
	// for string => ''
	// bool => false

	// multiple value assignment:
	var one, two = 1, 2
	fmt.Println("Multiple assignment:", one, two)
}
