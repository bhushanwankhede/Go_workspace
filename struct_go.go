//struct in go 



package main 


import ( "fmt" )

type porson struct {
	first string 
	last string 
	age int 
	city string 
}

//Embedded struct 

type S_porson struct {
	porson
	isset bool
}
func main () {

	p1 := porson{
			first: "Bhushan",
			last: "Patil",
			age: 26 ,
			city: "Vishwanath", 	
		}
	
	var p2 porson
	
	p2.first = "Kavita"
	p2.last = "Patil"
	p2.age = 23
	p2.city = "Vishwanath"
	
	fmt.Println(p1, p2)
	//embedded struc 
	
	var p3 S_porson
	
	p3.first = "Makan"
	p3.last = "Patil"
	p3.age = 67
	p3.city = "Vishwanath"
	p3.isset = true
	
	fmt.Println(p3)
	
	
	//anonymous struct : structure with no identifer  
	
	s := struct {
		first string
		friend map[string]rune
	}{
		first : "Bhushan",
		friend: map[string]rune {
						"kavita": 12345,
						"Dinesh" : 45345,
						"me": 342445,
					},
	}
	fmt.Println(s.first)
	for k, v := range s.friend {
			fmt.Println(k,v)
	}
	

}