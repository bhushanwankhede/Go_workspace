package main

import (
	"fmt"
	
)

//create your own type using "type" keyword
type own_t int 
var x own_t

func main () {

	x = 43
	fmt.Println(x)
	fmt.Printf(" Type of the variable is : %T\n", x)

}