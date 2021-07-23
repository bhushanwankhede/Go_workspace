package main


import (
	"fmt"
	_ "os" // black import 
)

func main() {
	// this is for only we are using os package 
	//os.Getgid() 
	/*//exercise 1
	ch := make(chan int)

	go func (){
		fmt.Println("ch is push ")
		ch <- 42 
		close(ch)
		fmt.Println("ch is close ")
	}()
	
	fmt.Println("ch is ", <- ch)*/

	/* //exercise 2 
	//sending channel
	ch := make(chan int)

	go func(){
		ch <- 42
	}()

	fmt.Println(<-ch) // err

	fmt.Println("............\n")
	fmt.Printf("ch\t %T \n",ch)*/

	/*  //exercise 4 

	ch1 := make(chan int)
	ch2 := make(chan int)
	quit := make(chan int)

	go func(){
		for i := 0 ; i < 100; i++ {
			if i % 2 == 0{
				ch1 <- i
			} else {
				ch2 <- i
			}
		}
		quit <- 1
		close(ch1)
		close(ch2)
		close(quit)
	}()

	for {
		select {
		case c1 := <- ch1 :
			fmt.Println("channel one value : ", c1)
		case c2 := <-ch2 : 
			fmt.Println("channel two value : ", c2)
		case c3 := <- quit :
			fmt.Println("channel is closed : ", c3)
			os.Exit(0)	
		//default :
			//fmt.Println("meight be channel is close ", <-ch1, <-ch2)
		}

	}*/

	/* // comma ok idiom 
	ch := make(chan int )

	go func (){
		ch <- 42
		close(ch)
	}()

	v, ok := <- ch
	fmt.Println(v, ok)

	v,ok = <- ch
	fmt.Println(v , ok)*/

	/*// put 100 number to achannel 
	// read that number form that channel value */
	 ch := sender()

	 for v := range ch {
	 	fmt.Println(v)
	 }
	


}


// sender function returing receiver channnel 
func sender() <- chan int {
	ch := make(chan int)

	go func (){
		for i := 0; i< 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}