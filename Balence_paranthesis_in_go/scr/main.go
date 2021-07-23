package main 


import (
	"fmt"
	"src/stack"
)

var sym = map[byte]int{
	40 : -1,
	41 : 1,
	91 : -2,
	93 : 2,
	123 : -3,
	125 : 3,
}

func Is_balence(str string) bool {
	st := stack.New(256)

	//fmt.Println(st, "str" ,str, str[0]))
	for i := 0; i < len(str); i++ {
		if v, isset := sym[str[i]]; isset && v < 0 {
			st.Push(str[i])
			//fmt.Println(isset, v, str[i])
		} else {
			if st.Is_empty() {
				return false
			}else if sym[st.Peek().(byte)] + sym[str[i]] == 0 {
				st.Pop()
			}
		}
	}

	if (st.Is_empty()) {
		return true
	}
	return false
}
func main() {

	fmt.Println("Enter Parenthesis string :" )

	var str string;

	fmt.Scanf("%s", &str)

	fmt.Println(str)
	fmt.Println(Is_balence(str))
}