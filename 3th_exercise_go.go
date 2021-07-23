package main

import (
	"fmt"
)

var x int = 43
var y string = "Bhushan patil"
var z bool = false 

func main() {
	s := fmt.Sprintf("%d\t%s\t%t\n", x,y,z)
	fmt.Println(s)
}
