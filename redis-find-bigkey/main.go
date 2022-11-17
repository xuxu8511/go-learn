package main

import (
	"github.com/go-redis/redis"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	count := len(os.Args)
	if count != 5 {
		log.Printf("args error:  %#v, host port password threvalue\n", os.Args)
		return
	}
	dbHost := os.Args[1]
	dbPort := os.Args[2]
	dbPassword := os.Args[3]
	threValue, _ := strconv.Atoi(os.Args[4])

	options := redis.Options{Addr: dbHost + ":" + dbPort, Password: dbPassword, DB: 0}
	redisClient := redis.NewClient(&options)

	findBigKey(redisClient, threValue)
}

func findBigKey(redisClient *redis.Client, threValue int) {
	var cursor uint64
	var err error
	for {
		var keys []string
		keys, cursor, err = redisClient.Scan(cursor, "", 1000).Result()
		if err != nil {
			panic(err)
		}

		for _, key := range keys {
			doFindNoTtlKey(redisClient, key)
		}

		if cursor == 0 {
			break
		}
	}
}

func doFindBigKey(redisClient *redis.Client, key string, threValue int) {
	length, err := redisClient.MemoryUsage(key).Result()
	if err != nil {
		return
	}
	if int(length) > threValue {
		log.Println(key, length)
	}
}

func doFindNoTtlKey(redisClient *redis.Client, key string) {
	ttl, err := redisClient.TTL(key).Result()
	if err != nil {
		return
	}
	if ttl == -1*time.Second {
		log.Println(key)
	}
}
