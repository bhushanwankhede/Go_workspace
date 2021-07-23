// wirte program define TYPE and UNTYPE const and print it's value and type


package main 

import ( "fmt" )

const PI = 3.14

const (
	a int = iota
	b float64 = 0.0
	c string = ""
	d byte = 0
	e rune = 0
)

func main () { 
		fmt.Println(PI)
		fmt.Printf("%T\n", PI)
		fmt.Println(a,b,c,d,e)
		fmt.Printf("%T\n%T\n%T\n%T\n%T\n",a,b,c,d,e)
}