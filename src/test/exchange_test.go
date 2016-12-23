/*
Package exchange_test implement the exchange process using mock Data

As we're not able to use the test command
without having a file named *_test.go
Therefore consider exchange_test.go as the exchange Class.

Author : Marc Intha-amnouay
Mail : marc.inthaamnouay@gmail.com
*/

package exchange_test

import (
	"errors"
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

// Define a structure of our Exchange
// Consider a struct as a constructor like in Java.
type ExchangeMock struct {
	ReceiverM *mock_receiver.Mockuser
	SenderM   *mock_receiver.Mockuser
	ProductM  *mock_product.MockP
	DbM       *mock_db.MockdbInterface
	MailM     *mock_mail.MockMail
	StartDate time.Time
	EndDate   time.Time
}

// GetExchangeMock is the function where we're creating a mock of our Package
// Here is the list of package that we're mocking
// ---- Receiver
// ---- Product
// ---- Database
// ---- Mail
// Note that Receiver and Sender is the same package
// As Go lang is pretty much new there're not a lot of mock library or framework..
// Therefore after reading articles & testing out some library I use the official gomock library + testify library for the assertion.
// GetExchangeMock(t *testing.T) ExchangeMock
func GetExchangeMock(t *testing.T) ExchangeMock {

	// Define the controller of our different package that we're mocking

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

	receiverObj := mock_receiver.NewMockuser(receiverMockCtrl)
	// Mock the create user constructor
	receiverObj.EXPECT().CreateUser("joe@doe.com", "doe", "john", 10).Return("john@gmail.com", "doe", "john", 30)
	receiverObj.EXPECT().GetEmail().Return("angela.zhang@icbc.com").AnyTimes()
	// Define the flow of execution of receiverObj.isValid
	// Call x5 times IsValid -> return true
	// Then xInfinite times IsValid -> return false
	gomock.InOrder(
		receiverObj.EXPECT().IsValid().Return(true).MaxTimes(5),
		receiverObj.EXPECT().IsValid().Return(false).AnyTimes(),
	)

	// Define the flow of execution of receiverObj.GetAge
	gomock.InOrder(
		receiverObj.EXPECT().GetAge().Return(21),
		receiverObj.EXPECT().GetAge().Return(16),
		receiverObj.EXPECT().GetAge().Return(17),
		receiverObj.EXPECT().GetAge().Return(19).AnyTimes(),
	)

	// Define the sender
	senderObj := mock_receiver.NewMockuser(senderMockCtrl)
	senderObj.EXPECT().GetEmail().Return("jiang.jianqing@icbc.com").AnyTimes()

	// Mock a product
	pr := mock_product.NewMockP(productMockCtrl)
	//mail, firstname, name, age := receiverObj.CreateUser("joe@doe.com", "doe", "john", 10)
	// Creating an empty receiver struct
	// We don't need anything as we're mocking it
	person := &receiver.UserStruct{}
	// Mocking the create product constructor
	pr.EXPECT().CreateProduct("chopstick", 2, person).Return("chopstick", 2, person)

	// Define the flow of how we're calling our pr.IsValid()
	gomock.InOrder(
		pr.EXPECT().IsValid().Return(true).MaxTimes(4),
		pr.EXPECT().IsValid().Return(false).AnyTimes(),
	)

	// Mocking the Database
	dbObj := mock_db.NewMockdbInterface(dbMockCtrl)
	// Set the flow
	gomock.InOrder(
		dbObj.EXPECT().SaveExchange().Return(true, nil).MaxTimes(2),
		dbObj.EXPECT().SaveExchange().Return(false, errors.New("can't save in the database")).MaxTimes(1),
	)

	// Mocking the SendMail
	mailObj := mock_mail.NewMockMail(mailMockCtrl)
	// Set the flow
	gomock.InOrder(
		mailObj.EXPECT().SendMail("angela.zhang@icbc.com", "jiang.jianqing@icbc.com", "Please take care of our new chopstick collection").Return(true, nil).MaxTimes(1),
		mailObj.EXPECT().SendMail("angela.zhang@icbc.com", "jiang.jianqing@icbc.com", "Please take care of our new chopstick collection").Return(false, errors.New("An error occured while sending the mail")).AnyTimes(),
	)

	// Now we can test our Exchange Here but we want to avoid the circular depedency..
	// Therefore we're creating our save function Here
	// It's call SaveMock

	// Building our ExchangeMock structure
	sv := ExchangeMock{
		ReceiverM: receiverObj,
		SenderM:   senderObj,
		ProductM:  pr,
		DbM:       dbObj,
		MailM:     mailObj,
		StartDate: time.Now().AddDate(0, 0, 1),
		EndDate:   time.Now().AddDate(0, 0, 2),
	}

	return sv
}

// SaveMock
// Save the exchange and handle error when it has
// SaveMock(bool, error)
// We're using composition therefore when calling the saveMock function
// We will have to do this way
// ---------------------------
// struct.saveMock()
func (e ExchangeMock) saveMock() (bool, error) {
	dateNow, _errNow := exchange.GetFormatTime(time.Now())
	formatStartDate, _errStart := exchange.GetFormatTime(e.StartDate)
	formatEndDate, _errEnd := exchange.GetFormatTime(e.EndDate)

	if _errStart != nil || _errEnd != nil || _errNow != nil {
		return false, errors.New("time is not valid")
	}

	var (
		userValid          bool = e.ReceiverM.IsValid()
		productValid       bool = e.ProductM.IsValid()
		age                int  = e.ReceiverM.GetAge()
		error              error
		startDateAfter     = formatStartDate.After(dateNow)
		endDateAfter       = formatEndDate.After(dateNow)
		startDateNow       = formatStartDate.Equal(dateNow)
		startDateBeforeEnd = formatStartDate.Before(formatEndDate)
		endDateAfterStart  = formatEndDate.After(formatStartDate)
	)

	if userValid && productValid {
		// convert the start date and the end date to a real date
		if startDateAfter && endDateAfter || startDateNow && endDateAfter {
			if startDateBeforeEnd && endDateAfterStart {
				if age < 18 {
					_, _err := e.MailM.SendMail(e.ReceiverM.GetEmail(), e.SenderM.GetEmail(), "Please take care of our new chopstick collection")
					if _err != nil {
						return false, errors.New("An error occured while sending the mail")
					}
				}

				// Send the data in the Database
				_, _errDB := e.DbM.SaveExchange()
				if _errDB != nil {
					return false, errors.New("An error occured while saving the exchange")
				}

				return true, nil
			} else {
				error = exchange.BuildError(": enter a valid value", startDateBeforeEnd, endDateAfterStart, "start date", "end date")

				return false, error
			}
		} else {
			error = exchange.BuildError(": enter a date after or today", startDateAfter, endDateAfter, "start date", "end date")
			return false, error
		}
	}

	error = exchange.BuildError("is not valid", userValid, productValid, "user", "product")

	return false, error
}

////////////////////////////////////////////////
//											  //
//				  UNIT TEST 				  //
//	   			  DATE TEST					  //
////////////////////////////////////////////////

// TestGoodOrderDate
// Testing the order of the date
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
	}

	assert.Equal(t, r, true, "Time should be equal to true")
}

