package redisManager

import (
	"context"
	"fmt"
	redis "github.com/go-redis/redis"
	redisv "github.com/go-redis/redis/v8" // 注意导入的是新版本
	"go_demo/configManager"
	"time"
)

var rdb_v8 *redisv.Client

func initClient_v8()(err error) {
	redisConfig := configManager.GetRedisConfig()
	redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host + ":" + redisConfig.Port,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
		PoolSize: 100, // 连接池大小
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = rdb_v8.Ping(ctx).Result()
	return err
}

func V8Example() {
	ctx := context.Background()
	if err := initClient_v8(); err != nil {
		return
	}

	err := rdb_v8.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb_v8.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb_v8.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
