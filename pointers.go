package main

import (
	"fmt"
)

func main(){	
	i := 3
	// & is the memory address of the variable
	// * is the value of the memory address
	// &i is the memory address of i
	increment(&i) 
	fmt.Println("i:", i) //

}

//You need to add * show that we will reference the memory address of x
func increment(x *int){
	*x++
}
