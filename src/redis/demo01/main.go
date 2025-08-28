package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type BikeInfo struct {
	Model string `redis:"model"`
	Brand string `redis:"brand"`
	Type  string `redis:"type"`
	Price int    `redis:"price"`
}

func main() {

	//设置空白文本流和Redis客户端
	ctx := context.Background()
	rdb2 := redis.NewClient(&redis.Options{
		Addr:     "redis-10059.c325.us-east-1-4.ec2.redns.redis-cloud.com:10059",
		Password: "",
		DB:       0,

		// 连接池配置
		PoolSize:     10,               // 最大连接数
		MinIdleConns: 5,                // 最小空闲连接数
		MaxIdleConns: 10,               // 最大空闲连接数
		PoolTimeout:  30 * time.Second, // 获取连接的超时时间

	})
	defer rdb2.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "120.26.105.128:6379",
		Password: "666666",
		DB:       0,
	})
	defer rdb.Close()

	//string
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("foo", val)

	// hash
	hashFields := []string{
		"model", "demos",
		"brand", "Ergnom",
		"year", "2023",
		"type", "Enduro bikes",
		"price", "4972",
	}
	res1, err := rdb.HSet(ctx, "bike:1", hashFields).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("HSet", res1)

	res2, err := rdb.HGet(ctx, "bike:1", "model").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("HGet", res2)
	res2, err = rdb.HGet(ctx, "bike:1", "brand").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("HGet", res2)
	res2, err = rdb.HGet(ctx, "bike:1", "year").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("HGet", res2)

	res4, err := rdb.HGetAll(ctx, "bike:1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("HGetAll", res4)

	var bike BikeInfo
	err = rdb.HGetAll(ctx, "bike:1").Scan(&bike)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Model: %v Brand: %v Type: %v Price: %v\n", bike.Model, bike.Brand, bike.Type, bike.Price)

}
