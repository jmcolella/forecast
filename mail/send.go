package mail

import (
	"net/smtp"
	"os"
)

/*
Mail is a struct that holds the base credentials
for sending an email from a specific user.
*/
type Mail struct {
	Username string
	Password string
	Host     string
}

/*
Send uses the base credentials from the Mail struct
to send a given subject/body from the same user to itself
*/
func (m *Mail) Send(subject, body string) error {
	msg := []byte("Subject: " + subject + "\r\n" + body)
	to := []string{m.Username}

	return smtp.SendMail(m.Host+":587", *m.auth(), m.Username, to, msg)
}

func (m *Mail) auth() *smtp.Auth {
	auth := smtp.PlainAuth("", m.Username, m.Password, m.Host)

	return &auth
}

/*
NewMail creates a Mail struct with the base credentials setup
*/
func NewMail() *Mail {
	return &Mail{
		Username: "colella.john@gmail.com",
		Password: os.Getenv("GMAIL_PASSWORD"),
		Host:     "smtp.gmail.com",
	}
}
