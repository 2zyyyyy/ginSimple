package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

// go-redis基本使用

// 声明全局变量
var rdb *redis.Client

// 初始化链接
func initClient() error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	p, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println("rdb.Ping failed, err:%", err)
		return err
	}
	fmt.Println("rdb.Ping().Result():", p)
	return nil
}

// 连接redis哨兵模式
func initSentry() (err error) {
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master",
		SentinelAddrs: []string{"x.x.x.x:26379", "xx.xx.xx.xx:26379", "xxx.xxx.xxx.xxx:26379"},
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

// 连接redis集群
func initCluster() (err error) {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

// set/get
func redisSetGet() error {
	if err := rdb.Set("name", "万里", 0).Err(); err != nil {
		fmt.Printf("set name failed, err:%v\n", err)
		return err
	}
	// get
	value, err := rdb.Get("name").Result()
	if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return err
	}
	fmt.Println("user:", value)

	// 获取不到场景
	value, err = rdb.Get("info").Result()
	if err == redis.Nil {
		fmt.Printf("info dose not exist, err:%v\n", err)
	} else if err != nil {
		fmt.Printf("get info failed, err:%v\n", err)
		return err
	} else {
		fmt.Println("info:", value)
	}
	return nil
}

// zset
func redisZset() {
	zsetKey := "language_rank"
	languages := []redis.Z{
		redis.Z{Score: 90.0, Member: "Golang"},
		redis.Z{Score: 98.0, Member: "Java"},
		redis.Z{Score: 95.0, Member: "Python"},
		redis.Z{Score: 97.0, Member: "JavaScript"},
		redis.Z{Score: 99.0, Member: "C/C++"},
	}
	// ZADD
	num, err := rdb.ZAdd(zsetKey, languages...).Result()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Println("zadd success.", num)

	// 将Go的分数加10
	newScore, err := rdb.ZIncrBy(zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincyby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang new score:%f\n", newScore)

	// 获取分数前三
	ret, err := rdb.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrange failed, err:%v\n", err)
		return
	}
	for _, v := range ret {
		fmt.Println("获取分数前三:", v.Member, v.Score)
	}

	// 取95~100
	op := redis.ZRangeBy{
		Min: "98",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangbyscore failed, err:%v\n", err)
		return
	}
	for _, v := range ret {
		fmt.Println("取98~100", v.Member, v.Score)
	}
}

func main() {
	defer func(rdb *redis.Client) {
		err := rdb.Close()
		if err != nil {
			fmt.Println("rdb close failed, err:", err)
			return
		}
	}(rdb)
	err := initClient()
	if err != nil {
		fmt.Println("initClient() failed, err:%", err)
		return
	}

	err = redisSetGet()
	if err != nil {
		fmt.Println("redisSetGet() failed, err:%", err)
		return
	}

	redisZset()
}
