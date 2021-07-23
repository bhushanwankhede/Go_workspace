//Array

package main

import ( "fmt" )

func main () {
	
	arr := []int{1,2,3,4,5}
	
	fmt.Println(arr)
	
	fmt.Println("using for loop ")
	
	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i])
	}
	
	fmt.Println("using for loop  with range")
	
	for i , v := range arr {
		fmt.Println(i, v)
	}
	fmt.Println("Type of the Array : ")
	
	fmt.Printf("%T\n",arr);
	
}