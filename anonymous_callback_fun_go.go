package main 

import  ("fmt")

func main() {

   cfunc  := func (num  []int) int {
   			for i, v := range num {
   				for j, val := range num {
   					//fmt.Println(i, " ",v, " and ", val)
   					if v > val {
   						num[i] = val	
   						num[j] = v
   						v = val
   						//fmt.Println(i, " ", j, " ",num)
   					}
   				}
   			}
   			//fmt.Println(num)
   			return num[0]
   }

   fmt.Println(get_max_number(cfunc,[]int{-1,-13,-14,0,-6,-7,-8,-9}))
}


func get_max_number(f func(x []int)int, ii []int) int {
	return f(ii)
}