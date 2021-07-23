package main 

import ( "fmt" )

var allUserInfo map[string]user_info

type user_info struct {
	fristName string 
	lastNmae string
	phoneNumbr rune
	email string
	userName string
	password string
}

func addNewUser(user *user_info) bool {
	if user == nil || allUserInfo == nil{
		fmt.Println("nill fund ", user, " | ", allUserInfo)
		return false
	}
	allUserInfo[user.userName] = *user
	fmt.Println("Your username is  : ", allUserInfo[user.userName].userName)

	return true
}

func editUserInfo(userName string) bool{
	 var user_intput string = "user Name"
	if v, isset := allUserInfo[userName]; isset {
		fmt.Println(v)
		fmt.Println("What you want to edit : ")

		//fmt.Scanf("%s", &user_intput)

		switch user_intput {
		case "user Name":
				un := ""
				cun := ""
				fmt.Print("Enter yout New user name : ")
				fmt.Scanf("%s", &un)
				fmt.Print("Re-Enter your user name : ")
				fmt.Scanf("%s", &cun)
				fmt.Print("Processing......")
				if un == cun {
					return true
				} else {
					fmt.Print("somthing wrong happen")
				}
		case "password" :
		case "email id" :
		case "phone no" :
		default :
		}

	}
	return false
}


func main () {
	userName := ""
	allUserInfo = make (map[string]user_info, 10)
	newUser := user_info{"Bhushan", "Patil", 12345 , "xyz@gmail.com", "Bhushna@123", "password",}
	addNewUser(&newUser)
	fmt.Print("Enter your user name for edit info : ")
	fmt.Scanf("%s", &userName)
	editUserInfo(userName)
}