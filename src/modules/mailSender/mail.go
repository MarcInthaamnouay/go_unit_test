package mail

import "errors"

type Mail interface {
	SendMail(s string, t string, o string) (bool, error)
}

func SendMail(sender string, to string, object string) (bool, error) {
	// Throwing an error without stopic the go program
	err := errors.New("Mail has not been implemented")

	return false, err
}
