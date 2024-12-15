package main

import (
	"fmt"
)

// struct is the way of declaring out own types
type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

// in go we can declare some function that will be tight to the struct
// like classes but without class

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

// declare just like struct but instead interface
// now we can use it to say, accept anything that has milesLeft method
type engine interface {
	milesLeft() uint8
}

// what if we want to make this func accept more than one type of Engine
// Interfaces
// we are using engine interface to say that it must have milesLeft() method
func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("Can make it!")
	} else {
		fmt.Println("Can't make it :(")
	}
}

func main() {
	// to create some variable from struct just give it this type
	var myEngineDefault gasEngine

	fmt.Println("Engine with default values:", myEngineDefault.mpg, myEngineDefault.gallons)
	// by default it will contain default values for types => 0 for int etc.

	// to initialize with values
	// we can even omit the values name to declare variables in order
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15}
	fmt.Println("Engine with default values:", myEngine.mpg, myEngine.gallons)

	// with nested types (some type in some type) we can use
	// variable from nested type directly!! that's great
	// myEngine.owner.name => myEngine.name

	// anonymous struct are possible
	var nonReusableEngine = struct {
		mpg     uint8
		gallons uint8
	}{25, 15}

	fmt.Println("nonReusableEngine(declared + initialized in place):", nonReusableEngine)

	// FUNC ON STRUCT (methods)
	// like in classes
	milesLeft := myEngine.milesLeft()

	fmt.Println("Miles left:", milesLeft)
	canMakeIt(myEngine, 100)
	canMakeIt(electricEngine{12, 23}, 100)
}
