package profile

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Foxtrot-14/FitRang/profile-service/graph/model"
	"github.com/Foxtrot-14/FitRang/profile-service/middleware"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (p *ProfileService) UploadProfile(ctx context.Context, file graphql.Upload) (*model.Profile, error) {
	emailID, ok := ctx.Value(middleware.EmailKey).(string)
	if !ok || emailID == "" {
		return nil, errors.New("unauthenticated")
	}

	const maxSize = 5 << 20

	if file.Size > maxSize {
		return nil, errors.New("file size exceeds 5MB")
	}

	contentType := file.ContentType
	switch contentType {
	case "image/jpeg", "image/png", "image/webp":
	default:
		return nil, errors.New("unsupported image format")
	}

	uploadDir := "uploads/profile"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return nil, err
	}

	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%s_%d%s", emailID, time.Now().Unix(), ext)
	path := filepath.Join(uploadDir, filename)

	dst, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file.File); err != nil {
		return nil, err
	}

	imageURL := fmt.Sprintf("/%s/%s", uploadDir, filename)

	res := p.Repo.Col.FindOneAndUpdate(
		ctx,
		bson.M{"email": emailID},
		bson.M{
			"$set": bson.M{
				"profilePicture": imageURL,
				"updatedAt":      time.Now().UTC().Format(time.RFC3339),
			},
		},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)

	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return nil, errors.New("profile not found")
		}
		return nil, res.Err()
	}

	var profile model.Profile
	if err := res.Decode(&profile); err != nil {
		return nil, err
	}

	p.sendProfileToIndex(profile)

	return &profile, nil
}
