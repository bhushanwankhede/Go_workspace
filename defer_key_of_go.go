package main


import ( "fmt" )

func call_last() {

	fmt.Println("Hi i am call_last func and i should called at end of main ")
}

func name() {
	fmt.Println("name of the code is : defer keyword ")
}

func version() {
	fmt.Println("version of code is : 1.0")
}

func main() {
	defer call_last()
	name()
	version()
}