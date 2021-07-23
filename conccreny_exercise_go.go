package main 


import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	 go foo()
	 go bar()
	 for i := 0; i < 100; i++{
		fmt.Println("printing in main routine : ", i)
	}
	 wg.Wait()
	 fmt.Println("Code Exit...")


}


func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("printing in foo routine : ", i)
	}
	wg.Done()
}


func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("printing in bar routine : ", i)
	}
	wg.Done()
}