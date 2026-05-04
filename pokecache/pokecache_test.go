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

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cases := []struct{
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("value1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("value2"),
		},
		{
			inputKey: "",
			inputVal: []byte("value3"),
		},
	}

	for _, cs := range cases {
		cache.Add(cs.inputKey, cs.inputVal)
		actual, ok := cache.Get(cs.inputKey)

		if !ok {
			t.Errorf("%s not found in cache", cs.inputKey)
			continue
		}

		if string(actual) != string(cs.inputVal) {
			t.Errorf("expected %s, got %s", string(cs.inputVal), string(actual))
			continue
		}	
	}
		

	cache.Add("key1", []byte("value1"))
	actual, ok := cache.Get("key1")

	if !ok {
		t.Error("key1 not found in cache")
	}

	if string(actual) != "value1" {
		t.Errorf("expected value1, got %s", string(actual))
	}
}

func TestReap(t *testing.T) {
	interval := 10 * time.Millisecond
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("value1"))

	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get(keyOne)
	if ok {
		t.Errorf("%s should have been reaped from cache", keyOne)
	}
}

func TestReapFail(t *testing.T) {
	interval := 10 * time.Millisecond
	cache := NewCache(interval)

	keyOne := "key1"
	cache.Add(keyOne, []byte("value1"))

	time.Sleep(interval/2)

	_, ok := cache.Get(keyOne)
	if !ok {
		t.Errorf("%s should not have been reaped from cache", keyOne)
	}
}