package main

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisStruct struct {
	Tags      string		`php:"tags"`
	Categories string	`php:"categories"`
}

var ctx = context.Background()

func getClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
	//
	//val, err := rdb.Get(ctx, "key").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("key", val)
	//
	//val2, err := rdb.Get(ctx, "key2").Result()
	//if err == redis.Nil {
	//	fmt.Println("key2 does not exist")
	//} else if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("key2", val2)
	//}
	//// Output: key value
	//// key2 does not exist
}

func putSiteTerms(key string, data []string) bool {

	rdb := getClient()
	err := rdb.Set(ctx, key, data, 0).Err()
	if err != nil {
		panic(err)
	}

	return true
}
