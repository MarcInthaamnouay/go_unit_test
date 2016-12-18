package mail

import "errors"

type Mail interface {
	SendMail() error
}

type MailContructor struct {
	sender string
	to     string
	object string
}

func (m MailContructor) SendMail() error {
	// Throwing an error without stopic the go program
	err := errors.New("Mail has not been implemented")

	return err
}
