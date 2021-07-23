// go "if else statement"

package main 

import ( "fmt" )

func main() {
	var num int 
	fmt.Printf("Enter Number : ")
	fmt.Scanf("%d", &num)
	
	if num == 3 {
		fmt.Println("Enter Number is Match And Your The Winner.......!")
	} else {
		fmt.Println("Please Try Again Enter Number Did not Match")
	}
}
 
