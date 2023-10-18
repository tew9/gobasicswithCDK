package main

import (
	"fmt"
)

func main(){
	// the only loop allowed in go is the for loop	
	// explicit loop declaration gives us FOR LOOP
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//implicit loop declaration gives us WHILE LOOP
	j := 0

	for j < 5 {
		fmt.Printf("while loop: %v\n", j)
		j++
	}

	//same loop to iterate over an array
	for index, value := range []int{1,2,3,4,5} {
		fmt.Println("index:", index, "value:", value)
	}

	//same loop to iterate over a map
	for key, val := range map[string]int{"key1": 1, "key2": 2, "key3": 3} {
		fmt.Println("key:", key, "value:", val)
	}
}