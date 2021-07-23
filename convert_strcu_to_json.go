package main 

import  (
	"fmt"
	"encoding/json"
	"os"
)


//Note : for converting struct into json ,
// structure member name should start with upper case leter 
type user_info struct {
	Frist string
	Last string
	Email string
	Age int
	Mobilenumber rune
}


func main () {
	u1 := user_info{
		"Bhushan", "Patil", "patilbhushan8595@gmail.com", 27, 123456789,
	}
	fmt.Println(u1)
	bs, err := json.Marshal(u1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
	fmt.Println()
	os.Stdout.Write(bs)
}