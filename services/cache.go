package services

import (
	"gin-todo-app/constants"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

func GetRedis(c *gin.Context) (redis.Pool, bool) {
	value, isExists := c.Get(constants.MiddlewareKeysConst.RedisAppKey)
	redis := value.(redis.Pool)

	return redis, isExists
}

func GetCacheData(redis redis.Pool, k string) (interface{}, error) {
	conn := redis.Get()
	if conn.Err() != nil {
		return nil, conn.Err()
	}
	defer conn.Close()

	data, err := conn.Do("GET", k)
	if err != nil {
		return nil, err
	}
	return data, err
}

func SetCacheData(redis redis.Pool, k string, v string) error {
	conn := redis.Get()
	if conn.Err() != nil {
		return conn.Err()
	}
	defer conn.Close()

	_, err := conn.Do("SET", k, v)
	return err
}

func CheckCacheExists(redis redis.Pool, k string) (interface{}, error) {
	conn := redis.Get()
	if conn.Err() != nil {
		return nil, conn.Err()
	}
	defer conn.Close()

	data, err := conn.Do("EXISTS", k)
	if err != nil {
		return nil, err
	}
	return data, err
}
