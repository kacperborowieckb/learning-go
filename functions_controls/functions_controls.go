package main

import (
	"errors"
	"fmt"
	"log"
)

// pass parameters with type <param_name> <type>
func printString(printValue string) {
	fmt.Println(printValue)
}

// pass many parameters separated by coma, and return type after (or many return using () and separate by come)
func intDivision(numerator int, denominator int) (int, int, error) {
	// ERRORS IS TYPE IN GO (default value is nil)
	// ?? if error is possible always return it as the last return value
	// return nil as error value if everything is good
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}

	result := numerator / denominator

	reminder := numerator % denominator

	return result, reminder, nil
}

func main() {
	message := "Hello, World!"
	printString(message)

	firstNum, secondNum := 6, 0
	result, remainder, err := intDivision(firstNum, secondNum)

	if err != nil {
		log.Fatal(err.Error())
	}

	// with Printf we can use %v for value and than pass the values, it will replace %v
	fmt.Printf("Result: %v with remainder: %v", result, remainder)

	// if syntax in go

	// if condition {
	// } else if condition2 {
	// } else {
	// }

	// and => &&
	// or => ||

}
