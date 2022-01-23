package db

import (
	"github.com/go-redis/redis"
)

var instanceDBRedis *redis.Client

// ConnectDatabase - get connection database
func ConnectDatabaseRedis() *redis.Client {
	//dbPassword := os.Getenv("DB_REDIS_PASSWORD")
	//dbHost := os.Getenv("DB_REDIS_HOST")
	// dbPort := os.Getenv("DB_REDIS_PORT")

	if instanceDBRedis != nil {
		if _, err := instanceDBRedis.Ping().Result(); err != nil {
			instanceDBRedis = nil
		}
	}

	if instanceDBRedis == nil {
		client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "SUASENHA",
			DB:       0,
		})

		if _, err := client.Ping().Result(); err != nil {
			panic(err.Error())
		}

		instanceDBRedis = client
	}

	return instanceDBRedis
}
