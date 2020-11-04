package ggl

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type Firestore struct {
	Store *firestore.Client
}

func (f *Firestore) GetCollectionDocuments(ctx context.Context, collectionName string) ([]*firestore.DocumentSnapshot, error) {
	iter := f.Store.Collection(collectionName).Documents(ctx)

	var data []*firestore.DocumentSnapshot

	for {
		item, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		data = append(data, item)
	}

	return data, nil
}

func NewFireStoreClient(ctx context.Context) (*Firestore, error) {
	app, err := NewFirebaseApp(ctx)
	if err != nil {
		return nil, err
	}

	store, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	return &Firestore{
		Store: store,
	}, nil
}
