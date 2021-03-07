package redisManager

import (
	_ "context"
	"fmt"
	_ "fmt"
	redis "github.com/go-redis/redis"
	"go_demo/configManager"
	_ "time"
)

var rdb *redis.Client

func initClient()(err error) {
	redisConfig := configManager.GetRedisConfig()
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host + ":" + redisConfig.Port,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
		PoolSize: 100, // 连接池大小
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}


//连接Redis哨兵模式
func initClientSentinel()(err error) {
	rdb = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master",
		SentinelAddrs: []string{"x.x.x.x:26379", "xx.xx.xx.xx:26379", "xxx.xxx.xxx.xxx:26379"},
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

//连接Redis集群
func initClientCluster()(err error){
	rdb  := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func Set(key string,value interface{}) {
	err := rdb.Set(key, value, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
	}
}

func Get(key string) interface{}{
	value,err:=rdb.Get(key).Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
	}
	return value
}