package db

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var Rdbs *redis.Client

func InitRedis(ctx context.Context) {
	rd := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	Rdbs = rd
	//清空当前数据库中的所有key，只要加了这个每次重新启动服务器，所有数据被清空
	// Rdbs.FlushDB(ctx)
	_, err := Rdbs.Ping(ctx).Result() // PING, <nil>
	if err != nil {
		fmt.Println("connect redis failed:", err)
		return
	}

}
func Close() {
	Rdbs.Close()
}
