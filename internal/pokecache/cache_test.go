package pokecache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cache := NewCache(time.Second * 1)
	cache.Add("test1", getValue("test1"))
	value, ok := cache.Get("test1")
	if !ok {
	    t.Errorf("Key not present in cache after adding")
	}
	if string(value) != string(getValue("test1")) {
	    t.Errorf("Wrong value returned")
	}
	time.Sleep(2 * time.Second)
	if _, ok := cache.Get("test1"); ok {
		t.Errorf("Key present in cache after expiration interval")
	}
}

func getValue(key string) []byte {
	return []byte("value_for_" + key)
}

