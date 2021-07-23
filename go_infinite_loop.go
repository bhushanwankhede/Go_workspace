package main 


import ( "fmt" )


func main() {
	bd := 1995
	i := 0
	for {
		if bd > 2080 {
			break
		}
		fmt.Println(i, " :- ",bd)
		bd++; i++
	}

}