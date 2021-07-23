//Create slice to hold all state in india, and print it.

// {"Andhra Pradesh" , "Arunachal Pradesh" , "Assam" , "Bihar" , "Chhattisgarh" , "Goa" , "Gujarat" , 
// "Haryana" , "Himachal Pradesh" , "Jammu and Kashmir" , "Jharkhand" , "Karnataka" , "Kerala" , 
// "Madhya Pradesh" , "Maharashtra" , "Manipur" , "Meghalaya" , "Mizoram" , "Nagaland" , "Odisha" , 
// "Punjab" , "Rajasthan" , "Sikkim" , "Tamil Nadu" , "Telangana" , "Tripura" , "Uttarakhand" , "Uttar Pradesh" ,
// "West Bengal" , "Andaman and Nicobar Islands" , "Chandigarh" , "Dadra and Nagar Haveli" , "Daman and Diu" ,
// "Delhi" , "Lakshadweep" , "Puducherry" }


package main 

import ( "fmt" )

func main () {
	
	//create slice using "make" func
	// func make(s []T, len, capacity) []T
	var slice_var []string
	slice_var = make([]string, 36, 50)
	
	slice_var = []string{"Andhra Pradesh" , "Arunachal Pradesh" , "Assam" , "Bihar" , "Chhattisgarh" , "Goa" , "Gujarat" , 
						"Haryana" , "Himachal Pradesh" , "Jammu and Kashmir" , "Jharkhand" , "Karnataka" , "Kerala" , 
						"Madhya Pradesh" , "Maharashtra" , "Manipur" , "Meghalaya" , "Mizoram" , "Nagaland" , "Odisha" , 
						"Punjab" , "Rajasthan" , "Sikkim" , "Tamil Nadu" , "Telangana" , "Tripura" , "Uttarakhand" , "Uttar Pradesh" ,
						"West Bengal" , "Andaman and Nicobar Islands" , "Chandigarh" , "Dadra and Nagar Haveli" , "Daman and Diu" ,
						"Delhi" , "Lakshadweep" , "Puducherry" }
	
	fmt.Println("Index\t\tState")
	
	for i := 0; i < len(slice_var); i++ {
			fmt.Println(i , slice_var[i])
	}

}