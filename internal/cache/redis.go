package cache

import (
	"backend_template/internal/user/model"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	Client *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{Client: client}
}

func (r *RedisCache) GetUserProfile(ctx context.Context, userID uint) (*model.UserModel, error) {
	key := fmt.Sprintf("user_profile:%d", userID)
	val, err := r.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var user model.UserModel
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *RedisCache) SetUserProfile(ctx context.Context, user *model.UserModel, duration time.Duration) error {
	key := fmt.Sprintf("user_profile:%d", user.ID)
	b, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return r.Client.Set(ctx, key, b, duration).Err()
}

func (r *RedisCache) DeleteUserProfile(ctx context.Context, userID uint) error {
	key := fmt.Sprintf("user_profile:%d", userID)
	return r.Client.Del(ctx, key).Err()
}
