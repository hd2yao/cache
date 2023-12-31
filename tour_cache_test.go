package cache

import (
    "log"
    "sync"
    "testing"

    "github.com/matryer/is"

    "github.com/hd2yao/cache/lru"
)

func TestTourCacheGet(t *testing.T) {
    db := map[string]string{
        "key1": "val1",
        "key2": "val2",
        "key3": "val3",
        "key4": "val4",
    }
    getter := GetFunc(func(key string) interface{} {
        log.Println("[From DB] find key", key)

        if val, ok := db[key]; ok {
            return val
        }
        return nil
    })
    tourCache := NewTourCache(getter, lru.New(0, nil))

    is := is.New(t)

    var wg sync.WaitGroup

    for k, v := range db {
        wg.Add(1)
        go func(k, v string) {
            defer wg.Done()
            is.Equal(tourCache.Get(k), v)

            is.Equal(tourCache.Get(k), v)
        }(k, v)
    }
    wg.Wait()

    is.Equal(tourCache.Get("unknown"), nil)
    is.Equal(tourCache.Get("unknown"), nil)

    is.Equal(tourCache.Stat().NGet, 10)
    is.Equal(tourCache.Stat().NHit, 4)
}
