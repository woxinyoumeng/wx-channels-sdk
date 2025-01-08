package demo

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/cache"
	"github.com/woxinyoumeng/wx-channels-sdk/apis"
	"log"
	"time"
)

var (
	ca = cache.NewMemoryCache()
)

// 实现 access_token 的 内存 存取方案
type DcsTokenByMem struct{}

func (w DcsTokenByMem) Get(cacheKey string) apis.TokenInfo {
	var tokenInfo apis.TokenInfo
	get, err := ca.Get(ctx, cacheKey)
	if err != nil {
		log.Println("cache get  err", err)
		return tokenInfo
	}
	if get != nil {
		data := get.(string)
		err = json.Unmarshal([]byte(data), &tokenInfo)
		if err != nil {
			log.Println("cache set err", err)
			return tokenInfo
		}
	}
	return tokenInfo
}

func (w DcsTokenByMem) Set(cacheKey string, tokenInfo apis.TokenInfo, ttl time.Duration) error {
	marshal, _ := json.Marshal(tokenInfo)
	return ca.Put(ctx, cacheKey, string(marshal), ttl)
}

func (w DcsTokenByMem) Del(cacheKey string) error {
	return ca.Delete(ctx, cacheKey)
}

func (w DcsTokenByMem) Lock(cacheKey string, ttl time.Duration) bool {
	err := ca.Put(ctx, cacheKey, "1", ttl)
	if err != nil {
		log.Println("cache lock err", err)
		return false
	}
	return true
}

func (w DcsTokenByMem) Unlock(cacheKey string) error {
	return ca.Delete(ctx, cacheKey)
}
