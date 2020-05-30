package mail

import "os"

/*
Recipient is a struct that holds information
for sending weather emails to the specified people
*/
type Recipient struct {
	Email    string
	Location RecipientLocation
}

type RecipientLocation struct {
	CityID string
	Lat    string
	Long   string
}

/*
GetRecipients returns a slice of email addresses
to send forecast data to
*/
func (m *Mail) GetRecipients() []*Recipient {
	var recipients []*Recipient

	recipients = append(recipients, &Recipient{
		Email: os.Getenv("JOHN_EMAIL"),
		Location: RecipientLocation{
			CityID: "5110302",
			Lat:    "40.650101",
			Long:   "-73.949577",
		},
	})
	recipients = append(recipients, &Recipient{
		Email: os.Getenv("MAGGIE_EMAIL"),
		Location: RecipientLocation{
			CityID: "5133268",
			Lat:    "40.7001",
			Long:   "-73.799583",
		},
	})

	return recipients
}
