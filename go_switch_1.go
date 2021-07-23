//write programe "switch statement without' expression


package main

import ( "fmt" )

func main() {
	var num int
	
	fmt.Printf("1 -> Hi\t 2 -> Hello\t3 -> Hey\n : ")
	fmt.Scanf("%d",  &num)
	
	switch {
		case (num == 1):
			fmt.Println("Hi :)")
		case (num == 2):
			fmt.Println("Hello .<)")
		case (num == 3):
			fmt.Println("Hey :^")
		default : 
			fmt.Println(":( Try Again ")
	
	}

}