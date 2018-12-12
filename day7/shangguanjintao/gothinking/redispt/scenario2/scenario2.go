package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"math/rand"
	"strconv"
	"time"
)

/*

场景：排行榜应用，取TOP N操作
使用：Redis Sorted Set, 有序集合


最新N个数据是以某个条件为权重，比如按点赞的次数排序，这时候就需要sorted set
将你要排序的值设置成sorted set的score，将具体的数据设置成相应的value，每次只需要执行一条ZADD命令即可

重点：
1. 如何设计score值


*/

var pool *redis.Pool

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
	go CreateComments()
	ShowComments("user_1")
	time.Sleep(5 * time.Second)
	ShowComments("user_1")
}

func CreateComments() {
	conn := GetRedisConn()
	defer conn.Close()

	userNum := 10

	// 写入评论数据
	for {
		userId := rand.Intn(userNum)
		account := "user_" + strconv.Itoa(userId)

		//点赞次数
		thumbUpNum := rand.Intn(1000)
		comment := fmt.Sprintf("comment at %s, thumb-up num:%d",
			time.Now().Format("2006-01-02 15:04:05"), thumbUpNum)

		//fmt.Printf("%s has %s\n", account, comment)
		_, err := conn.Do("ZADD", account, thumbUpNum, comment)
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(200 * time.Millisecond)
	}
}

func ShowComments(account string) {
	conn := GetRedisConn()
	defer conn.Close()

	// ZREVRANGE， 按照score值，从大到小排序，取5条记录
	comments, err := redis.Values(conn.Do("ZREVRANGE", account, 0, 4))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%s lastest comments:\n", account)
	for _, comment := range comments {
		fmt.Printf("\t%s\n", comment)
	}
}
