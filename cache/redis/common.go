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

func Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	return _redis.Do(commandName, args...)
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

func HDel(key interface{}, fields ...interface{}) (res int, err error) {
	var reply interface{}
	keys := []interface{}{key}
	keys = append(keys, fields...)

	reply, err = _redis.Do("HDEL", keys)
	if err != nil {
		return
	}
	res = reply.(int)
	return
}

func HSet(key, field string, value interface{}) (res int, err error) {
	var reply interface{}
	reply, err = _redis.Do("HSET", key, field, value)
	if err != nil {
		return
	}
	res = reply.(int)
	return
}

func HGet(key, field string) (res string, err error) {
	var reply interface{}
	reply, err = _redis.Do("HGET", key, field)
	if err != nil {
		return
	}
	res = reply.(string)
	return
}

func HGetInt(key, field string) (res int, err error) {
	var reply interface{}
	reply, err = _redis.Do("HGET", key, field)
	if err != nil {
		return
	}
	res = reply.(int)
	return
}

func HMSet(key string, fields ...interface{}) (res string, err error) {
	var reply interface{}
	keys := []interface{}{key}
	keys = append(keys, fields)
	reply, err = _redis.Do("HMSET", keys)
	res = reply.(string)
	return
}
