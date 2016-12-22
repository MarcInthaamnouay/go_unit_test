package exchange_test

import (
	"errors"
	"fmt"
	"modules/db_mock"
	"modules/exchange"
	"modules/mail_mock"
	"modules/product_mock"
	"modules/receiver"
	"modules/receiver_mock"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// Define the object that use the saveMock function
type ExchangeMock struct {
	ReceiverM *mock_receiver.Mockuser
	SenderM   *mock_receiver.Mockuser
	ProductM  *mock_product.MockP
	DbM       *mock_db.MockdbInterface
	MailM     *mock_mail.MockMail
	StartDate time.Time
	EndDate   time.Time
}

/*
 * GetExchangeMock
 *      TestExchange is the whole function which go will use to test our Exchange Package (there's no Class in GO)
 *      As we need to use the testing package therefore we can't mock the Exchange package in the exchange package but only in the *_test.go files.
 *      Using the go test -v *_test.go
 *
 *      Mock based on https://godoc.org/github.com/golang/mock/gomock
 *      Test based on https://golang.org/pkg/testing/
 *
 * @Params {t} *testing
 *
 */
func GetExchangeMock(t *testing.T) ExchangeMock {

	// Define the controller of our different package
	// Receiver
	receiverMockCtrl := gomock.NewController(t)
	// Sender
	senderMockCtrl := gomock.NewController(t)
	// Product
	productMockCtrl := gomock.NewController(t)
	// Database
	dbMockCtrl := gomock.NewController(t)
	// MailSender
	mailMockCtrl := gomock.NewController(t)

	/*
	   In this test we're going to test the behavior of our save function by giving the normal parameters
	*/

	// First we're creating a new user using our Mock receiverMockCtrl
	receiverObj := mock_receiver.NewMockuser(receiverMockCtrl)
	// Mock the create user constructor
	receiverObj.EXPECT().CreateUser("joe@doe.com", "doe", "john", 10).Return("john@gmail.com", "doe", "john", 30)
	receiverObj.EXPECT().GetEmail().Return("angela.zhang@icbc.com").AnyTimes()
	receiverObj.EXPECT().IsValid().Return(true).AnyTimes()

	// Chaining the call of GetAge
	gomock.InOrder(
		receiverObj.EXPECT().GetAge().Return(21),
		receiverObj.EXPECT().GetAge().Return(16),
		receiverObj.EXPECT().GetAge().Return(17),
	)

	mail, firstname, name, age := receiverObj.CreateUser("joe@doe.com", "doe", "john", 10)

	// Now in GO you can't create an Empty Object without defining a strict structure.
	// Therefore we use the original receiver structure
	person := &receiver.UserStruct{
		Email:     mail,
		Firstname: firstname,
		Name:      name,
		Age:       age,
	}

	senderObj := mock_receiver.NewMockuser(senderMockCtrl)
	senderObj.EXPECT().GetEmail().Return("jiang.jianqing@icbc.com").AnyTimes()

	// fmt.Println("person ", person)

	// Mock a product
	pr := mock_product.NewMockP(productMockCtrl)
	// Mocking the create product constructor
	pr.EXPECT().CreateProduct("chopstick", 2, person).Return("chopstick", 2, person)
	pr.EXPECT().IsValid().Return(true).MaxTimes(3)
	//@TODO to see if it works..
	pr.EXPECT().IsValid().Return(false).MaxTimes(3)

	// Just like the product, we can't mock the product struct in go so we pass the original product struct

	// Mocking the Database
	dbObj := mock_db.NewMockdbInterface(dbMockCtrl)
	dbObj.EXPECT().SaveExchange().Return(true, nil).MaxTimes(3)

	// Mocking the SendMail
	mailObj := mock_mail.NewMockMail(mailMockCtrl)
	gomock.InOrder(
		mailObj.EXPECT().SendMail("angela.zhang@icbc.com", "jiang.jianqing@icbc.com", "Please take care of our new chopstick collection").Return(true, nil),
		mailObj.EXPECT().SendMail("angela.zhang@icbc.com", "jiang.jianqing@icbc.com", "Please take care of our new chopstick collection").Return(false, errors.New("An error occured while sending the mail")),
	)

	// Now we can test our Exchange class here....
	// Creating a new Exchange (we don't need a constructor as GO prefer composition over inheritance..)

	// Using our saveMock function

	// Building our fake struct

	sv := ExchangeMock{
		ReceiverM: receiverObj,
		SenderM:   senderObj,
		ProductM:  pr,
		DbM:       dbObj,
		MailM:     mailObj,
		StartDate: time.Now().AddDate(0, 0, 1),
		EndDate:   time.Now().AddDate(0, 0, 2),
	}

	// Making some test using our mock...

	// validTest, _errValid := sv.saveMock()

	// if _errValid != nil {
	// 	fmt.Println("error during save mock test")
	// }

	// fmt.Println("test valid ! mock save", validTest)

	return sv
}

/*
 * saveMock
 *		this function is mimicking the exchange.mock function in the original
 *		Exchange package. The difference is that it used the Mock struct instead
 *		of the original one.
 *		We did not put it in the Original Exchange module as it would create a circular depedency
 * @Private
 * @Compositon parameters {ExchangeMock} structure
 * @Return {bool} boolean
 * @Return {error} errors
 */
func (e ExchangeMock) saveMock() (bool, error) {
	dateNow, _errNow := exchange.GetFormatTime(time.Now())
	formatStartDate, _errStart := exchange.GetFormatTime(e.StartDate)
	formatEndDate, _errEnd := exchange.GetFormatTime(e.EndDate)

	if _errStart != nil || _errEnd != nil || _errNow != nil {
		return false, errors.New("time is not valid")
	}

	if e.ReceiverM.IsValid() && e.ProductM.IsValid() {
		// convert the start date and the end date to a real date
		if formatStartDate.After(dateNow) && formatEndDate.After(dateNow) {
			if formatStartDate.Before(formatEndDate) && formatEndDate.After(formatStartDate) {
				age := e.ReceiverM.GetAge()
				fmt.Println(age)

				if age < 18 {
					_, _err := e.MailM.SendMail(e.ReceiverM.GetEmail(), e.SenderM.GetEmail(), "Please take care of our new chopstick collection")

					if _err != nil {
						return false, errors.New("An error occured while sending the mail")
					}
				}

				// Send the data in the Database
				res, _errDB := e.DbM.SaveExchange()

				if _errDB != nil {
					return false, errors.New("An error occured while saving the exchange")
				}

				fmt.Println("exchange saved ! ", res)

				return true, nil
			} else {
				return false, errors.New("Please check the date")
			}
		} else {
			return false, errors.New("Please use a date that's after today")
		}
	}

	return false, errors.New("Product or User isn't valid")
}

////////////////////////////////////////////////
//											  //
//				  UNIT TEST 				  //
//	   			  DATE TEST					  //
////////////////////////////////////////////////

// @TODO rename TestMock into a getMock with the param that we want
// So we can return the struct the object that interest us or directly the struct
// Therefore we can test in the Test function
func TestGoodOrderDate(t *testing.T) {
	// Call the GetMock function
	localStruct := GetExchangeMock(t)
	// We set the date of today + 1
	localStruct.StartDate = time.Now().AddDate(0, 0, 1)
	// We set the date of today + 2
	localStruct.EndDate = time.Now().AddDate(0, 0, 2)

	r, _e := localStruct.saveMock()

	if _e != nil {
		assert.Fail(t, "Error catch, an error happened while testing the date")
		fmt.Println(_e)
	}

	result := assert.Equal(t, r, true, "Time should be equal to true")

	if result {
		t.Logf("TestDate using normal date is a success")
	}
}

func TestBadOrderDate(t *testing.T) {
	// First get a copy of our structure
	badDateStructure := GetExchangeMock(t)
	// Now set the date we're going to set the following date
	// Today date - 1
	// Today date + 1
	badDateStructure.StartDate = time.Now().AddDate(0, 0, -1)
	badDateStructure.EndDate = time.Now().AddDate(0, 0, 1)

	// Call our saveMock function
	res, _e := badDateStructure.saveMock()
	// Testing the exchange class using EqualError
	// If the saveMock function return true then the test has failed
	if res != false {
		assert.Fail(t, "Test failed TestBadOrderDate")
	}
	// assert the error and chekc if the string is the one that we input
	result := assert.EqualError(t, _e, "Please use a date that's after today", "working")

	if result {
		t.Logf("Test TestBadOrderDate pass")
	}
}

func TestStartAfter(t *testing.T) {
	// First get a copy of our structure
	badStartDateStructure := GetExchangeMock(t)
	// Now set the date we're going to set the following date
	// Today date + 3
	// Today date + 1
	badStartDateStructure.StartDate = time.Now().AddDate(0, 0, 3)
	badStartDateStructure.EndDate = time.Now().AddDate(0, 0, 1)

	// Call our saveMock function
	res, _e := badStartDateStructure.saveMock()
	// Testing the exchange class using EqualError
	// If the saveMock function return true then the test has failed
	if res != false {
		assert.Fail(t, "Test failed TestStartAfter")
	}

	// assert the error and chekc if the string is the one that we input
	result := assert.EqualError(t, _e, "Please check the date", "working")

	if result {
		t.Logf("Test TestStartAfter pass")
	}
}

////////////////////////////////////////////////
//											  //
//				  UNIT TEST 				  //
// 			AGE AND MAIL SENDING 			  //
////////////////////////////////////////////////

// 2 test in one function
func TestAgeForMail(t *testing.T) {
	// Copy our struct
	assert := assert.New(t)
	ageTestStruct := GetExchangeMock(t)

	// Testing the first call with an age of 21
	firstCall, _fErr := ageTestStruct.saveMock()
	if _fErr != nil {
		assert.Fail("error while trying to save the exchange with a user of 21")
	}

	assert.Equal(firstCall, true, "call is ok")

	// Testing the first call with an age of 16 it should send a mail
	secondCall, _sErr := ageTestStruct.saveMock()
	if _sErr != nil {
		assert.Fail("error happened while trying to send a mail, user of 18")
	}

	assert.Equal(secondCall, true, "second call is ok..")

	// Testing the first call with an age of 16 it should send a mail
	thirdCall, _tErr := ageTestStruct.saveMock()
	if thirdCall != false {
		assert.Fail("the mail should have not been send")
	}

	assert.EqualError(_tErr, "An error occured while sending the mail")
}

/*
 * TestRealObj
 *		Making the same mock test using the real package
 * @Params {t} testing
 */
func testRealObj(t *testing.T) {}
