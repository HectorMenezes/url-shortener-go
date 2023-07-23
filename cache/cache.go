package cache

import (
    "log"
    "time"
    "encoding/json"

	"github.com/HectorMenezes/url-shortener-go/models"
    "github.com/patrickmn/go-cache"
)

type allCache struct{
    urls *cache.Cache
}

var c *allCache

const (
    defaultExpiration = 5 * time.Minute
    purgeTime = 10 & time.Minute
)

func newCache() *allCache {
    Cache := cache.New(defaultExpiration, purgeTime)
    return &allCache{
        urls: Cache,
    }
}

func (c *allCache) Read(id string) (item []byte, ok bool){
    url, ok := c.urls.Get(id)
    if ok {
        log.Println("From cache")
        res, err := json.Marshal(url.(models.Url))
        if err != nil {
            log.Fatalln("Error")
        }
        return res, true
    }
    return nil, false
}

func (c *allCache) Update(id string, url models.Url){
    c.urls.Set(id, url, cache.DefaultExpiration)
}

func Start() {
    c = newCache()
}
func GetCache() *allCache {
    return c
}
