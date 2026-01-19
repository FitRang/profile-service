package profile

import (
	"context"

	"github.com/Foxtrot-14/FitRang/profile-service/apperror"
	"github.com/Foxtrot-14/FitRang/profile-service/middleware"
)

func getEmailFromContext(ctx context.Context) (string, error) {
	email, ok := ctx.Value(middleware.EmailKey).(string)
	if !ok || email == "" {
		return "", apperror.New(
			apperror.Unauthenticated,
			"Authentication required",
		)
	}
	return email, nil
}
