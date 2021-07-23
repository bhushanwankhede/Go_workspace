package  calculator


//Add two given number 
func Add(a,b int) int {
	return a + b
}


//Sub two number's 
func Sub(a,b float64) float64 {
	return a - b
}

// Multiple two given number 
func Mul(a,b rune) rune{
	return a * b
}

//Divide two given number 
func Divide(a, b float64) float64 {
	if b == 0 {
		return -1
	}
	return a / b
}

