package userApp


//var userInfoData user_info

import (
	"fmt"
)
func NewUser()  *user_info {
	user := new(user_info)

	user.userFristName = ""
	user.userLastName  = ""
	user.userEmailID.emailID  = make(map[precednce]string, 3)
	user.userEmailID.isvalidated = false
	user.userMobileNumber.mnumber = make(map[precednce] rune, 3)
	user.userMobileNumber.isvalidated = false
	user.userGender.male  = false
	user.userGender.female = false
	user.userGender.other = false
	user.passWord = ""
	user.userName  = ""

	return user
}

func (user_info u) Create_account() {
	fmt.print("You Want to Create new account [y/n]: ")
	var c string 
	fmt.scanf("%s",&c)
	if c == "y" {
		fmt.println("Please feel below given info [* is mandetory] ")
	}
} 