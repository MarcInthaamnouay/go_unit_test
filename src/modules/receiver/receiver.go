package receiver

/*
 * @author       Intha-amnouay Marc
 * @mail         marc.inthaamnouay@gmail.com
 * @goversion    1.7
 * @run once     go run src/start/user.go
 * @run routine  $GOPATH/bin/start
 */

import (
	"fmt"
	"net/mail"
)

type UserStruct struct {
	Email     string
	Firstname string
	Name      string
	Age       int
}

type user interface {
	CreateUser() *UserStruct
	IsValid() bool
	SetEmail(string)
	GetEmail() string
	SetFirstName(string)
	GetFirstname() string
	SetName(string)
	GetName() string
	SetAge(int)
	GetAge() int
}

/*
 * CreateUser
 *		Method which create a new user it mimick the Java constructor ...
 *		Note that in GO it's recommend to use directly the strucutre to create an object..
 *	@Params {email} string
 *	@Params {Firstname} string
 *	@Params {name} string
 *	@Params {age} int
 */
func CreateUser(email string, firstname string, name string, age int) *UserStruct {
	user := &UserStruct{
		Email:     email,
		Firstname: firstname,
		Name:      name,
		Age:       age,
	}

	return user
}

/* ----------------------------------- USER INTERFACE METHOD ----------------------------------- */

/*
 * IsValid
 *     Check if the user is Valid
 * @Params {User} UserStruct
 * @Return {Bool} Boolean
 */
func (user *UserStruct) IsValid() bool {
	if len(user.Firstname) > 0 && len(user.Name) > 0 && user.Age > 13 {
		// Checking the mail adress
		checkAdd, err := mail.ParseAddress(user.Email)

		if err != nil {
			// Throw an error
			return false
		}

		fmt.Println(checkAdd)
	}

	return true
}

/*
 * SetMail
 *      Set the mail of the user
 * @Params {mail} string
 */
func (user *UserStruct) SetEmail(mail string) {
	fmt.Println(mail)
	user.Email = mail
	fmt.Println(user)
}

/*
 * SetFirstname
 *      Set the firstname of the user
 * @Params {name} string
 */
func (user *UserStruct) SetFirstName(firstname string) {
	user.Firstname = firstname
}

/*
 * SetLastName
 *      Set the lastname of the user
 * @Params {lastname} string
 */
func (user *UserStruct) SetName(lastname string) {
	user.Name = lastname
}

/*
 * SetAge
 *      Set the age of the user
 * @Params {age} Int
 */
func (user *UserStruct) SetAge(age int) {
	user.Age = age
}

/*
 * GetEmail
 *      Get the email of the user
 * @Params {user} UserStruct
 * @Return {mail} string
 */
func (user *UserStruct) GetEmail() string {
	return user.Email
}

/*
 * GetFirstname
 *      Get the firstname of the user
 * @Params {user} UserStruct
 * @Return {firstname} string
 */
func (user *UserStruct) GetFirstName() string {
	return user.Firstname
}

/*
 * GetName
 *      Get the last name of the user
 * @Params {user} UserStruct
 * @Return {name} string
 */
func (user *UserStruct) GetName() string {
	return user.Name
}

/*
 * GetAge
 *      Get the age of the user
 * @Prams {user} UserStruct
 * @Return {age} int
 */
func (user *UserStruct) GetAge() int {
	return user.Age
}
