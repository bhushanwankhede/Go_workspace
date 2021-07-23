package main 


import (
	"fmt"
	//"sync"
)



/*In Go language, a channel is a medium through which a goroutine communicates 
with another goroutine and this communication is lock-free. 
Or in other words, a channel is a technique which allows to let one goroutine 
to send data to another goroutine. By default channel is bidirectional, 
means the goroutines can send or receive data through the same channel*/

// var idenfier-name chan Type [int, string, map, floate64 and so on]
//var channel_var chan  int
//var wg sync.WaitGroup

func main() {

	channel_var := make(chan int, 1)
	//wg.Add(2)
	fmt.Println("master_chan go routine is lonched ")
	go master_chan(channel_var)

	slave_chan(channel_var)
	//wg.Wait()
	fmt.Println("main routine exit ")
	close(channel_var)

} 

//sender only
func master_chan(ch chan <- int) {
	i := 1
	for {
		ch <- i
		i++
		//fmt.Println("master_chan %d ", <- ch)
		//if val, isset := <- ch; isset {
		//	if val == 5 {
			//	break
			//}
		//}
	}
	fmt.Println("master_chan close..")
	//wg.Done()
	//close(ch)
}
//receiver only
func slave_chan(ch <- chan int) {
	
	for {
	 
		if val, isset := <- ch; isset {
			if val == 5 {
				break
			} 
			fmt.Println("slave_chan value is : " ,<-ch)
		}
	}
	//wg.Done()
	//close(ch) : err : invalid operation: close(ch) (cannot close receive-only channel)
	fmt.Println("slave_chan exit...")
}