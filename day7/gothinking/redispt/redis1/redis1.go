package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var redisAddr = "192.168.3.158:6379"

func main() {
	c, err := redis.Dial("tcp", redisAddr)
	if err != nil {
		fmt.Println("connect to redispt error:", err)
		return
	}
	defer c.Close()

	key := "test1"
	_, err = c.Do("SET", key, "super")
	if err != nil {
		fmt.Println("redispt set failed:", err)
	}

	value, err := redis.String(c.Do("GET", key))
	if err != nil {
		fmt.Println("redispt get failed:", err)
	}

	fmt.Printf("get key=%s, value=%s", key, value)

}
