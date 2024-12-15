package main

import "fmt"

func main() {
	// Arrays
	// - Fixed Length
	// - Same Type
	// - Indexable
	// - Contiguous in Memory

	// var <variable_name> [<length>]<type>
	// it will fill this array with default values
	// in case in int it is 0
	var intArr [3]int32

	fmt.Println("First element in intArr:", intArr[0])

	// we can use syntax like in python [start:end(exclude):step]
	fmt.Println("Element [1:3] from index 1 to 3 without 3", intArr[1:3])

	// we can use & to access the memory location of an item
	fmt.Println("Memory address of first element:", &intArr[0])

	// to immediately initialize an array
	// var intArr [3]int32 = [3]int32{1, 2, 3}
	// OR
	// intArr := [3]int32{1, 2, 3}
	// OR
	// int Arr := [...]int32{1, 2, 3}
	// it will know that this array has length of 3 (infer from elements)

	// SLICES
	// slices are wrapper of arrays,
	// so they are arrays with some additional/better functionality

	// we can omit the length and now we have a SLICE
	var slice []int32 = []int32{1, 2, 3}

	// we can append an element to a slice using the append method
	intSlice := append(slice, 4)
	// it will also return the updated slice
	// it will check if there is a space for a new value
	// in this case we have [1, 2, 3] so no more space for 4
	// it will create new array with enough space (with new capacity) for it and return it
	fmt.Println("Slice after append 4:", intSlice)

	fmt.Printf("The length is %v, the capacity is %v\n", len(intSlice), cap(intSlice))
	// we can check the length => number of elements
	// or for capacity => how many values can be stored in the same slice

	// we can append multiple values using spread operator
	// intSlice = append(intSlice, intSlice2...)

	// we can create a slice with make() function

	// make(<type>, <length>, <capacity>)
	sliceWithMake := make([]int32, 3, 8)
	fmt.Println("Slice created with make():", sliceWithMake)

	// MAPS
	// key value pairs

	// we can create it using the make function
	// map[<key_type>]<value_type>
	myMap := make(map[string]uint8)
	fmt.Println("Map:", myMap)

	// to initialize immediately:
	// map[string]uint8{"Adam": 23, "John": 54}

	// if we tried to access the value that do not exist in the map
	// we will get the default value of value type, int this case 0
	// it will also return the second value that is bool if key exists

	val, exists := myMap["someKey"]

	fmt.Println("Value:", val, "Exists:", exists)

	// to delete some key use:
	// delete(<map>, <key>)
	// no return value for this

	// loops

	// iterating over a maps do not preserve order!!

	// range <map, slice or array> to iterate
	// with maps => key, value := range myMap
	// with arrays, slices => index, value := range mySlice
	for name, age := range myMap {
		fmt.Println("Name:", name, "Age:", age)
	}

	// no while loops, instead:
	// for i < 10 {} or sth like this
	// without condition is infinite loop
}
