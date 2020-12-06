package user

import (
	"github.com/fajarAnd/learn-go-with-test/gomock/doer"
)

type User struct {
	Doer doer.Doer
}

func (u *User) Use() error {
	return u.Doer.DoSomething(123, "Hello GoMock")
}
