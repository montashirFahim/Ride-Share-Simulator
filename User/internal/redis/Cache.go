package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

func SetClient(client *redis.Client) {
	Client = client
}

func GetClient() *redis.Client {
	return Client
}

func InitRedis(host, port string) error {
	Client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := Client.Ping(ctx).Err(); err != nil {
		return err
	}
	return nil
}

// func GetUserCache(userID int) (*model.UserResponse, error) {
// 	ctx := context.Background()
// 	key := fmt.Sprintf("user:info:%d", userID)

// 	val, err := Client.Get(ctx, key).Result()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var user model.UserResponse
// 	if err := json.Unmarshal([]byte(val), &user); err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func SetUserCache(userID int, user *model.UserResponse) error {
// 	ctx := context.Background()
// 	key := fmt.Sprintf("user:info:%d", userID)

// 	data, err := json.Marshal(user)
// 	if err != nil {
// 		return err
// 	}

// 	return Client.Set(ctx, key, data, 1*time.Minute).Err()
// }
