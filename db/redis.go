package db

import (
	"github.com/go-redis/redis"
	"os"
)

var instanceDBRedis *redis.Client

func ConnectDatabaseRedis() *redis.Client {

	if instanceDBRedis != nil {
		if _, err := instanceDBRedis.Ping().Result(); err != nil {
			instanceDBRedis = nil
		}
	}

	if instanceDBRedis == nil {
		client := redis.NewClient(&redis.Options{
			Addr:     os.Getenv("DB_REDIS_HOST"),     //"localhost:6379",
			Password: os.Getenv("DB_REDIS_PASSWORD"), //"SUASENHA",       //
			DB:       0,
		})

		if _, err := client.Ping().Result(); err != nil {
			panic(err.Error())
		}

		instanceDBRedis = client
	}

	return instanceDBRedis
}

func RemoveCacheRedisByKey(key string) {
	db := ConnectDatabaseRedis()
	db.Del(key)
}
