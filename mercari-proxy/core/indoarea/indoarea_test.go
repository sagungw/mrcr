package indoarea_test

import (
	context "context"
	"sagungw/mercari/core/cache"
	"sagungw/mercari/core/config"
	"sagungw/mercari/core/indoarea"
	"testing"
	"time"

	"github.com/pkg/errors"

	grpc "google.golang.org/grpc"
)

func TestGetProvince(t *testing.T) {
	cache, err := cache.NewRedisCache(config.RedisAddress())
	if err != nil {
		t.Error(err)
	}

	t.Run("returns error", func(t *testing.T) {
		indoareaService := indoarea.NewCacheAwareClient(cache, 1*time.Millisecond, NewClientStub())
		_, err := indoareaService.GetProvinces(context.Background())
		if err == nil {
			t.Error("should error")
		}

		_, err = indoareaService.GetCities(context.Background())
		if err == nil {
			t.Error("should error")
		}

		_, err = indoareaService.GetDistricts(context.Background())
		if err == nil {
			t.Error("should error")
		}

		_, err = indoareaService.GetSubDistricts(context.Background())
		if err == nil {
			t.Error("should error")
		}
	})

	t.Run("returns result", func(t *testing.T) {
		indoareaService := indoarea.NewCacheAwareClient(cache, 24*time.Hour, indoarea.NewClientStub())
		_, err := indoareaService.GetProvinces(context.Background())
		if err != nil {
			t.Error("should not error")
		}
	})
}

type alwaysErrorClientStub struct {
}

func NewClientStub() indoarea.ApiClient {
	return &alwaysErrorClientStub{}
}

func (c *alwaysErrorClientStub) GetProvinces(ctx context.Context, in *indoarea.Void, opts ...grpc.CallOption) (*indoarea.GetProvincesResponse, error) {
	return nil, errors.New("mocked error")
}

func (c *alwaysErrorClientStub) GetCities(ctx context.Context, in *indoarea.Void, opts ...grpc.CallOption) (*indoarea.GetCitiesResponse, error) {
	return nil, errors.New("mocked error")
}

func (c *alwaysErrorClientStub) GetDistricts(ctx context.Context, in *indoarea.Void, opts ...grpc.CallOption) (*indoarea.GetDistrictsResponse, error) {
	return nil, errors.New("mocked error")
}

func (c *alwaysErrorClientStub) GetSubDistricts(ctx context.Context, in *indoarea.Void, opts ...grpc.CallOption) (*indoarea.GetSubDistrictsResponse, error) {
	return nil, errors.New("mocked error")
}
