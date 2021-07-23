package main 


import (
	"fmt"
	"runtime"
)


func main() {
	fmt.Println("os : ", runtime.GOOS)
	fmt.Println("No . of CPUs : ", runtime.NumCPU())
}