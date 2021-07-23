// "if statement" with  variable initilastion


package main 


import ( "fmt" )

func main ()  {
	
	map_var := map[int] string {
		1 : "Bhushan",
		2 : "Makan",
		3 : "Patil",
	}
	
	fmt.Println(map_var)
	
	fmt.Println(map_var[1])
	
	for i := 0; i < 5; i++ {
		
		if v, ok := map_var[i]; ok {
			fmt.Println(i, v)
		} else {
			fmt.Printf("for key %d value not fund \n",i)
		}
	}
	
	//add value into map 
	
	map_var[0] = "Kavita"
	map_var[4] = "Bhushan"
	map_var[5] = "Patil"
	
	// map with range
	for v, k := range map_var {
		fmt.Println(k, v)
	
	}
	
	// delete  value from map 
	// use delete func to delete value from map
	// delete(map, key)
	
	 if v, ok := map_var[0]; ok {
		fmt.Println(v)
		delete(map_var, 0)
	 }
	 
	 fmt.Println(map_var)
	 
	 //use make for map 
	 mm_map := make(map[int]string, 10)
	 
	 fmt.Println(len(mm_map))
	// fmt.Println(cap(mm_map))


}