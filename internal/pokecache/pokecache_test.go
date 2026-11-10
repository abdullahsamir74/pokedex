package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGet(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cases := []struct {
		inputKey   string
		inputValue []byte
	}{
		{
			inputKey:   "key1",
			inputValue: []byte("value1"),
		},
		{
			inputKey:   "key2",
			inputValue: []byte("value2"),
		},
		{
			inputKey:   "",
			inputValue: []byte("value3"),
		},
	}

	for _, c := range cases {
		cache.Add(c.inputKey, c.inputValue)
		actual, ok := cache.Get(c.inputKey)
		if !ok {
			t.Error("key not found")
		}
		if string(actual) != string(c.inputValue) {
			t.Error("value mismatch")
		}
	}

}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("value1"))

	time.Sleep(interval + time.Millisecond)

	if _, ok := cache.Get(keyOne); ok {
		t.Errorf("%s should have been reaped", keyOne)
	}

}
