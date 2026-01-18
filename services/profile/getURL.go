package profile

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/storage"
)

const bucketName = "fitrang-bucket"

func getURL(ctx context.Context, userUUID string) (string, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("storage.NewClient: %w", err)
	}

	objectPath := fmt.Sprintf(
		"avatars/raw/%s/original.webp",
		userUUID,
	)

	opts := &storage.SignedURLOptions{
		Method:      "PUT",
		Expires:     time.Now().Add(10 * time.Minute),
		ContentType: "image/webp",
	}

	url, err := client.Bucket(bucketName).SignedURL(objectPath, opts)
	if err != nil {
		return "", err
	}

	return url, nil
}
