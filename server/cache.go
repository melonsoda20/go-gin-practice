package server

import (
	"fmt"
	"gin-todo-app/models"
	"gin-todo-app/services"

	"github.com/gomodule/redigo/redis"
)

const maxConnections = 10

func ConnectRedis() (*redis.Pool, error) {
	configPath := "../config/redis.json"

	configFile, file_err := services.OpenFile(configPath)
	if file_err != nil {
		return nil, file_err
	}
	defer configFile.Close()

	configData, data_err := services.ReadFile(configFile)
	if data_err != nil {
		return nil, data_err
	}

	redisConfig := models.RedisConfig{}
	services.DeserializeFile(configData, &redisConfig)

	fmt.Printf("%#v", redisConfig)

	redisPool := &redis.Pool{
		MaxIdle: maxConnections,
		Dial:    func() (redis.Conn, error) { return redis.Dial("tcp", redisConfig.URL) },
	}

	return redisPool, nil
}
