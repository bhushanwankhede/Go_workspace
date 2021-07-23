package main 

import ( "fmt" )

type person struct{
	frist_name string
	last_name string
	age int8
}

//method attached with porsion
func (p person) speek() {
	fmt.Println("Hello every one I am ", p.frist_name," ",p.last_name," and my age is ",p.age)
}


func main() {
	p1 := person{
		frist_name: "Bhushan",
		last_name: "Patil",
		age: 26,
	}
	p2 := person{
		frist_name: "unknown",
		last_name: "someting",
		age: 00,
	}
	p1.speek()
	p2.speek()
}