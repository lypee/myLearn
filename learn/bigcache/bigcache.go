package main

import (
	"github.com/allegro/bigcache"
	"time"
)

type MyCache struct{
	bc *bigcache.BigCache
}
func NewMyCache() MyCache{
	return MyCache{
		bc: NewBigCache(),
	}
}

func main()  {

}

func NewBigCache()*bigcache.BigCache{
	bc , _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Second))
	return bc
}

func (m *MyCache)MyGet1(key string){
	m.bc.Get(key)
}
