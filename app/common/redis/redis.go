package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/king-tu/mallweb/app/global"
	"time"
)

var redisConn *redis.Pool

// InitRedis Initialize the Redis instance
func init() {
	redisConn = &redis.Pool{
		MaxIdle:     global.Config.Redis.MaxIdle,
		MaxActive:   global.Config.Redis.MaxActive,
		IdleTimeout: time.Second * time.Duration(global.Config.Redis.IdleTimeout),
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", global.Config.Redis.Addr)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// Set a key/value
func Set(key string, value interface{}) error {
	conn := redisConn.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	return nil
}

func Setex(key string, value interface{}, time int) error {
	conn := redisConn.Get()
	defer conn.Close()

	_, err := conn.Do("setex", key, time, value)
	if err != nil {
		return err
	}

	return nil
}

// Exists check a key
func Exists(key string) bool {
	conn := redisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

// Get get a key
func Get(key string) (string, error) {
	conn := redisConn.Get()
	defer conn.Close()

	reply, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}

	return reply, nil
}

// Delete delete a kye
func Delete(key string) (bool, error) {
	conn := redisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

// LikeDeletes batch delete
func LikeDeletes(key string) error {
	conn := redisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}
