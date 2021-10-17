package data

import (
    "download/internal/conf"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/go-redis/redis/v8"
    "github.com/google/wire"
    "gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDataBase, NewCache, NewObjectStorage, NewDownloadRepo)

// Data .
type Data struct {
    // 封装的数据库客户端
    DataBase      *gorm.DB       // sql数据库
    Cache         *redis.Client  // 缓存
    ObjectStorage *ObjectStorage // 对象存储服务的封装
}

// NewData .
func NewData(c *conf.Data,
    db *gorm.DB, red *redis.Client, oss *ObjectStorage,
    logger log.Logger) (*Data, func(), error) {

    cleanup := func() {
        log.NewHelper(logger).Info("closing the data resources")
    }
    return &Data{
        DataBase:      db,
        Cache:         red,
        ObjectStorage: oss,
    }, cleanup, nil
}
