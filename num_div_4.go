//write programe to fined divided by 4 (0 to 100)


package main 

import ( "fmt" )

func main() {
	for i:=0; i <= 100; i++ {
		if i % 4 == 0 {
			fmt.Println(i)
		}
	}

}