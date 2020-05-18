package writer

import (
  "fmt"
  "github.com/go-redis/redis/v7"
  "strings"
  "time"
)

var client redis.Client

func init() {
  client = redisConn()
}

func redisConn() redis.Client {
  return *redis.NewClient(&redis.Options{
		Addr:     "docker:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func WriteToRedis(path string, hostname string, log string, msg string) {
  defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered from ", r)
    }
  }()

  logname := strings.Split(log, "/")

  _, err := client.Do("lpush", "logs", "[" + time.Now().Format("2006-01-02 15:04:05") + "][" + hostname + "][" + logname[len(logname)-1] + "] " + jsonEscape(msg)).Result()
	if err != nil {
		fmt.Println("Error getting response: ", err)
	}
}
