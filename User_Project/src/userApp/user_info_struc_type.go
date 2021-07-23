
package userApp

const (
	NUMBEROFUSER = 10
	NUMBEROFRETRY = 3 // Number of retries for password 
)

type precednce struct {
	primery  int 
	seconder int 
}

type userEmailID struct {
	emailID map[precednce]string
	isvalidated bool
}

type userMobileNumber struct {
	mnumber map[precednce]rune
	isvalidated bool
}

type userGender struct {
	male bool
	female bool
	other bool
}

type user_info struct {
	userFristName string
	userMiddlName string  
	userLastName string 
	userEmailID
	userMobileNumber
	userGender
	userName string 
	passWord string
}

