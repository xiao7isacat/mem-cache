package cache

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type MemCache struct {
	//最大内存
	maxMemSize int64
	//最大内存字符串表示
	maxMemSizestr string
	//当前内存大小
	currMemSize int64
	//缓存键值对
	values map[string]*memCacheValue
	//锁，map和slice是非线程安全的，所以要加锁
	lock sync.RWMutex
	//清空过期缓存的时间周期
	exprireTime time.Duration
}

type memCacheValue struct {
	//value值
	val interface{}
	//过期时间
	expireTime time.Time
	//value 大小
	size int64
}

func (this *MemCache) SetMaxMemory(size string) bool {
	this.maxMemSize, this.maxMemSizestr = ParseSize(size)
	fmt.Println(this.maxMemSize, this.maxMemSizestr)
	return false
}

func (this *MemCache) Set(key string, val interface{}, expire time.Duration) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	v := &memCacheValue{
		val:        val,
		expireTime: time.Now().Add(expire),
		size:       GetValSize(val),
	}
	this.del(key)
	this.add(key, v)
	if this.currMemSize > this.maxMemSize {
		this.del(key)
		log.Fatalf("添加key:%s失败,总内存大小为%v,当前总内存大小%v,当前需要添加的内存大小%v，内存不足", key, this.maxMemSize, this.currMemSize, v.size)
		return false
	}
	return true
}

func (this *MemCache) Get(key string) (interface{}, bool) {
	this.lock.RLock()
	defer this.lock.RUnlock()
	val, ok := this.get(key)
	if ok {
		if val.expireTime.Before(time.Now()) {
			this.del(key)
			return nil, false
		}
		return val.val, ok
	}

	return nil, false
}

// 删除k
func (this *MemCache) Del(key string) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.del(key)
	return false
}

// 判断k是否存在
func (this *MemCache) Exist(key string) bool {
	this.lock.RLock()
	defer this.lock.RUnlock()
	if _, ok := this.get(key); ok {
		return true
	}
	return false
}

// 清空所有的k
func (this *MemCache) Flush() bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.values = make(map[string]*memCacheValue, 0)
	this.currMemSize = 0
	return false
}

// 获取缓存中的所有k
func (this *MemCache) Keys() []string {
	return []string{}
}

// 获取key的数量
func (this *MemCache) KeysNum() int64 {
	this.lock.RLock()
	defer this.lock.RUnlock()
	num := 0
	num = len(this.values)
	return int64(num)
}

func (this *MemCache) get(key string) (*memCacheValue, bool) {
	val, ok := this.values[key]
	if ok {
		return val, true
	}
	return nil, false
}

func (this *MemCache) del(key string) {
	val, ok := this.get(key)
	if ok && val != nil {
		this.currMemSize -= val.size
		delete(this.values, key)
	}

}

func (this *MemCache) add(key string, val *memCacheValue) {
	this.values[key] = val
	this.currMemSize += val.size
}

// 定期清空缓存
func (this *MemCache) clearExpiredTime() {
	timeTicker := time.NewTicker(this.exprireTime)
	defer timeTicker.Stop()
	for {
		select {
		case <-timeTicker.C:
			for key, val := range this.values {
				this.lock.Lock()
				if val.expireTime.Before(time.Now()) {
					this.del(key)
				}
				this.lock.Unlock()
			}

		}

	}
}
