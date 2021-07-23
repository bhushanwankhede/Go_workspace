//map exercise 


package main 

import ( "fmt" )

func main () {
	//map[kayType]valueType
	map_var := map[string]string {"1" : "Bhushan",
								"2" : "Makan",
								"3" : "Patil"}
	fmt.Println(map_var)
	
	for k, _ := range map_var {
		fmt.Println(k)
		if v, isset := map_var[k]; isset {
			fmt.Println(v)
		}
	}

}