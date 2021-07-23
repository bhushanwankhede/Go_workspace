package main 


import (
	"fmt"
	"src/calculator"
	"src/stringutil"
)



func main () {
	fmt.Println("Calculator Programme")

	fmt.Println("Add : ", calculator.Add(12,10))
 	fmt.Println("Sub : ", calculator.Sub(12,10))
  	fmt.Println("Mul : ", calculator.Mul(12,10))
   	fmt.Println("divide : ", calculator.Divide(12,10))

   	fmt.Println(stringutil.Reverse("Hi I am Bhushan"))

   	fmt.Println(stringutil.Swapstring("Bhushan", "Kavita"))
   	fmt.Println(stringutil.ConcatString("Bhushan", "Kavita"))
   }
