package user_test

import (
	"github.com/fajarAnd/learn-go-with-test/gomock/mocks"
	"github.com/fajarAnd/learn-go-with-test/gomock/user"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestUse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &user.User{Doer:mockDoer}

	// Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return nil from the mocked call.
	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(nil).Times(1)

	testUser.Use()
}