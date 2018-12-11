package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

/*
目标：  读写json到redis

*/

var redisAddr = "192.168.3.158:6379"

func main() {
	c, err := redis.Dial("tcp", redisAddr)
	if err != nil {
		fmt.Println("connect to redispt error:", err)
		return
	}
	defer c.Close()

	//testJson(c)
	testList(c)
}

func testJson(c redis.Conn) {
	key := "testjson"

	data := map[string]string{"username": "tom", "mobile": "13912345678"}

	value, _ := json.Marshal(data)

	n, err := c.Do("SETNX", key, value)
	if err != nil {
		fmt.Println("redispt set failed:", err)
	}

	if n == int64(1) {
		fmt.Println("set success")
	}

	var dataGet map[string]string
	valueGet, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		fmt.Printf("redispt get %s failed:%s\n", key, err)
	}

	errShal := json.Unmarshal(valueGet, &dataGet)
	if errShal != nil {
		fmt.Println(errShal)
	}

	fmt.Printf("get key=%s, value=%s\n", key, dataGet)

	fmt.Printf("username:%s, ", dataGet["username"])
	fmt.Printf("mobile:%s\n", dataGet["mobile"])

}

func testList(c redis.Conn) {
	key := "testlist"

	_, err := c.Do("LPUSH", key, "redispt")
	if err != nil {
		fmt.Println("redispt set failed:", err)
	}

	_, err = c.Do("LPUSH", key, "mongodb")
	if err != nil {
		fmt.Println("redispt set failed:", err)
	}

	_, err = c.Do("LPUSH", key, "mysql")
	if err != nil {
		fmt.Println("redispt set failed:", err)
	}

	values, err := redis.Values(c.Do("LRANGE", key, "0", "10"))
	if err != nil {
		fmt.Printf("redispt get %s failed:%s\n", key, err)
	}

	for _, v := range values {
		fmt.Println(string(v.([]byte)))
	}
}
