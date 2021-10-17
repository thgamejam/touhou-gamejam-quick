package data

import (
    "download/internal/conf"
    "github.com/go-redis/redis/v8"
)

// NewCache 初始化缓存
func NewCache(c *conf.Data) (*redis.Client, error) {
    rdb := redis.NewClient(&redis.Options{
        Network:      c.Redis.Network,
        Addr:         c.Redis.Addr,
        Password:     c.Redis.Password, // no password set
        ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
        WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
        DB:           0, // use default DB
    })
    return rdb, nil
}
