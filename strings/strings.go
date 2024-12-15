package main

import (
	"fmt"
	"strings"
)

func main() {
	// strings can be indexed
	var name string = "John"

	// but this will return the uint8 with some number
	fmt.Println("First letter of name: ", name[0])

	// each letter will have some representation of bytes
	// it will depend on the char (UTF-8)
	// some emojis will take e.g two places in array

	// when using range to iterate through string
	// it will do some extra work and encode correctly
	// that some char will take two places

	// we can instead use runes array
	// runes => number that will represent a char
	// in other words, just alias for int32
	var runesArr = []rune("SomeRune")
	fmt.Println(runesArr)

	// we can declare a single rune type by putting char in ''
	var singleRune = 'a'
	fmt.Println("Single Rune:", singleRune)

	// string are immutable in go!
	// hm

	// we can build out string using strings package
	var stringBuilder strings.Builder
	stringBuilder.WriteString("Some string")

	fmt.Println("String with string builder:", stringBuilder.String())
}
