package users

import (
	"context"
	"fmt"
	"forecast/ggl"
	"log"

	"google.golang.org/genproto/googleapis/type/latlng"
)

type Location struct {
	Lat  string
	Long string
}

type User struct {
	Name     string
	Email    string
	Location Location
}

type firestoreUser struct {
	Name     string         `json:"name"`
	Email    string         `json:"email"`
	Location *latlng.LatLng `json:"location"`
}

func Fetch(ctx context.Context, firestoreClient *ggl.Firestore) ([]*User, error) {
	data, err := firestoreClient.GetCollectionDocuments(ctx, "users")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var users []*User

	for _, d := range data {
		fu := firestoreUser{}

		err = d.DataTo(&fu)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		users = append(users, &User{
			Name:  fu.Name,
			Email: fu.Email,
			Location: Location{
				Lat:  fmt.Sprintf("%f", fu.Location.GetLatitude()),
				Long: fmt.Sprintf("%f", fu.Location.GetLongitude()),
			},
		})
	}

	return users, nil
}
