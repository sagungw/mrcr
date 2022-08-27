package indoarea

import (
	context "context"
	"time"

	"sagungw/mercari/core/cache"

	"github.com/pkg/errors"
	grpc "google.golang.org/grpc"
)

//go:generate protoc -I${PWD}/core/indoarea --gofast_out=plugins=grpc:${PWD}/core/indoarea/. indoarea.proto

type IndoareaService interface {
	GetProvinces(ctx context.Context) (*GetProvincesResponse, error)
	GetCities(ctx context.Context) (*GetCitiesResponse, error)
	GetDistricts(ctx context.Context) (*GetDistrictsResponse, error)
	GetSubDistricts(ctx context.Context) (*GetSubDistrictsResponse, error)
}

type cacheAwareClient struct {
	c        cache.Cache
	cacheTtl time.Duration
	client   ApiClient
}

func NewCacheAwareClient(c cache.Cache, cacheTtl time.Duration, client ApiClient) IndoareaService {
	return &cacheAwareClient{
		c:        c,
		cacheTtl: cacheTtl,
		client:   client,
	}
}

func (c *cacheAwareClient) GetProvinces(ctx context.Context) (*GetProvincesResponse, error) {
	key := "proxy:GetProvinces"
	b := &GetProvincesResponse{}
	err := c.c.Get(ctx, key, b)
	if err != nil {
		if errors.Is(err, cache.CacheMissError{}) {
			message, err := c.client.GetProvinces(ctx, &Void{})
			if err != nil {
				return nil, errors.Wrap(err, "indoarea: error fetching data")
			}

			err = c.c.Set(ctx, key, message, c.cacheTtl)
			if err != nil {
				return nil, err
			}

			return message, nil
		}

		return nil, err
	}

	return b, nil
}

func (c *cacheAwareClient) GetCities(ctx context.Context) (*GetCitiesResponse, error) {
	key := "proxy:GetCities"
	b := &GetCitiesResponse{}
	err := c.c.Get(ctx, key, b)
	if err != nil {
		if errors.Is(err, cache.CacheMissError{}) {
			message, err := c.client.GetCities(ctx, &Void{})
			if err != nil {
				return nil, errors.Wrap(err, "indoarea: error fetching data")
			}

			err = c.c.Set(ctx, key, message, c.cacheTtl)
			if err != nil {
				return nil, err
			}

			return message, nil
		}

		return nil, err
	}

	return b, nil
}

func (c *cacheAwareClient) GetDistricts(ctx context.Context) (*GetDistrictsResponse, error) {
	key := "proxy:GetDistricts"
	b := &GetDistrictsResponse{}
	err := c.c.Get(ctx, key, b)
	if err != nil {
		if errors.Is(err, cache.CacheMissError{}) {
			message, err := c.client.GetDistricts(ctx, &Void{})
			if err != nil {
				return nil, errors.Wrap(err, "indoarea: error fetching data")
			}

			err = c.c.Set(ctx, key, message, c.cacheTtl)
			if err != nil {
				return nil, err
			}

			return message, nil
		}

		return nil, err
	}

	return b, nil
}

func (c *cacheAwareClient) GetSubDistricts(ctx context.Context) (*GetSubDistrictsResponse, error) {
	key := "proxy:GetSubDistricts"
	b := &GetSubDistrictsResponse{}
	err := c.c.Get(ctx, key, b)
	if err != nil {
		if errors.Is(err, cache.CacheMissError{}) {
			message, err := c.client.GetSubDistricts(ctx, &Void{})
			if err != nil {
				return nil, errors.Wrap(err, "indoarea: error fetching data")
			}

			err = c.c.Set(ctx, key, message, c.cacheTtl)
			if err != nil {
				return nil, err
			}

			return message, nil
		}

		return nil, err
	}

	return b, nil
}

type clientStub struct {
}

func NewClientStub() ApiClient {
	return &clientStub{}
}

func (c *clientStub) GetProvinces(ctx context.Context, in *Void, opts ...grpc.CallOption) (*GetProvincesResponse, error) {
	// to mock long response time
	time.Sleep(5 * time.Second)

	return &GetProvincesResponse{
		Provinces: []*Province{
			{Id: "1", Name: "Jawa Tengah"},
		},
	}, nil
}

func (c *clientStub) GetCities(ctx context.Context, in *Void, opts ...grpc.CallOption) (*GetCitiesResponse, error) {
	// to mock long response time
	time.Sleep(5 * time.Second)

	return &GetCitiesResponse{
		Cities: []*City{
			{Id: "1", Name: "Semarang", ProvinceName: "Jawa Tengah"},
			{Id: "2", Name: "Surakarta", ProvinceName: "Jawa Tengah"},
		},
	}, nil
}

func (c *clientStub) GetDistricts(ctx context.Context, in *Void, opts ...grpc.CallOption) (*GetDistrictsResponse, error) {
	// to mock long response time
	time.Sleep(5 * time.Second)

	return &GetDistrictsResponse{
		Districts: []*District{
			{Id: "1", Name: "Banyumanik", CityName: "Semarang"},
			{Id: "2", Name: "Banjarsari", CityName: "Surakarta"},
		},
	}, nil
}

func (c *clientStub) GetSubDistricts(ctx context.Context, in *Void, opts ...grpc.CallOption) (*GetSubDistrictsResponse, error) {
	// to mock long response time
	time.Sleep(5 * time.Second)

	return &GetSubDistrictsResponse{
		SubDistricts: []*SubDistrict{
			{Id: "1", Name: "Gedawang", DistrictName: "Banyumanik"},
			{Id: "1", Name: "Banyuanyar", DistrictName: "Banjarsari"},
		},
	}, nil
}
