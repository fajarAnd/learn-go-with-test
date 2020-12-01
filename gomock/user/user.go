package user

import (
	"../doer"
)

type User struct {
	Doer doer.Doer
}

func (u *User) use() error {
	return u.Doer.DoSomething(123, "Hello Gomock")
}
