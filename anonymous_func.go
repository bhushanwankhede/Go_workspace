package main

import ( "fmt" )

func main() {

	//fmt.Println(foo()())
	x := foo()
	fmt.Printf("%T\n", x)
	fmt.Println(x())
}

//function that returns function 

func foo() func() int{
	fmt.Println("I am in foo\n")
	return func()int{
		return 1995
	}
}