package main 

import ( 
	"fmt" 
)

const PI float64 = 3.14159

//square
type square struct{
	side  float64
}

type circle struct{
	radius float64
}

func (c circle) area() float64{
	return (PI * (c.radius * c.radius))
}

func (s square) area() float64{
	return (s.side * s.side)
}

type shape interface{
	area() float64
}

func math_func(a shape)  {

	switch a.(type) {
	case circle:
		fmt.Println("circle  Radius is : ", a.(circle).radius)
	case square:
		fmt.Println("square side is : ", a.(square).side)
	}
	fmt.Println("Area is : ", a.area())
}


func main () {
	var c circle 
	c.radius = 3
	var s square 
	s.side = 4
	math_func(c)
	math_func(s) 
}