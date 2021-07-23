package main

import (
	"fmt"
	"src/userApp"
)


func main() {
	user := userApp.NewUser()
	fmt.println(user)
	user.Create_account()
}