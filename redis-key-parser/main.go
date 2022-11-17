package main

import (
	"github.com/go-redis/redis"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var expiredDay = int64(604800)

func main() {
	count := len(os.Args)
	if count != 5 {
		log.Printf("args error:  %#v, host port password expiredDay\n", os.Args)
		return
	}
	dbHost := os.Args[1]
	dbPort := os.Args[2]
	dbPassword := os.Args[3]
	expiredDay, _ = strconv.ParseInt(os.Args[4], 10, 64)
	expiredDay = expiredDay * 24 * 3600

	options := redis.Options{Addr: dbHost + ":" + dbPort, Password: dbPassword, DB: 0}
	redisClient := redis.NewClient(&options)

	findBigKey(redisClient)
}

func findBigKey(redisClient *redis.Client) {
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
func doFindNoTtlKey(redisClient *redis.Client, key string) {
	ttl, err := redisClient.TTL(key).Result()
	if err != nil {
		return
	}
	if ttl == -1*time.Second {
		if strings.Contains(key, "eyJhbGciOiJIUzI1NiJ9") {
			result, err := redisClient.Get(key).Result()
			if err != nil {
				return
			}
			result = strings.Trim(result, "\"")
			isDelete := false
			nowTime := time.Now().Unix()
			v := int64(0)
			keyCreateTime, err := strconv.ParseInt(result, 10, 64)
			if err == nil {
				v = nowTime - keyCreateTime/1000
				if v >= expiredDay {
					redisClient.Del(key)
					isDelete = true
				}
			}

			log.Println(key, result, nowTime, v, isDelete)
		}
	}
}
