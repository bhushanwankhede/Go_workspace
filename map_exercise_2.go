
//create map[string] []string 



package main 

import ( "fmt" )

func main () {

	//map_arr := make(map[string][]string, 100)
	
	map_arr := map[string][]string{ "1" : []string{"Makan", "Girdhar", "Patil"},
				"2" : []string{"Bhushan", "Makan", "Patil"},
				"3" : []string{"Jotsanabai", "Makan", "Patil"},
			}
	fmt.Println(map_arr)
	
	
	for k, _ := range map_arr {
		if v, isset := map_arr[k]; isset {
			fmt.Println(k,v)
		}
	
	}
	
	
	//aad recoed in map 
	map_arr["4"] = []string{"Kavita", "Bhushan", "Patil"}
	
	for k, _ := range map_arr {
		if v, isset := map_arr[k]; isset {
			fmt.Println(k,v)
		}
	
	}
	
	//delete "4" recoed from map_arr
	//func delete(map, key)
	
	key := "4"
	
	if _ , isset := map_arr[key]; isset {
		delete(map_arr, key)
	}
	
	fmt.Println(map_arr)


}