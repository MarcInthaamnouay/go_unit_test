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
	StartDate string
	EndDate   string
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
	receiverObj.EXPECT().GetAge().Return(21)
	receiverObj.EXPECT().GetEmail().Return("angela.zhang@icbc.com")
	receiverObj.EXPECT().IsValid().Return(true)

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
	senderObj.EXPECT().GetEmail().Return("jiang.jianqing@icbc.com")

	// fmt.Println("person ", person)

	// Mock the isValid function by returning true
	isValid := mock_receiver.NewMockuser(receiverMockCtrl)
	isValid.EXPECT().IsValid().Return(true)

	// Mock a product
	pr := mock_product.NewMockP(productMockCtrl)
	// Mocking the create product constructor
	pr.EXPECT().CreateProduct("chopstick", 2, person).Return("chopstick", 2, person)
	pr.EXPECT().IsValid().Return(true)
	// Just like the product, we can't mock the product struct in go so we pass the original product struct

	// Mocking the Database
	dbObj := mock_db.NewMockdbInterface(dbMockCtrl)
	dbObj.EXPECT().SaveExchange().Return("", nil)

	// Mocking the SendMail
	mailObj := mock_mail.NewMockMail(mailMockCtrl)
	mailObj.EXPECT().SendMail("john@doe.com", "otherdoe@.com", "test").Return(true, nil)

	// Now we can test our Exchange class here....
	// Creating a new Exchange (we don't need a constructor as GO prefer composition over inheritance..)

	// Using our saveMock function

	// Building our fake struct

	sv := ExchangeMock{
		ReceiverM: receiverObj,
		ProductM:  pr,
		DbM:       dbObj,
		MailM:     mailObj,
		StartDate: "2016-Dec-24",
		EndDate:   "2016-Dec-25",
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
	dateNow, _errNow := exchange.GetFormatTime("now")
	formatStartDate, _errStart := exchange.GetFormatTime(e.StartDate)
	formatEndDate, _errEnd := exchange.GetFormatTime(e.EndDate)

	if _errStart != nil || _errEnd != nil || _errNow != nil {
		return false, errors.New("time is not valid")
	}

	if e.ReceiverM.IsValid() && e.ProductM.IsValid() {
		// convert the start date and the end date to a real date
		if formatStartDate.After(dateNow) && formatEndDate.After(dateNow) {
			if formatStartDate.Before(formatEndDate) && formatEndDate.After(formatStartDate) {
				if e.ReceiverM.GetAge() < 18 {
					e.MailM.SendMail(e.ReceiverM.GetEmail(), e.SenderM.GetEmail(), "Please take care of our new chopstick collection")
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
//	   We use the testify assert package	  //
////////////////////////////////////////////////

// @TODO rename TestMock into a getMock with the param that we want
// So we can return the struct the object that interest us or directly the struct
// Therefore we can test in the Test function
func TestDate(t *testing.T) {
	// Call the GetMock function
	localStruct := GetExchangeMock(t)

	localStruct.StartDate = "2016-Dec-24"
	localStruct.EndDate = "2016-Dec-25"

	r, _e := localStruct.saveMock()

	if _e != nil {
		assert.Fail(t, "an error happened while comparing 2 dates")
		fmt.Println(_e)
	}

	result := assert.Equal(t, r, true, "Time should be equal to true")

	if result {
		t.Logf("TestDate using normal date is a success")
	}
}

/*
 * TestRealObj
 *		Making the same mock test using the real package
 * @Params {t} testing
 */
func testRealObj(t *testing.T) {}
