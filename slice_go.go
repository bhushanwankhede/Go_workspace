// slice 
// Note in go use slice instead of array 


package main

import ( "fmt" )

func main() {
	// short hand oprator , declearation with initilisation 
	slice_var := []int{1,2,3,4,5,6,7,8,9,10}
	
	fmt.Println(slice_var)
	
	for i, v := range slice_var {
		fmt.Println(i,v)
	}
	
	fmt.Printf("%T\n", slice_var)
	
	//use ':' oprator for slicing data
	//output should be bellow given
	//1-> [1,2,3,4]
	//2-> [5,6,7,8]
	//3-> [3,4,5,6]
	//4-> [ 7,8,9,10]
	fmt.Println("array slicing ")
	fmt.Println(slice_var[:4])
	fmt.Println(slice_var[4:8])
	fmt.Println(slice_var[2:6])
	fmt.Println(slice_var[6:])
	
	//use "append" func  for append value into the slice 
	//func append (s []T, x ...T) []T --> func defination 
	// append value 11,  int slice_var variable using append func
	
	slice_var = append(slice_var, 11)
	fmt.Println(slice_var)
	//append 3 more value into some slice 
	
	slice_var = append(slice_var, 12,13,14)
	fmt.Println(slice_var)
	
	//create new slice 
	
	x := []int{15,16,17,18,19,20}
	fmt.Println(x)
	
	//append x slice into the slice_var using "dst = append(dst, src...)"
	slice_var = append(slice_var, x...)
	fmt.Println(slice_var)
	
	fmt.Println("Delete element from slice 12, 13, 14")
	//for delete element from slice we use were simple method 
	// we use append func and : oprator
	// Delete 12, 13, 14 from given slice 
	
	slice_var = append(slice_var[:11], slice_var[14:]...)
	fmt.Println(slice_var)

}