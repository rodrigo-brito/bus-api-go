package memcached

import (
	"fmt"

	"sync"

	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/rodrigo-brito/bus-api-go/lib/memcached/iface"
	"github.com/spf13/viper"
)

var (
	once  sync.Once
	cache iface.MemcachedIface
)

func getConnection() iface.MemcachedIface {
	once.Do(func() {
		initConnection()
	})
	return cache
}

func initConnection() {
	settings := viper.GetStringMapString("memcached")
	cache = memcache.New(fmt.Sprintf("%s:%s", settings["address"],
		settings["port"]))
}

func Get(key string, value interface{}) (bool, error) {
	conn := getConnection()
	item, err := conn.Get(key)
	if err == memcache.ErrCacheMiss {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, decode(item.Value, value)
}

func Set(key string, value interface{}, TTL time.Duration) error {
	conn := getConnection()

	bin, err := encode(value)
	if err != nil {
		return err
	}

	return conn.Set(&memcache.Item{
		Key:        key,
		Value:      bin,
		Expiration: int32(TTL.Seconds()),
	})
}
