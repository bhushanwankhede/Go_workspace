
package main 

import ( "fmt" )

type vehicle struct {
	door int8
	color string 
}

type truck struct {
	vehicle
	fourwheeler bool
}

type bmw struct {
	vehicle
	luxury bool
}

func main () {
	t := truck {
		vehicle: vehicle{
			door: 2,
			color: "red",
		},
		fourwheeler: true,
	}
	
	fmt.Println(t)
}