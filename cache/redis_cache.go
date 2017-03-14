package cache

import (
	"strconv"
	"time"

	"github.com/michelaquino/golang_api_skeleton/api_errors"
	"github.com/michelaquino/golang_api_skeleton/context"

	redis "gopkg.in/redis.v4"
)

var cacheLogger context.Logger

func init() {
	cacheLogger = context.GetLogger()
}

// RedisCache is a redis cache object
type RedisCache struct {
	client *redis.Client
}

// NewRedisCache returns a new instance of the RedisCache
func NewRedisCache() *RedisCache {

	var redisClient *redis.Client
	cacheLogger.Info("NewRedisCache", "Constructor", "", "", "", "", "")
	redisClient = redis.NewClient(&redis.Options{
		ReadTimeout:  time.Duration(1) * time.Second,
		WriteTimeout: time.Duration(1) * time.Second,
		Addr:         "http://localhost",
		Password:     "123456",
		DB:           0,
		PoolSize:     5000,
	})

	if _, err := redisClient.Ping().Result(); err != nil {
		cacheLogger.Error("NewRedisCache", "Constructor", "", "", "Ping Redis", err.Error(), "Error when connect on redis")
	}

	return &RedisCache{
		client: redisClient,
	}
}

// Get is a method that get an value from cache
func (r RedisCache) Get(key string) (string, error) {
	cacheValue, err := r.client.Get(key).Result()
	if err == redis.Nil {
		cacheLogger.Debug("RedisCache", "Get", "", "", "Get key", redis.Nil.Error(), "Key "+key+" not found on cache")
		return "", apiErrors.ErrNotFoundOnCache
	}

	if err != nil {
		cacheLogger.Error("RedisCache", "Get", "", "", "Get key", err.Error(), "An error occur when get cache value with key "+key)
		return "", apiErrors.ErrGetCacheValue
	}

	return cacheValue, nil
}

// Set is a method that set a value to cache
func (r RedisCache) Set(key, value string, expireInSec int) error {
	expire := time.Duration(expireInSec) * time.Second

	err := r.client.Set(key, value, expire).Err()
	if err != nil {
		cacheLogger.Error("RedisCache", "Set", "", "", "Set key/value on cache", err.Error(), "An error occur when set cache value with key "+key)
	}

	cacheLogger.Debug("RedisCache", "Set", "", "", "Set key/value on cache", "Success", "Set key "+key+" on cache with expire in "+strconv.Itoa(expireInSec)+" seconds")
	return err
}
