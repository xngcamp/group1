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

--List 使用场景--

在Web应用中，“列出最新的回复”之类的查询非常普遍，要输出这个顺序却不得不进行排序操作，这通常会带来性能和可扩展性问题。
类似的问题就可以用Redis来解决：
比如，我们的一个Web应用想要列出用户贴出的最新20条评论
假设数据库中的每条评论都有一个唯一的递增的ID字段，每次新评论发表时，我们会将它的ID添加到一个Redis列表


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
	ShowLatestComments("user_1")
	time.Sleep(5 * time.Second)
	ShowLatestComments("user_1")
}

func CreateComments() {
	conn := GetRedisConn()
	defer conn.Close()

	userNum := 10

	// 写入评论数据
	for {
		userId := rand.Intn(userNum)
		account := "user_" + strconv.Itoa(userId)
		comment := "comment at " + time.Now().Format("2006-01-02 15:04:05")
		//fmt.Printf("%s has %s\n", account, comment)
		_, err := conn.Do("LPUSH", account, comment)
		if err != nil {
			log.Fatal(err)
		}

		// 设置列表长度
		_, err = conn.Do("LTRIM", account, 0, 9)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func ShowLatestComments(account string) {
	conn := GetRedisConn()
	defer conn.Close()

	// 通过LRANGE， 取5条记录
	comments, err := redis.Values(conn.Do("LRANGE", account, 0, 4))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%s lastest comments:\n", account)
	for _, comment := range comments {
		fmt.Printf("\t%s\n", comment)
	}
}
