package main

import (
	"fmt"
	"errors"
)

func main(){
	fmt.Printf("sum: %v\n", sum(5, 7))

	res, err := swap(5, 5)
	fmt.Printf("swap: %v\n", res, err)

}

// function declaration with parameters and return type
func sum(x int, y int) int {
	return x + y
}

// go function can return multiple values
func swap(x int, y int) (string, error) {
	if x == y {
		return "", errors.New("x and y are the same")
	}else if y < x{
		return "y is less than x", nil
	}
	return "swapped them", nil
}