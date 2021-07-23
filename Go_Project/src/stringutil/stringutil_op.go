package stringutil 

//Hi Hello
//i -> H
// j -> o
//i -> o , j -> H

//
func Reverse(s string) string {
	r := []rune(s)

	for i, j := 0, (len(r) -1); i < len(r)/2; i, j =  i + 1, j - 1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}

//a = 2 
// b = 3
// b = a ^ b
// a = a ^ b
// b = a ^ b
//
func Swapstring(dst, src string) (string, string){
	size := 0
	var d, s []rune
	if len(src) == len(dst) {
		size = len(src) - 1
		d = []rune(dst)//a
		s = []rune(src)//b
	} else if len(src) > len(dst) {
		size = len(src) 
		d = make([]rune, size)
		s = make([]rune, size)
		//fmt.Println(len(d), len(s))
		copy(d, []rune(dst))
		copy(s, []rune(src))
		//fmt.Println(len(d), len(s))
	} else if len(src) < len(dst) {
		size = len(dst) 
		d = make([]rune, size)
		s = make([]rune, size)
		//fmt.Println(len(d), len(s))
		copy(d, []rune(dst))
		copy(s, []rune(src))
		//fmt.Println(len(d), len(s))
	}
	
	for i := 0; i < size; i++ {
		s[i] = xro_op(d[i], s[i])
	}
	for i := 0; i < size; i++ {
		d[i] = xro_op(d[i], s[i])
	}
	for i := 0; i < size; i++ {
		s[i] = xro_op(d[i], s[i])
	}
	//fmt.Println("s :" , string(s) , "d : ", string(d))
	return string(d), string(s)
}

func xro_op(a, b rune) rune {
	return a ^ b
}

func ConcatString(str1, str2  string) string {
	return str1 + str2
}



