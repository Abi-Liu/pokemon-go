package cache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cases := []struct {
		key   string
		value []byte
	}{
		{
			key:   "https://test.com",
			value: []byte("Testdata"),
		},
		{
			key:   "https://test.com/test",
			value: []byte("More data"),
		},
	}

	interval := 5 * time.Millisecond

	for _, c := range cases {
		cache := CreateCache(interval)
		cache.Add(c.key, c.value)

		val, ok := cache.Get(c.key)
		if !ok {
			t.Errorf("Expected value to be at key - %s", c.key)
			return
		}

		for i, b := range val {
			if b != c.value[i] {
				t.Errorf("Expected %v\nReceived %v", c.value[i], b)
				return
			}
		}
	}
}

func TestReapLoop(t *testing.T) {
	cases := []struct {
		key   string
		value []byte
	}{
		{
			key:   "https://test.com",
			value: []byte("Testdata"),
		},
		{
			key:   "https://test.com/test",
			value: []byte("More data"),
		},
	}

	interval := 5 * time.Millisecond
	buffer := interval + 5*time.Millisecond

	for _, c := range cases {
		cache := CreateCache(interval)
		cache.Add(c.key, c.value)

		time.Sleep(buffer)

		_, ok := cache.Get(c.key)
		if ok {
			t.Error("Key is accessed when it should be deleted")
			return
		}
	}
}
