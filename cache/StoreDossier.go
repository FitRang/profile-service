package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Foxtrot-14/FitRang/profile-service/graph/model"
)

func (r *RedisClient) StoreDossier(
	ctx context.Context,
	key string,
	dossier model.Dossier,
) error {

	bytes, err := json.Marshal(dossier)
	if err != nil {
		return err
	}

	return r.rdb.Set(
		ctx,
		"dossier:"+key,
		bytes,
		24*time.Hour,
	).Err()
}
