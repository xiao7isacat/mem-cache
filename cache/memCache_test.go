package cache

import (
	"testing"
	"time"
)

func TestCacheOP(t *testing.T) {
	testData := []struct {
		key    string
		val    interface{}
		expire time.Duration
	}{
		{"test1", "testval1", time.Second * 1000},
		{"test2", "testval2", time.Second * 7},
		{"test3", "testval3", time.Second * 1000},
		{"test4", "testval4", time.Second * 7},
		{"test5", "testval5", time.Second * 7},
		{"test6", "testval6", time.Second * 7},
		{"test7", "testval7", time.Second * 7},
	}

	c := NewCache()
	c.SetMaxMemory("20kb")
	for _, item := range testData {
		c.Set(item.key, item.val, item.expire)
		_, ok := c.Get(item.key)
		if !ok {
			t.Errorf("%s缓存读取失败", item.key)
		}
	}

	if int64(len(testData)) != c.KeysNum() {
		t.Errorf("数量不一致")
	}

	c.Del(testData[0].key)
	c.Del(testData[1].key)
	if int64(len(testData)) != c.KeysNum()+2 {
		t.Errorf("删除失败")
	}

	time.Sleep(time.Second * 10)
	if c.KeysNum() != 1 {
		t.Errorf("缓存清空失败%v", c.KeysNum())
	}
}
