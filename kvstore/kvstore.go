package kvstore

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

type KVStore interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Delete(key string) error
	Close() error
}

type RedisStore struct {
	client *redis.Client
}

// come back to this and look at alternate methods and error handling...
// https://redis.io/docs/connect/clients/go/
func ConnectRedis() *RedisStore {
	c := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	fmt.Println(c)

	return &RedisStore{
		client: c,
	}
}

func (r *RedisStore) Set(key string, value string) error {
	// Implement using Redis client
	return nil
}

func (r *RedisStore) Get(key string) (string, error) {
	// Implement using Redis client
	return "test", nil
}

func (r *RedisStore) Delete(key string) error {
	// Implement using Redis client
	return nil
}

func (r *RedisStore) Close() error {
	return r.client.Close()
}
