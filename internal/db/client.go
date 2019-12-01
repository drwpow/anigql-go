package db

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func Client(ctx context.Context) (*firestore.Client, error) {
	return firestore.NewClient(ctx, os.Getenv("FIRESTORE_PROJECT_ID"))
}
