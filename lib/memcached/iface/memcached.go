package iface

import "github.com/bradfitz/gomemcache/memcache"

type MemcachedIface interface {
	Get(key string) (item *memcache.Item, err error)
	Set(item *memcache.Item) error
	GetMulti(keys []string) (map[string]*memcache.Item, error)
	Delete(key string) error
}
