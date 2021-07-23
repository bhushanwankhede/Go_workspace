
package main 

import ( "fmt" ) 

type porson struct {
	fristName string 
	lastName string
	favorate []string
}


func main () {
	 p1 := porson{
			fristName: "Bhushan",
			lastName: "Patil",
			favorate: []string{"I","like","write code", "c", "c++", "go","so on"},
	}
	 p2 := porson{
			fristName: "Bhushan",
			lastName: "Patil",
			favorate: []string{"I","like","write code", "c", "c++", "go","so on"},
	}
	fmt.Println(p1, p2)
	
	//creare map using porson 
	//note : map not contain dublicate key entries
	map_var := map[string]porson{
						p1.lastName : p1,
						p2.lastName : p2 }
						
	fmt.Println(map_var)
	
	//range over it 
	
	for k, v := range map_var {
		fmt.Println(k)
		fmt.Println(v.fristName,v.lastName)
		for i , v := range v.favorate {
			fmt.Println(i, v)
		}
	}
}