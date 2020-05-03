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
GetRecipients returns a slice of email addresses
to send forecast data to
*/
func (m *Mail) GetRecipients() []string {
	return []string{os.Getenv("JOHN_EMAIL"), os.Getenv("MAGGIE_EMAIL")}
}

/*
Send uses the base credentials from the Mail struct
to send a given subject/body from the same user to itself
*/
func (m *Mail) Send(subject, body string) error {
	msg := []byte("Subject: " + subject + "\r\n" + body)
	to := m.GetRecipients()

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
		Username: os.Getenv("JOHN_EMAIL"),
		Password: os.Getenv("GMAIL_PASSWORD"),
		Host:     "smtp.gmail.com",
	}
}
