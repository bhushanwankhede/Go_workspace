package main


import ( "fmt" )

func main () {
	// rune is alise for int32
	var num rune 
	
	fmt.Printf("Enter the number here : ")
	fmt.Scanf("%d", &num)
	fmt.Printf("%d\t\t%b\t\t%#X\n", num,num,num)
	// num left shift over 1 
	new_var := num << 1 
	fmt.Printf("%d\t\t%b\t\t%#x\n", new_var,new_var,new_var)
}