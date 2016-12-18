package mockUser

import (
	"fmt"
	"modules/mock"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestMock(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockObj := mock_mail.NewMockMail(mockCtrl)
	mockObj.EXPECT().SendMail()

	u := mockObj

	test := u.SendMail()

	fmt.Println(test)
	// pass mockObj to a real object and play with it.
}
