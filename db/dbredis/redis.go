package dbredis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

type RedisConf struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	DefaultDb   int
	IdleTimeout time.Duration
}

var (
	cfg       = &RedisConf{}
	redisPool *redis.Pool
)

func GetRedis() redis.Conn {
	return redisPool.Get()
}

func newpool() (*redis.Pool, error) {
	pool := &redis.Pool{
		MaxIdle:     cfg.MaxIdle,
		MaxActive:   cfg.MaxActive,
		IdleTimeout: cfg.IdleTimeout * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", cfg.Host)
			if err != nil {
				return nil, err
			}
			if cfg.Password != "" {
				if _, err := c.Do("AUTH", cfg.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			if _, err := c.Do("SELECT", cfg.DefaultDb); err != nil {
				c.Close()
				return nil, err
			}

			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
	return pool, nil
}

func SetUp(c *RedisConf) {
	cfg = c
	var err error
	redisPool, err = newpool()
	if err != nil {
		panic(err.Error())
	}
}
