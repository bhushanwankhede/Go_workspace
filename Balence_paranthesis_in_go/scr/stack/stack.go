package stack
import "fmt"
type (
	Stack struct {
		top *node
		length int
		capacity int
	}
	node struct {
		value interface{}
		prev *node
	}	
)

// Create a new stack
func New(capacity int) *Stack {
	return &Stack{nil,0,capacity}
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
func (this *Stack) Pop() {
	if this.length == 0 {
		return
	}
	this.top = this.top.prev
	this.length--
}
// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {
	if this.length == this.capacity {
		fmt.Println("[ERROR] : Stack is full...!")
		return
	}
	n := &node{value,this.top}
	this.top = n
	this.length++
}

func (this *Stack) Is_empty() bool {
	return this.length == 0
}
