package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool

/*

 */
func init() {
	redisAddress := "192.168.3.158:6379"
	poolSize := 20
	pool = &redis.Pool{
		MaxIdle:     poolSize,
		IdleTimeout: time.Minute,
		Dial: func() (conn redis.Conn, err error) {
			conn, err = redis.Dial("tcp", redisAddress)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
}

func GetRedisConn() redis.Conn {
	return pool.Get()
}

func main() {
	testRedis()
}

func testRedis() {
	conn := GetRedisConn()
	fmt.Printf("conn=%p, %v\n", conn, conn)

	//记得销毁本次链连接
	defer conn.Close()

	key := "key1000"

	//写入数据
	_, err := conn.Do("SET", key, "redigo")
	if err != nil {
		fmt.Println("error while setting")
	}

	//判断key是否存在
	is_key_exit, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		fmt.Println("error while existing")
	}
	fmt.Printf("key:%s exist %b", key, is_key_exit)

	//获取value并转成字符串
	account_balance, err := redis.String(conn.Do("GET", key))
	if err != nil {
		fmt.Println("error while getting")
	}
	fmt.Printf("account_balance=%s", account_balance)

	//删除key
	_, err = conn.Do("DEL", "go_key")
	if err != nil {
		fmt.Println("error while deleting")
	}

	//设置key过期时间
	_, err = conn.Do("SET", "mykey", "superWang", "EX", "5")
	if err != nil {
		fmt.Println("redispt set failed:", err)
	}

	//创建key时设置5s过期
	_, err = conn.Do("SET", "go_key:ex", "redigo", "EX", 5)
	if err != nil {
		fmt.Println("error while setting")
	}

	//对已有key设置5s过期时间
	n, err := conn.Do("EXPIRE", "go_key", 5)
	if err != nil {
		fmt.Println("error while setting")
	} else if n != int64(1) {
		fmt.Println("failed")
	}
}
