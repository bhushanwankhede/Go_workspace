package main


import (
	"os"
	"bufio"
	"fmt"
	"strings"
)

//32 to 127 make map map[byte]rune

func Make_map(cap rune) map[byte]rune {
	m := make(map[byte]rune, cap)

	for i := 23; i < 127; i++ {
		m[byte(i)] = 0;
	}

	return m
 }

 func Print_char_freq(m map[byte]rune) {
 	for i := range m{
 		if v, isset := m[byte(i)]; isset && v != 0{
 			fmt.Printf("%#U == %d \n",i, m[byte(i)])
 		}
 	}
 }

 func Count_char_freq(str string) map[byte]rune {
 	m := Make_map(128)
 	fmt.Println()
 	for i := 0; i < len(str); i++ {
 		m[str[i]] = m[str[i]]+1
 	} 
 	return m
 }
 func main() {
 	fmt.Println("Enter String : ")
 	reader := bufio.NewReader(os.Stdin)
 	str, _ := reader.ReadString('\n')

 	m := Count_char_freq(strings.Trim(str, "\r\n"))
 	Print_char_freq(m)
 }


