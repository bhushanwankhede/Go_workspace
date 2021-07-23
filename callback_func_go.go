package main 


import ( "fmt" )

func sum(num ...int) int {
	if len(num) <= 0 {
		return -1
	}
	var total int
	for _, v := range num {
		total += v
	} 
	return total
}

func even(Cfunc func(num ...int) int, num ...int) int {
	if len(num) <= 0 {
		return -1
	}
	var s_var []int
	for _ , v := range num {
		if v % 2 == 0 {
			s_var = append(s_var, v)
		}
	}

	return Cfunc(s_var...)
}


func main() {
	i := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
	fmt.Println(even(sum,i...))
}