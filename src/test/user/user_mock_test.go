package mockUser

import (
	"errors"
	"fmt"
	"modules/db_mock"
	"modules/mock"
	"modules/receiver_mock"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestMock(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockObj := mock_mail.NewMockMail(mockCtrl)
	mockObj.EXPECT().SendMail().Return(errors.New("Mail has not been implemented"))

	u := mockObj
	test := u.SendMail()
	fmt.Println(test)
	// pass mockObj to a real object and play with it.
}

func TestReceiver(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	receiverMockObj := mock_receiver.NewMockuser(mockCtrl)
	receiverMockObj.EXPECT().IsValid().Return(true)

	testIsValid := receiverMockObj
	isTrue := testIsValid.IsValid()

	fmt.Println(isTrue)
}

func TestDB(t *testing.T) {
	dbmockCtrl := gomock.NewController(t)
	defer dbmockCtrl.Finish()

	dbMockObj := mock_db.NewMockdbInterface(dbmockCtrl)
	dbMockObj.EXPECT().SaveExchange()

	dbMockObj.SaveExchange()

}

func TestExchange(t *testing.T) {

}
