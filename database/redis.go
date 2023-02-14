package database

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/kennedybaraka/app-server/config"
)

var (
	REDISHOST     = config.Env("REDISHOST")
	REDISPASSWORD = config.Env("REDISPASSWORD")
	REDISPORT     = config.Env("REDISPORT")
	REDISUSER     = config.Env("REDISUSER")
	Client        *redis.Client
)

func init() {
	// database stuff
	ConnectToRedis()
}

func ConnectToRedis() {

	redisAddr := fmt.Sprintf("%s:%s", REDISHOST, REDISPORT)
	Client = redis.NewClient(&redis.Options{

		Addr:     redisAddr,
		Password: REDISPASSWORD,
		DB:       0,
	})
	_, err := Client.Ping(Client.Context()).Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[#] REDIS CONNECTION HAS BEEN ESTABLISHED.")

}
