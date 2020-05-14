package redis

import (
	"time"

	"github.com/strconv/valyria/config"
)

var _redis *Redis

func Init(c *config.Conf) {
	conf := c.Redis
	_redis = NewRedis(&RedisOpts{
		Host:        conf.Host,
		Password:    conf.Auth,
		Database:    conf.DB,
		MaxIdle:     conf.MaxIdle,
		MaxActive:   conf.MaxActive,
		IdleTimeout: conf.IdleTimeout,
	})
}

func Get(key string) interface{} {
	return _redis.Get(key)
}

func Set(key string, val interface{}, timeout time.Duration) (err error) {
	return _redis.Set(key, val, timeout)
}

func IsExist(key string) bool {
	return _redis.IsExist(key)
}

func Delete(key string) error {
	return _redis.Delete(key)
}
