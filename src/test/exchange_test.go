package exchange_test

import (
	"errors"
	"fmt"
	"modules/db_mock"
	"modules/mail_mock"
	"modules/product_mock"
	"modules/receiver"
	"modules/receiver_mock"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
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
 * TestExchange
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
func TestExchange(t *testing.T) {

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

	validTest, _errValid := sv.saveMock()

	if _errValid != nil {
		fmt.Println("error during save mock test")
	}

	fmt.Println("test valid ! mock save", validTest)
}

/*
 * saveMock
 *		this function is mimicking the exchange.mock function in the original
 *		Exchange package. The difference is that it used the Mock struct instead
 *		of the original one.
 * @Private
 * @Compositon parameters {ExchangeMock} structure
 * @Return {bool} boolean
 * @Return {error} errors
 */
func (e ExchangeMock) saveMock() (bool, error) {
	const shortForm = "2006-Jan-02"
	dateNowStr := time.Now().Format(shortForm)
	formatStartDate, _errStart := time.Parse(shortForm, e.StartDate)
	formatEndDate, _errEnd := time.Parse(shortForm, e.EndDate)
	dateNow, _errNow := time.Parse(shortForm, dateNowStr)

	if _errStart != nil || _errEnd != nil || _errNow != nil {
		return false, errors.New("time is not valid")
	}

	fmt.Println("before start date ? ", formatStartDate.Before(formatEndDate))
	fmt.Println("after start date ? ", formatEndDate.After(formatStartDate))

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

/*
 * TestRealObj
 *		Making the same mock test using the real package
 * @Params {t} testing
 */
func testRealObj(t *testing.T) {}
