package caches

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisCache implements the Cache interface using Redis
type RedisCache struct {
	client *redis.Client
	ctx    context.Context
}

// NewRedisCache creates a new instance of RedisCache
func NewRedisCache(addr string) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &RedisCache{
		client: rdb,
		ctx:    context.Background(),
	}
}

// Set adds a key-value pair to the cache with an expiration time
func (rc *RedisCache) Set(key string, value interface{}, expiration time.Duration) error {
	return rc.client.Set(rc.ctx, key, value, expiration).Err()
}

// Get retrieves the value associated with a key from the cache
func (rc *RedisCache) Get(key string) (string, error) {
	val, err := rc.client.Get(rc.ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key not found")
	}
	return val, err
}

// Update modifies the value associated with an existing key
func (rc *RedisCache) Update(key string, value interface{}) error {
	// Check if the key exists
	if _, err := rc.Get(key); err != nil {
		return err
	}
	return rc.Set(key, value, 0) // Set with no expiration to update the value
}

// Delete removes a key-value pair from the cache
func (rc *RedisCache) Delete(key string) error {
	return rc.client.Del(rc.ctx, key).Err()
}
