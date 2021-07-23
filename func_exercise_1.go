package main 

import ( "fmt" )
//variadic parameter 
func foo(x ...int) int {
	if len(x) <= 0 {
		return -1
	}
	var total int
	for _, v := range x {
		total += v
	}

	return total
}

func bar(x []int) int {
 	if len(x) <= 0 {
 		return -1
 	}
 	var total int
 	for _, v := range x {
 		total += v
 	}

 	return total
}

func main() {
	slice_var := []int{1,2,3,4,5,6,7,8,9,}
	//var slice_var []int

	fmt.Println("Sum return from foo func : ", foo(slice_var...)) //slice_var.. called unfirling
	fmt.Println("sum return from bar func : ", bar(slice_var))
}