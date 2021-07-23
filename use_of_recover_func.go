package main 



import (
	"fmt"
)

func main () {
	foo()
	fmt.Println("return normally from foo")
}


func foo() {
	//check how this work 
	defer func(){
		//comma ok idiom sytle
		if r := recover(); r != nil {
			fmt.Println("recovering in foo : ", r)
		}
	}()
	fmt.Println("calling func get  ")
	get(0)
	fmt.Println("return normally from get")
}


func get(i int ) {
	if i > 3 {
		fmt.Println("Panicing here !")
		panic(fmt.Sprintf("%v",i))
	}
	defer fmt.Println("Defer in get : ",i)
	fmt.Println("print in get :",i)
	get(i+1)
}