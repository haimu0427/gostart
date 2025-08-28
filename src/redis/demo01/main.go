package main

import (
	"context"
	"fmt"

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
