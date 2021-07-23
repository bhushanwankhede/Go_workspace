package main 


import (
	"fmt"
	"sync"
)
//fanin is a tequenic that use in conccerency 
// fanin means pull data from  two or more channles and push into one channele 
// it's called fanIN
//  ch1 --------
//				\
//			     |-->fanIN-->
//				/
//  ch2 --------

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	fanin := make(chan int)


	go sender(ch1, ch2)

	go receiverAndFanIn(ch1, ch2, fanin)

	for v := range fanin {
		fmt.Println(v)
	}
}

func sender(ch1, ch2 chan<- int) {
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			ch1 <- i
		} else {
			ch2 <- i
		}
	}
	close(ch1)
	close(ch2)
	//if we are not close ch1 and ch2 go comiler throws error deadlock
}

func receiverAndFanIn(ch1, ch2 <- chan int, fanin chan <- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		for v := range ch1{
			fanin <- v
		}
		wg.Done()
	}()

	go func(){
		for v := range ch2 {
			fanin <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanin)

}