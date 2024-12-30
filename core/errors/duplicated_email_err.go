package errors

import "fmt"

type DuplicatedEmailErr struct {
	email string
}

func NewDuplicatedEmailErr(email string) error {
	return DuplicatedEmailErr{email: email}
}

func (d DuplicatedEmailErr) Error() string {
	return fmt.Sprintf("Email %v has aready been taken.", d.email)
}
