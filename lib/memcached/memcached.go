package memcached

import (
	"fmt"

	"time"

	"reflect"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/rodrigo-brito/bus-api-go/lib/memcached/iface"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

type contextKey struct{}

type CacheManager struct {
	BPC   bool
	cache iface.MemcachedIface
}

func NewMemcached(BPC bool) *CacheManager {
	cache := new(CacheManager)
	cache.initConnection()
	cache.BPC = BPC
	return cache
}

func (c *CacheManager) initConnection() {
	settings := viper.GetStringMapString("memcached")
	c.cache = memcache.New(fmt.Sprintf("%s:%s", settings["address"],
		settings["port"]))
}

func (c *CacheManager) Get(key string, value interface{}) (bool, error) {
	if c.BPC {
		return false, nil
	}

	item, err := c.cache.Get(key)
	if err == memcache.ErrCacheMiss {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, decode(item.Value, value)
}

func (c *CacheManager) Set(key string, value interface{}, TTL time.Duration) error {
	if c.BPC {
		return nil
	}

	bin, err := encode(value)
	if err != nil {
		return err
	}

	return c.cache.Set(&memcache.Item{
		Key:        key,
		Value:      bin,
		Expiration: int32(TTL.Seconds()),
	})
}

func (c *CacheManager) GetSet(key string, value interface{}, getValue func() (interface{}, error), TTL time.Duration) error {
	hit, err := c.Get(key, value)
	if err != nil {
		return err
	}

	if !hit || c.BPC {
		fromDB, err := getValue()
		if err != nil {
			return err
		}
		reflectValue := reflect.Indirect(reflect.ValueOf(value))
		reflectValue.Set(reflect.Indirect(reflect.ValueOf(fromDB)))
		return c.Set(key, value, TTL)
	}

	return nil
}

func NewContext(parent context.Context, BPC bool) context.Context {
	return context.WithValue(parent, contextKey{}, NewMemcached(BPC))
}

func FromContext(ctx context.Context) *CacheManager {
	return ctx.Value(contextKey{}).(*CacheManager)
}
