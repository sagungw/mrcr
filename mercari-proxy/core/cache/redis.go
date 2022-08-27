package cache

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/go-redis/redis/v8"
	"github.com/golang/protobuf/proto"
)

type redisCache struct {
	client *redis.Client
}

func NewRedisCache(redisAddress string) (*redisCache, error) {
	opt, err := redis.ParseURL(redisAddress)
	if err != nil {
		return nil, errors.Wrap(err, "cache: error connecting to redis")
	}

	opt.MaxRetries = 5

	return &redisCache{
		client: redis.NewClient(opt),
	}, nil
}

func (r *redisCache) Set(ctx context.Context, key string, val proto.Message, ttl time.Duration) error {
	b, err := proto.Marshal(val)
	if err != nil {
		return errors.Wrap(err, "cache: error encoding bytes")
	}

	r.client.Set(ctx, key, b, ttl)
	return nil
}

func (r *redisCache) Get(ctx context.Context, key string, val proto.Message) error {
	s := r.client.Get(ctx, key)
	t := r.client.TTL(ctx, key)

	ttl, err := t.Result()
	if err != nil {
		return errors.Wrap(err, "cache: error fetching ttl")
	}

	if ttl <= 0 {
		return CacheMissError{}
	}

	b, err := s.Bytes()
	if err != nil {
		return errors.Wrap(err, "cache: error fetching bytes")
	}

	err = proto.Unmarshal(b, val)
	if err != nil {
		return errors.Wrap(err, "cache: error decoding bytes")
	}

	return nil
}
