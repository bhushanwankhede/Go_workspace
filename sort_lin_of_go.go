package main


import (
	"fmt"
	"sort"
)


type porson struct {
	Name string
	Age int 
}
type Byname []porson

type Byage []porson

func (bn Byname) Len() int           { return len(bn) }
func (bn Byname) Less(i, j int) bool { return bn[i].Name < bn[j].Name }
func (bn Byname) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }

func (a Byage) Len() int { return len(a)}
func (a Byage) Less(i, j int) bool {return a[i].Age < a[j].Age}
func (a Byage) Swap(i, j int)  {a[i], a[j] = a[j], a[i]}

func main() {

	var peopls []porson

	p1 := porson{"Bhushan", 27}
	p2 := porson{"Dinesh", 30}
	p3 := porson{"Datta", 28}
	p4 := porson{"suraj", 32}
	p5 := porson{"bhupi", 25}
	p6 := porson{"shubham", 33}
	p7 := porson{"vikrant", 34}
	p8 := porson{"vishal", 37}

	peopls = []porson{p1,p2,p3,p4,p5,p6,p7,p8,}
	fmt.Println(peopls)
	sort.Sort(Byname(peopls))
	fmt.Println(peopls)

	sort.Sort(Byage(peopls))
	fmt.Println(peopls)

}