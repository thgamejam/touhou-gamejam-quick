package data

import (
    "download/internal/conf"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/go-redis/redis/v8"
    "github.com/google/wire"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewMySQL, NewRedis)

// Data .
type Data struct {
    // 封装的数据库客户端
    DataBase *gorm.DB
    Cache    *redis.Client
}

// NewData .
func NewData(c *conf.Data, db *gorm.DB, red *redis.Client, logger log.Logger) (*Data, func(), error) {
    cleanup := func() {
        log.NewHelper(logger).Info("closing the data resources")
    }
    return &Data{
        DataBase: db,
        Cache:    red,
    }, cleanup, nil
}

// NewRedis 初始化缓存
func NewRedis(c *conf.Data) (*redis.Client, error) {
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

// NewMySQL 初始化数据库
func NewMySQL(c *conf.Data, logger log.Logger) (*gorm.DB, error) {
    db, err := gorm.Open(
        mysql.Open(c.Database.Dsn),
        &gorm.Config{
            //Logger: , // TODO 绑定log 未完成
        })
    if err != nil {
        return nil, err
    }

    selDb, err := db.DB()
    if err != nil {
        return nil, err
    }
    // 设置连接池
    // 空闲
    selDb.SetMaxIdleConns(int(c.Database.MaxIdleConn))
    // 打开
    selDb.SetMaxOpenConns(int(c.Database.MaxOpenConn))
    // 超时 time.Second * 30
    selDb.SetConnMaxLifetime(c.Database.ConnMaxLifetime.AsDuration())

    //err = db.AutoMigrate(&model.User{})

    if err != nil {
        return nil, err
    }

    return db, nil
}

func SQLAutoMigrate() {

}
