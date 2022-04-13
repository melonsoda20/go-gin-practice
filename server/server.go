package server

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// Use a service account
func ConnectFirestore() (c *firebase.App, f_c *firestore.Client, err error) {
	ctx := context.Background()
	opts := option.WithCredentialsFile("../config/firebaseAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, opts)
	if err != nil {
		return nil, nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, nil, err
	}

	return app, client, err
}
