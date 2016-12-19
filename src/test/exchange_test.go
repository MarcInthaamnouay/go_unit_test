package exchange_test

import (
	"fmt"
	"modules/product"
	"modules/product_mock"
	"modules/receiver"
	"modules/receiver_mock"
	"testing"

	"github.com/golang/mock/gomock"
)

/*
 * TestExchange
 *      TestExchange is the whole function which go will use to test our Exchange Package (there's no Class in GO)
 * @Params {t} *testing
 *
 */
func TestExchange(t *testing.T) {

	// variable of the user
	var mail, name, firstname string
	var age int

	// variable of the product
	var name_product string
	var status int
	var p *receiver.UserStruct
	// Define the controller of our different package
	// Receiver
	receiverMockCtrl := gomock.NewController(t)
	// Product
	productMockCtrl := gomock.NewController(t)
	// Database
	//dbMockCtrl := gomock.NewController(t)
	// MailSender
	//mailMockCtrl := gomock.NewController(t)

	/*
	   In this test we're going to test the behavior of our save function by giving the normal parameters
	*/

	// First we're creating a new user using our Mock receiverMockCtrl
	receiverObj := mock_receiver.NewMockuser(receiverMockCtrl)
	// Mock the create user constructor
	receiverObj.EXPECT().CreateUser("joe@doe.com", "doe", "john", 100).Return("john@gmail.com", "doe", "john", 300)
	receiverObj.EXPECT().GetAge().Return(true)
	mail, name, firstname, age = receiverObj.CreateUser("joe@doe.com", "doe", "john", 100)

	// Now in GO you can't create an Empty Object without defining a strict structure.
	// Therefore we use the original receiver structure
	person := &receiver.UserStruct{
		Email:     mail,
		Firstname: firstname,
		Name:      name,
		Age:       age,
	}

	fmt.Println("person ", person)

	// Mock the isValid function by returning true
	isValid := mock_receiver.NewMockuser(receiverMockCtrl)
	isValid.EXPECT().IsValid().Return(true)

	// Mock a product
	pr := mock_product.NewMockP(productMockCtrl)
	// Mocking the create product constructor
	pr.EXPECT().CreateProduct("chopstick", 2, person).Return("chopstick", 2, person)
	pr.EXPECT().IsValid().Return(true)
	// Just like the product, we can't mock the product struct in go so we pass the original product struct
	name_product, status, p = pr.CreateProduct("chopstick", 2, person)

	// create a product... using the mock value
	chopstick := product.Product{
		Name:   name_product,
		Status: status,
		Owner:  p,
	}

	fmt.Println("product ", chopstick)
	// @TODO return true to a user and a product

}
