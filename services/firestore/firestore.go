package firestore

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
	"lendotopia.com/originator/config"
	"lendotopia.com/originator/services/gcp"
)

var Client *firestore.Client

func Init() {
	var err error

	cfg := config.Get()

	credentials, err := gcp.GetCredentials()

	// Create an option with the service account key file
	opt := option.WithCredentialsJSON(credentials)

	// Create a Firestore client with the project ID and option
	Client, err = firestore.NewClient(context.Background(), cfg.ProjectID, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firestore client: %v", err)
	}
}
