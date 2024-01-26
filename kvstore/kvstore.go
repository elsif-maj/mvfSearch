package kvstore

import (
	"context"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type KVStore interface {
	Set(key string, value int) error
	SAdd(userId string, searchterm string, docId int) error
	Get(key string) error
	Delete(key string) error
	Close() error
}

type RedisStore struct {
	client *redis.Client
}

// come back to this and look at alternate methods and error handling...
// https://redis.io/docs/connect/clients/go/
// extract to env
func ConnectRedis() *RedisStore {
	c := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	fmt.Println(c.Ping(context.Background()))

	return &RedisStore{
		client: c,
	}
}

// change this to look like SAdd
func (r *RedisStore) Set(key string, value int) error {
	strDocId := strconv.Itoa(value)
	err := r.client.Set(context.Background(), strDocId, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisStore) SAdd(userId string, searchterm string, docId int) error {
	key := fmt.Sprintf("user:%s:search:%s", userId, searchterm)
	strDocId := strconv.Itoa(docId)

	err := r.client.SAdd(context.Background(), key, strDocId).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisStore) Get(key string) error {
	// strKey := strconv.Itoa(key)
	return nil
}

func (r *RedisStore) SMembers(key string) ([]string, error) {
	result, err := r.client.SMembers(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *RedisStore) Delete(key string) error {
	// strKey := strconv.Itoa(key)
	return nil
}

func (r *RedisStore) Close() error {
	return r.client.Close()
}
