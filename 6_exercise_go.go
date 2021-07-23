//write programme to print enter number in dec, bin, and hex formate 

package main

import (
	"fmt"
)

func main() {
	var num int = 0
	fmt.Printf("Enter Number :" )
	fmt.Scanf("%d", &num)
	
	fmt.Printf("Number IN Decimal :%d\n", num)
	fmt.Printf("Number IN Binary :%b\n", num)
	fmt.Printf("Number IN Hexadecimal :%#X\n", num)
}