package main

import (
	"fmt"
)

func input(x []int, err error) []int {
    if err != nil {
        return x
    }
    var d int
    n, err := fmt.Scanf("%d", &d)
    if n == 1 {
        x = append(x, d)
    }
    return input(x, err)
}

func main() {

	fmt.Println("Enter Element :")
	x := input([]int{}, nil)

	fmt.Println(x,len(x))
	PrintNGE(x)
	fmt.Println("PrintNGE_using_stack : ")
	PrintNGE_using_stack(x)
}

func PrintNGE(a []int) {

	for i := 0; i < len(a); i++ {
		next := -1;
		for j := i + 1; j < len(a); j++{
			if a[j] > a[i] {
				next = a[j]
				break
			}
		}
		fmt.Println(a[i], "-.- ", next)
	}
}

type (
	Stack struct {
		top *node
		length int
	}
	node struct {
		value interface{}
		prev *node
	}	
)
// Create a new stack
func New() *Stack {
	return &Stack{nil,0}
}
// Return the number of items in the stack
func (this *Stack) Len() int {
	return this.length
}
// View the top item on the stack
func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}
// Pop the top item of the stack and return it
//func (this *Stack) Pop() interface{} {
func (this *Stack) Pop() {
	if this.length == 0 {
		return
	}
	
	//n := this.top
	//this.top = n.prev
	this.top = this.top.prev
	this.length--
	//return n.value
}
// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {
	n := &node{value,this.top}
	this.top = n
	this.length++
}

func (this *Stack) Is_empty() bool {
	return this.length == 0
}

func PrintNGE_using_stack(a []int) {
	s := New()

	for i := 0; i < len(a); i++ {
		if s.Is_empty() {
			s.Push(a[i])
			continue
		}

		for ;(s.Is_empty() == false) && (s.Peek().(int) < a[i]); {
			fmt.Println(s.Peek().(int)," -.- ", a[i])
			s.Pop()
		}

		s.Push(a[i])
	}

	for ;!s.Is_empty(); {
		fmt.Println(s.Peek().(int), " -.- ", "-1")
		s.Pop()
	}
}