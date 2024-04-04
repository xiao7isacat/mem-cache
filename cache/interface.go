package cache

import "time"

func NewCache() CacheInterface {
	memCache := &MemCache{
		exprireTime: 5 * time.Second,
		values:      make(map[string]*memCacheValue),
	}
	go memCache.clearExpiredTime()
	return memCache
}

type CacheInterface interface {
	//size: 1kb 100kb 1mb 2mb 1gb
	SetMaxMemory(size string) bool
	//将k写入缓存
	Set(key string, val interface{}, expire time.Duration) bool
	//根据k获取v
	Get(key string) (interface{}, bool)
	//删除k
	Del(key string) bool
	//判断k是否存在
	Exist(key string) bool
	//清空所有的k
	Flush() bool
	//获取缓存中的所有k
	Keys() []string
	//获取key的数量
	KeysNum() int64
}
