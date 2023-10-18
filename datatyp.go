package main

import (
	"fmt"	
)

//struc is a collection of fields that organizes datatype
type person struct {
	name string
	age int
}

func main(){
	var explicitArray [5]int // explicitly declare an array of 5 integers with default value of 0
	explicitArray[4] = 100 // set the 5th element to 100
	fmt.Println("emp:", explicitArray)

	implicitArray := [4]int{1,2,3,4} //implicit type declare and initialize an array of 5 integers
	fmt.Println(implicitArray)

	sliceArray := []int{1,2,3,4,5} // slice is a dynamically-sized array
	//why useful: slice allow appending to the end of the array which creates
	// a new array with the appended value in the background
	sliceArray = append(sliceArray, 7)
	fmt.Printf("sliceArray: %v\n", sliceArray)

	// map is a key-value pair
	myMap := map[string]int{"key1": 1, "key2": 2} //have to be initialize or else use make since it'll create a nil data structure which we can't add to
	fmt.Println(myMap)

	//explicit type declaration
	vertices := make(map[string]int) //make is a keyword to create an empty dictionary
	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12

	//delete a key-value pair
	delete(vertices, "square")

	fmt.Println(vertices)

	//initialize a struct
	p := person{name: "Jake", age: 23}
	println(p.name, p.age)

}