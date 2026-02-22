package repository

import (
	"User/internal/model"
	"User/internal/redis"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

func GetUserCache(userID int) (*model.UserResponse, error) {
	if redis.Client == nil {
		return nil, errors.New("redis client not initialized")
	}

	ctx := context.Background()
	key := fmt.Sprintf("user:info:%d", userID)

	val, err := redis.Client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var user model.UserResponse
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func SetUserCache(userID int, user *model.UserResponse) error {
	if redis.Client == nil {
		return nil // Silently fail if redis not initialized
	}

	ctx := context.Background()
	key := fmt.Sprintf("user:info:%d", userID)

	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return redis.Client.Set(ctx, key, data, 1*time.Minute).Err()
}
