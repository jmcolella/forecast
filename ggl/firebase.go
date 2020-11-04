package ggl

import (
	"os"

	firebase "firebase.google.com/go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

func NewFirebaseApp(ctx context.Context) (*firebase.App, error) {
	creds := []byte(os.Getenv("GOOGLE_FIREBASE_CREDENTIALS"))
	opt := option.WithCredentialsJSON(creds)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	return app, nil
}
