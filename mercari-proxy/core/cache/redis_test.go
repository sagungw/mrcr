package cache_test

import (
	"context"
	"sagungw/mercari/core/cache"
	"sagungw/mercari/core/config"
	"sagungw/mercari/core/indoarea"
	"testing"
	"time"

	"github.com/pkg/errors"
)

func TestCache(t *testing.T) {
	c, err := cache.NewRedisCache(config.RedisAddress())
	if err != nil {
		t.Error(err)
	}

	t.Run("cache miss", func(t *testing.T) {
		val := &indoarea.Province{}
		err := c.Get(context.Background(), "cache-key-2", val)
		if err == nil {
			t.Error("cache should miss")
		} else {
			if !errors.Is(err, cache.CacheMissError{}) {
				t.Error(err)
			}
		}
	})

	t.Run("cache hit", func(t *testing.T) {
		c.Set(context.Background(), "cache-key", &indoarea.Province{Id: "1", Name: "Solo"}, 1*time.Hour)

		val := &indoarea.Province{}
		err := c.Get(context.Background(), "cache-key", val)
		if err != nil {
			t.Error(err)
		}
	})
}
