package main

import (
	"fmt"
	cache_server "memCache/cache-server"
	"time"
)

func main() {
	cache := cache_server.Newmemcache()
	cache.SetMaxMemory("6kB")
	cache.Set("int", 1)
	cache.Set("bool", false, 100*time.Second)
	cache.Set("date", map[string]interface{}{"a": 1})
	cache.Get("int")
	cache.Del("int")
	//cache.Flush()
	cache.Keys()
	cache.KeysNum()
	//time.Sleep(time.Second * 1)
	fmt.Println(cache.Get("bool"))
	fmt.Println(cache.Get("date"))
}
