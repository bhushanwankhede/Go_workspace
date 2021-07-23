package main

import (
	"fmt"
	
)
//conversion of costum type 
type own_t int 

var x own_t
var y int 

func main () {

	x = 43 
	fmt.Printf("Type of x is : %T\n", x)
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("Type of y is : %T\n", y)
}