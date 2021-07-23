//write programe to use all reletinal/conditinal opratore of go 


package main 

import (
	"fmt"
)

func main() {
	a := (43 == 43)
	b := (43 <= 45)
	c := (43 >= 30)
	d := (43 != 43)
	e := (43 > 43)
	f := (43 < 43)
	
	fmt.Printf("%t\t%t\t%t\t%t\t%t\t%t\t\n",a,b,c,d,e,f)

}