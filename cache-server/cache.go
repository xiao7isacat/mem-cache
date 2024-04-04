package cache_server

import (
	"memCache/cache"
	"time"
)

type CacheServer struct {
	cache cache.CacheInterface
}

func Newmemcache() *CacheServer {
	var cacheServer = CacheServer{
		cache.NewCache(),
	}
	return &cacheServer
}

func (this *CacheServer) SetMaxMemory(size string) bool {
	return this.cache.SetMaxMemory(size)
}

func (this *CacheServer) Set(key string, val interface{}, expire ...time.Duration) bool {
	expireTs := time.Second * 1
	if len(expire) > 0 {
		expireTs = expire[0]
	}
	return this.cache.Set(key, val, expireTs)
}

func (this *CacheServer) Get(key string) (interface{}, bool) {
	return this.cache.Get(key)
}

// 删除k
func (this *CacheServer) Del(key string) bool {
	return this.cache.Del(key)
}

// 判断k是否存在
func (this *CacheServer) Exist(key string) bool {
	return this.Exist(key)
}

// 清空所有的k
func (this *CacheServer) Flush() bool {
	return this.cache.Flush()
}

// 获取缓存中的所有k
func (this *CacheServer) Keys() []string {
	return this.cache.Keys()
}

// 获取key的数量
func (this *CacheServer) KeysNum() int64 {
	return this.cache.KeysNum()
}
