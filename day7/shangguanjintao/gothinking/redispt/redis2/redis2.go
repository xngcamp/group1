package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

/*

1. Refactor: 从main中抽取出函数
2.
*/

var redisAddr = "192.168.3.158:6379"

func main() {
	c, err := redis.Dial("tcp", redisAddr)
	if err != nil {
		fmt.Println("connect to redispt error:", err)
		return
	}
	defer c.Close()

	testHash(c)
	testHash2(c)
}

func testHash(c redis.Conn) {
	key := "test1"
	_, err := c.Do("SET", key, "super")
	if err != nil {
		fmt.Println("redispt set failed:", err)
	}
	value, err := redis.String(c.Do("GET", key))
	if err != nil {
		fmt.Println("redispt get failed:", err)
	}
	fmt.Printf("get key=%s, value=%s\n", key, value)
}

func testHash2(c redis.Conn) {
	key := "test2"

	// EX设置过期时间
	_, err := c.Do("SET", key, "super", "EX", "5")
	if err != nil {
		fmt.Println("redispt set failed:", err)
	}
	value, err := redis.String(c.Do("GET", key))
	if err != nil {
		fmt.Printf("redispt get %s failed:%s\n", key, err)
	} else {
		fmt.Printf("get key=%s, value=%s\n", key, value)
	}

	fmt.Println("sleep ...")
	time.Sleep(6 * time.Second)

	value, err = redis.String(c.Do("GET", key))
	if err != nil {
		fmt.Printf("redispt get %s failed:%s\n", key, err)
	} else {
		fmt.Printf("get key=%s, value=%s\n", key, value)
	}

}
