package vacations

import "forecast/mail"

/*
Vacation is the struct that holds the
name and Lat/Long of the getaway place
*/
type Vacation struct {
	Name     string
	Location mail.RecipientLocation
}

/*
GetLocation returns the Long/Lat of the current vacation spot, if available
*/
func GetLocation() *Vacation {
	return &Vacation{
		Name: "Stowe",
		Location: mail.RecipientLocation{
			Lat:  "44.475278",
			Long: "-72.702222",
		},
	}
}
