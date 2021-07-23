//write program using "iota" for the last 4 years and print it 



package main 

import ( "fmt" )

const (
	y_1 = (2018 + iota)
	y_2 = (2018 + iota)
	y_3 = (2018 + iota)
	y_4 = (2018 + iota)
)

func main () {
	fmt.Println(y_1)
	fmt.Println(y_2)
	fmt.Println(y_3)
	fmt.Println(y_4)
}