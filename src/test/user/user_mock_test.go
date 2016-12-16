package mockUser

import (
	"start"
	"testing"
)

type UserTestStruct struct {
	user *start.UserStruct
	t    *testing.T
}

func (u UserTestStruct) checkUserExpectation() {
	// Check the integrety of the User

	if u.user == nil {
		u.t.Errorf("User creation has fail - MOCK FAIL")
	} else {
		u.t.Logf("User creation has been successfull")
	}

	// Checking if the user is valid

	userValid := u.user.IsValid()

	if userValid != true {
		u.t.Errorf("User is not valid - MOCK FAIL")
	} else {
		u.t.Logf("User verification has been successfull")
	}

	// Checking if we can set the age of Jenny
	// store the value of the age (not the reference)
	// Note that for more consistency we can declare our variable like this
	// var previousAge int = u.user.Age
	var previousAge = u.user.Age

	// Changing the age of the person
	u.user.SetAge(25)

	// Retrieve the age and test if the age has been changed
	if u.user.GetAge() != previousAge {
		u.t.Logf("User age has been changed - MOCK SUCCESS - SetAge() & GetAge()")
	} else {
		u.t.Errorf("User age has not changed - MOCK FAILED")
	}

}

/*
 * TestUser
 *      Test the creation of the user, the validation of it and set some params to it
 * @Params {test} testing
 *      Note that in GO there're framework which can help us to do mocks.
 *      However it's not that recommend as many dev find it too complicate or doing some sort of black magic.
 *      Therefore it's recommand to mock by yourself.
 * In order to mock this function I followed the articles down below..
 *      https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742#.ppovblxsk
 *      http://www.philosophicalhacker.com/2016/01/13/should-we-use-mocking-libraries-for-go-testing/
 */
func TestUser(test *testing.T) {

	jenny := UserTestStruct{
		user: &start.UserStruct{
			Email:     "jenny@gmail.com",
			Firstname: "Jenny",
			Name:      "Dembele",
			Age:       20,
		},
		t: test,
	}

	jenny.checkUserExpectation()

}
