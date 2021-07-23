//for loop 
// Note : In go lang only for loop is available 
// for [ conditinal | for [like c] | rangeClause ]

// while like loop 
// i := 0
// for i < some_num {
//     some code block
//		i++ or i = i + 1 
// }

//for i := 0; i < num; i++ { }

//for key, value := range data { }

//write programe to print number 1 to 10,000


package main

import ( "fmt" )

func main() {
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}
}