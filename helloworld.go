package main 

import (
	"fmt"
)

func main(){
	// explicit type declaration
	var x int = 5
	var y int = 7
	var sum int = x + y

	//implicit type declaration
	t := 5
	w := 8

	// if statement in go doesn't need parentheses
	if t < w {
		fmt.Println("t is less than w")
	} else {
		fmt.Println(sum)
	}
}
