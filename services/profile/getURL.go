package profile

import (
	"fmt"
	"time"

	"cloud.google.com/go/storage"
)

func getURL(uuid string) (string, error) {
	objectPath := fmt.Sprintf("avatars/%s/original.webp", uuid)

	opts := &storage.SignedURLOptions{
		Method:      "PUT",
		Expires:     time.Now().Add(10 * time.Minute),
		ContentType: "image/webp",
	}
	url, err := storage.SignedURL("avatars", objectPath, opts)
	if err != nil {
		return "", err
	}
	return url, nil
}
