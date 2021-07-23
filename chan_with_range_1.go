package main 

import (
	"fmt"
)

func main() {

   ch := make(chan int)

   //sender routine 
   fmt.Println("anonymaous go routine is launched ... ")
   go func(){
   		for i := 0; i < 10; i++{
   			ch <- i
   			//fmt.Println("anonymaous routine : ",<-ch) :err : dead-lock
   		}
   		close(ch)
   }()

   //Thise loop is alive until channle close by sender routine 
   for v := range ch {
   		fmt.Println(v)
   }

}