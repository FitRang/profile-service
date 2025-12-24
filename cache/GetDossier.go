package cache

import (
	"context"
	"encoding/json"

	"github.com/Foxtrot-14/FitRang/profile-service/graph/model"
	"github.com/redis/go-redis/v9"
)

func (r *RedisClient) GetDossier(
	ctx context.Context,
	key string,
) (*model.Dossier, error) {

	val, err := r.rdb.Get(ctx, "dossier:"+key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, err
		}
		return nil, err
	}

	var dossier model.Dossier
	if err := json.Unmarshal([]byte(val), &dossier); err != nil {
		return nil, err
	}

	return &dossier, nil
}
