package main

import (
	"fmt"
)

var x int 
var y string
var z bool
//q :-> The compiler assigned value to the variables , what these called?
// ans :-> Zero value
func main() {
	fmt.Println(x,"\n",y,"\n",z)
	fmt.Printf("%T\t%T\t%T\n", x,y,z)
}
