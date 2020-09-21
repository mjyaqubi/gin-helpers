package email

import (
	"gopkg.in/gomail.v2"
)

// Email type struct
type Email struct {
	Options *Options
}

// Options - options struct
type Options struct {
	Host     string
	Port     int
	Username string
	Password string
}

// New Email
func New(options *Options) *Email {
	return &Email{
		Options: options,
	}
}

// Send email
func (helper *Email) Send(to string, subject string, message string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", helper.Options.Username)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", message)

	d := gomail.NewDialer(
		helper.Options.Host,
		helper.Options.Port,
		helper.Options.Username,
		helper.Options.Password,
	)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
