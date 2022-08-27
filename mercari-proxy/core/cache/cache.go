package cache

import (
	"context"
	"time"

	"github.com/golang/protobuf/proto"
)

type Cache interface {
	Set(ctx context.Context, key string, val proto.Message, ttl time.Duration) error
	Get(ctx context.Context, key string, val proto.Message) error
}

type CacheMissError struct{}

func (c CacheMissError) Error() string {
	return "cache miss error"
}
