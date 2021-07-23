//create slice of slice of string ([][]string) 


package main 

import ( "fmt" )

func main () {
	var s_of_s_of_str [][]string
	
	s_of_s_of_str = make([][]string, 5, 10)
	
	x := []string{"Bhushan", "Makan", "Patil"}
	y := []string{"Jotsanabai", "Makan", "Patil"}
	
	s_of_s_of_str = [][]string{x,y}
	
	fmt.Println(s_of_s_of_str)

}