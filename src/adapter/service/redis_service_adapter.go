package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/KhaiHust/authen_service/adapter/properties"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/go-redis/redis/v8"
	"github.com/golibs-starter/golib/log"
	"time"
)

type RedisServiceAdapter struct {
	props       *properties.RedisProperties
	redisClient *redis.Client
}

func (r RedisServiceAdapter) SetToCache(ctx context.Context, key string, value interface{}, ttl int) error {
	data, err := json.Marshal(value)
	if err != nil {
		log.Error(ctx, "marshal data into redis error %v", err)
		return err
	}
	result, err := r.redisClient.Set(ctx, key, data, time.Duration(ttl)*time.Second).Result()
	if err != nil {
		log.Error(ctx, "set key %s, err %v", key)
		return err
	}
	log.Info(ctx, "set key %s, value %v, expire %d, result %v", key, value, ttl, result)
	return nil
}

func (r RedisServiceAdapter) GetFromCache(ctx context.Context, key string) (interface{}, error) {
	val, err := r.redisClient.Get(ctx, key).Result()
	switch {
	case errors.Is(err, redis.Nil):
		log.Info(ctx, "key does not exist")
		return nil, nil
	case err != nil:
		log.Error(ctx, "Get failed", err)
		return nil, err
	}
	return val, nil
}

func NewRedisServiceAdapter(props *properties.RedisProperties, redisClient *redis.Client) port.ICachePort {
	return &RedisServiceAdapter{props: props, redisClient: redisClient}
}
