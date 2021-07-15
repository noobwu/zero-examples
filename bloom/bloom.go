package main

import (
	"fmt"

	"github.com/tal-tech/go-zero/core/bloom"
	"github.com/tal-tech/go-zero/core/stores/redis"
)

var (
	rdx     = flag.String("redis", "localhost:6379", "the redis, default localhost:6379")
	rdxType = flag.String("redisType", "node", "the redis type, default node")
	rdxPass = flag.String("redisPass", "", "the redis password")
	rdxKey  = flag.String("redisKey", "testbloom", "the redis key, default testbloom")
	threads = flag.Int("threads", runtime.NumCPU(), "the concurrent threads, default to cores")
)
func main() {
	store := redis.NewRedis(*rdx, *rdxType,*rdxPass)
	filter := bloom.New(store, *rdxKey, 64)
	filter.Add([]byte("kevin"))
	filter.Add([]byte("wan"))
	fmt.Println(filter.Exists([]byte("kevin")))
	fmt.Println(filter.Exists([]byte("wan")))
	fmt.Println(filter.Exists([]byte("nothing")))
}
