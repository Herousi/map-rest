package client

import (
	"fmt"
	"strings"
	"time"

	"github.com/Herousi/map-rest/src/common/conf"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

type Redis struct {
	Conn   *redis.Client
	Prefix string
}

var rs Redis

func GetRedisClient() (rcon Redis, e error) {
	if rs.Conn != nil {
		return rs, nil
	}
	addr := conf.Options.RedisURL
	if strings.HasPrefix(addr, "redis://") {
		addr = strings.TrimPrefix(addr, "redis://")
	}
	rcon.Conn = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",                   // no password set
		DB:       conf.Options.RedisDB, // use default DB
	})
	if _, e := rcon.Conn.Ping().Result(); e != nil {
		zap.L().Error("connect redis err", zap.Error(e), zap.String("redis url", addr))
		return rcon, e
	}
	rcon.Prefix = conf.Options.RedisTag
	rs = rcon
	zap.L().Info("connect redis", zap.String("redis url", addr))
	return
}

func (r Redis) HGet(key string, tag string) (str string, err error) {
	key = fmt.Sprintf("%s-%s", strings.ToUpper(r.Prefix), strings.ToUpper(key))
	str, err = r.Conn.HGet(strings.ToUpper(key), strings.ToUpper(tag)).Result()
	return str, err
}

func (r Redis) HSet(key string, tag string, value interface{}) error {
	key = fmt.Sprintf("%s-%s", strings.ToUpper(r.Prefix), strings.ToUpper(key))
	bmd := r.Conn.HSet(strings.ToUpper(key), strings.ToUpper(tag), value)
	return bmd.Err()
}

func (r Redis) HDel(key string, tag string) error {
	key = fmt.Sprintf("%s-%s", strings.ToUpper(r.Prefix), strings.ToUpper(key))
	bmd := r.Conn.HDel(strings.ToUpper(key), strings.ToUpper(tag))
	return bmd.Err()
}

func (r Redis) Get(key string) (str string, err error) {
	key = fmt.Sprintf("%s-%s", strings.ToUpper(r.Prefix), strings.ToUpper(key))
	str, err = r.Conn.Get(strings.ToUpper(key)).Result()
	return str, err
}

func (r Redis) Set(key string, value interface{}, time time.Duration) error {
	key = fmt.Sprintf("%s-%s", strings.ToUpper(r.Prefix), strings.ToUpper(key))
	bmd := r.Conn.Set(strings.ToUpper(key), value, time)
	return bmd.Err()
}

func (r Redis) Del(key string) error {
	key = fmt.Sprintf("%s-%s", strings.ToUpper(r.Prefix), strings.ToUpper(key))
	bmd := r.Conn.Del(strings.ToUpper(key))
	return bmd.Err()
}
