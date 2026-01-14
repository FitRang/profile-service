package profile

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/storage"
)

func createUploadSignedURL(
	ctx context.Context,
	bucketName string,
	objectName string,
	expiry time.Duration,
	contentType string,
) (string, error) {

	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("storage.NewClient: %w", err)
	}
	defer client.Close()

	opts := &storage.SignedURLOptions{
		Scheme:      storage.SigningSchemeV4,
		Method:      "PUT",
		Expires:     time.Now().Add(expiry),
		ContentType: contentType,
	}

	url, err := client.Bucket(bucketName).SignedURL(objectName, opts)
	if err != nil {
		return "", fmt.Errorf("SignedURL: %w", err)
	}

	return url, nil
}
