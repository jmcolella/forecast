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
Recipient is a struct that holds information
for sending weather emails to the specified people
*/
type Recipient struct {
	Email  string
	CityID string
}

/*
GetRecipients returns a slice of email addresses
to send forecast data to
*/
func (m *Mail) GetRecipients() []*Recipient {
	var recipients []*Recipient

	recipients = append(recipients, &Recipient{Email: os.Getenv("JOHN_EMAIL"), CityID: "5110302"})
	recipients = append(recipients, &Recipient{Email: os.Getenv("MAGGIE_EMAIL"), CityID: "5133268"})

	return recipients
}

/*
Send uses the base credentials from the Mail struct
to send a given subject/body to the specified recipients
*/
func (m *Mail) Send(recipients []string, subject, body string) error {
	msg := []byte("Subject: " + subject + "\r\n" + body)

	return smtp.SendMail(m.Host+":587", *m.auth(), m.Username, recipients, msg)
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
