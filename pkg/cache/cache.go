package cache

import (
	c "github.com/patrickmn/go-cache"
	"time"
)

const (
	Default       = c.DefaultExpiration
	OneMinute     = time.Minute
	FiveMinute    = OneMinute * 5
	TenMinutes    = OneMinute * 10
	ThirtyMinutes = TenMinutes * 3
	OneHour       = time.Hour
	FourHours     = OneHour * 4
	EightHours    = OneHour * 8
	TwelveHours   = OneHour * 12
	Forever       = c.NoExpiration
)

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, item Cacheable)
	GetSet(key string, expiration time.Duration, getFunc func() (interface{}, error)) (interface{}, error)
	Delete(key string)
}

type Cacheable struct {
	value      interface{}
	expiration time.Duration
}

type cache struct {
	client *c.Cache
}

func NewCache() Cache {
	return &cache{
		client: c.New(FourHours, TenMinutes),
	}
}

func (c *cache) Get(key string) (interface{}, bool) {
	v, ok := c.client.Get(key)
	if !ok {
		return nil, false
	}

	return v, true
}

func (c *cache) Set(key string, item Cacheable) {
	c.client.Set(key, item.value, item.expiration)
}

func (c *cache) GetSet(key string, expiration time.Duration, getFunc func() (
	interface{}, error)) (interface{}, error) {
	if v, ok := c.Get(key); ok {
		return v, nil
	}

	value, err := getFunc()
	if err != nil {
		return nil, err
	}

	item := Cacheable{
		value:      value,
		expiration: expiration,
	}
	c.Set(key, item)

	return value, nil
}

func (c *cache) Delete(key string) {
	c.client.Delete(key)
}
