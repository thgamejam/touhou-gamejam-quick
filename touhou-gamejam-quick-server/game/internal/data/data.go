package data

import (
    "download/internal/conf"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/go-redis/redis/v8"
    "github.com/google/wire"
    "gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDataBase, NewCache)

// Data .
type Data struct {
    // 封装的数据库客户端
    DataBase      *gorm.DB       // sql数据库
    Cache         *redis.Client  // 缓存
}

// NewData .
func NewData(c *conf.Data,
    db *gorm.DB, red *redis.Client,
    logger log.Logger) (*Data, func(), error) {

    cleanup := func() {
        log.NewHelper(logger).Info("closing the data resources")
    }
    return &Data{
        DataBase:      db,
        Cache:         red,
    }, cleanup, nil
}
