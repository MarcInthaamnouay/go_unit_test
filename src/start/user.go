package main

/*
 * @author       Intha-amnouay Marc
 * @mail         marc.inthaamnouay@gmail.com
 * @goversion    1.7
 * @run once     go run src/start/user.go
 * @run routine  $GOPATH/bin/start
 */

import "net/mail"
import "fmt"

type UserStruct struct {
	Email     string
	Firstname string
	Name      string
	Age       int
}

type User interface {
	isValid() bool
	setEmail(string)
	setFirstname(string)
	setLastname(string)
	setAge(int)
	getEmail() string
	getFirstname() string
	getLastname() string
	getAge() int
}

func createUser(email string, firstname string, name string, age int) *UserStruct {
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
func (user *UserStruct) isValid() bool {
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
func (user *UserStruct) setEmail(mail string) {
	fmt.Println(mail)
	user.Email = mail
	fmt.Println(user)
}

/*
 * SetFirstname
 *      Set the firstname of the user
 * @Params {name} string
 */
func (user *UserStruct) setFirstname(firstname string) {
	user.Firstname = firstname
}

/*
 * SetLastName
 *      Set the lastname of the user
 * @Params {lastname} string
 */
func (user *UserStruct) setLastName(lastname string) {
	user.Name = lastname
}

/*
 * SetAge
 *      Set the age of the user
 * @Params {age} Int
 */
func (user *UserStruct) setAge(age int) {
	user.Age = age
}

/*
 * GetEmail
 *      Get the email of the user
 * @Params {user} UserStruct
 * @Return {mail} string
 */
func (user *UserStruct) getEmail() string {
	return user.Email
}

/*
 * GetFirstname
 *      Get the firstname of the user
 * @Params {user} UserStruct
 * @Return {firstname} string
 */
func (user *UserStruct) getFirstname() string {
	return user.Firstname
}

/*
 * GetLastName
 *      Get the last name of the user
 * @Params {user} UserStruct
 * @Return {name} string
 */
func (user *UserStruct) getLastName() string {
	return user.Name
}

/*
 * GetAge
 *      Get the age of the user
 * @Prams {user} UserStruct
 * @Return {age} int
 */
func (user *UserStruct) getAge() int {
	return user.Age
}

/*
 * Main <Void>
 *     Main function of the program
 */
func main() {
	user := createUser("marc@gmail.com", "marc", "intha", 18)
	fmt.Println(user)

	// Check if the user is valid
	isValid := user.isValid()

	fmt.Println("est-ce que je suis valide ? %n", isValid)

	// try to set the value

	fmt.Println(user.getAge())
	user.setAge(21)
	fmt.Println(user)
}
