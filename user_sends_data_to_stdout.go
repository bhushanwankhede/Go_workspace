package main 

import (
	"fmt"
	"encoding/json"
	"os"
)

type user_info struct {
	Frist string
	Last string
	Email string
	Age int
	Mobilenumber rune
}

// convert  go struct to json and snd to stdout using 
// func NewEncoder(w io.Writer) *Encoder
// func (enc *Encoder) Encode(v interface{}) error
// func (enc *Encoder) SetEscapeHTML(on bool)
// func (enc *Encoder) SetIndent(prefix, indent string)

func main () {
	u1 := user_info{
		"Bhushan", "Patil", "patilbhushan8595@gmail.com", 27, 123456789,
	}
	u2 := user_info{
		"Dinesh", "Sonwane", "dsb@gmail.com", 30, 123456789,
	}
	var uesrs []user_info
	uesrs = []user_info{u1, u2}
	fmt.Println(uesrs)
	/*bs, err := json.Marshal(uesrs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))*/

	err := json.NewEncoder(os.Stdout).Encode(uesrs)
	if err != nil {
		fmt.Println(err)
	}	
}