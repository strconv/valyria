package redis

import (
	"time"

	"go.uber.org/zap"

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

func HMSet(log *zap.SugaredLogger, key string, fields []string, values []interface{}) error {
	fs := make([]interface{}, len(fields)*2)
	for i := 0; i < len(fields)*2; i += 2 {
		fs[i] = fields[i/2]
		fs[i+1] = values[i/2]
	}

	keys := []interface{}{key}
	keys = append(keys, fields)
	_, err := _redis.Do("HMSET", key, fs)
	if err != nil {
		log.Errorf("redis HMSet fail｜error:%s｜", err)
		return err
	}
	log.Infof("redis HMSet|key:%s|fields:%+v|values:%+v|", key, fields, values)
	return nil
}