// TestBadOrderDate
// Testing if the saveMock function handle that the start date is before today
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
	assert.EqualError(t, _e, "start date : enter a date after or today", "working")
}

// TestStartAfter
// Testing if saveMock handle that start date is after end date
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
	assert.EqualError(t, _e, "start date & end date : enter a valid value", "working")
}

// TestBeforeDate
// Test if the saveMock function handle if the date use are before today
func TestBeforeDate(t *testing.T) {
	beforeToday := GetExchangeMock(t)
	// Now set the date we're going to set the following date
	// Today date - 1 months
	// Today date - 2 months
	beforeToday.StartDate = time.Now().AddDate(0, -1, 0)
	beforeToday.EndDate = time.Now().AddDate(0, -2, 0)

	_, _e := beforeToday.saveMock()
	assert.EqualError(t, _e, "start date & end date : enter a date after or today")
}

// TestToday
// Test if the saveMock function handle a start date equal to today
func TestToday(t *testing.T) {
	today := GetExchangeMock(t)
	// Now set the date we're going to set the following date
	// Today date - 1 months
	// Today date - 2 months
	today.StartDate = time.Now().AddDate(0, 0, 0)
	today.EndDate = time.Now().AddDate(0, 0, 1)

	res, _e := today.saveMock()

	if _e != nil {
		assert.Fail(t, "today failed test.")
	}
	assert.Equal(t, res, true)
}

////////////////////////////////////////////////
//											  //
//				  UNIT TEST 				  //
// 			AGE AND MAIL SENDING 			  //
////////////////////////////////////////////////

// TestAgeForMail
// --- 3 unit test in there
// - 1 : asserting if the exchange has been sent when the user is 21
// - 2 : asserting if we're sending a mail
// - 3 : asserting if we can't send a mail
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

////////////////////////////////////////////////
//											  //
//				  UNIT TEST 				  //
// 			USER & PRODUCT VALIDATION 	      //
////////////////////////////////////////////////

// TestIsValid
// 3 unit test for the price of one function
// 1 -- testing if the user and the product is valid
// 2 -- test if the saveMock handle the product isn't valid
// 3 -- test if the saveMock handle that both product and receiver isn't valid
func TestIsValid(t *testing.T) {
	assert := assert.New(t)
	s := GetExchangeMock(t)

	// User and Product is equal to true
	usAndProd, _err := s.saveMock()

	if _err != nil {
		assert.Fail("testing both user and product valid has failed")
	}

	// Asserting that our product & user is valid
	assert.Equal(true, usAndProd, "test user and product validity")

	// Executing multiple times our function
	// In order to go the the case where our product isn't valid
	for i := 0; i < 3; i++ {
		s.saveMock()
	}

	_, _e := s.saveMock()
	// Asserthing the fact that our product is not valid
	assert.EqualError(_e, "product is not valid", "test successfull user is not valid")

	// Asserting the fact that both product and user isn't valid
	_, _eProduct := s.saveMock()
	assert.EqualError(_eProduct, "user & product is not valid")
}

////////////////////////////////////////////////
//											  //
//				  UNIT TEST 				  //
// 				SAVE EXCHANGE 				  //
////////////////////////////////////////////////

// TestErrorExchange
// Test if saveMock handle the fact that there's an error when saving an exchange
func TestErrorExchange(t *testing.T) {
	// Get our structure
	s := GetExchangeMock(t)

	// passing directly to the 3rd call..
	for i := 0; i < 3; i++ {
		s.saveMock()
	}

	_, _e := s.saveMock()

	assert.EqualError(t, _e, "An error occured while saving the exchange")
}
