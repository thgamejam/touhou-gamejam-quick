package data

import (
    "game/internal/biz"
    "game/internal/conf"
    "github.com/go-kratos/kratos/v2/log"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

// NewDataBase 初始化数据库
func NewDataBase(c *conf.Data, logger log.Logger) (*gorm.DB, error) {
    db, err := gorm.Open(
        mysql.Open(c.DataBase.Source),
        &gorm.Config{
            //Logger: , // TODO 绑定 Log 未完成
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
    selDb.SetMaxIdleConns(int(c.DataBase.MaxIdleConn))
    // 打开
    selDb.SetMaxOpenConns(int(c.DataBase.MaxOpenConn))
    // 超时 time.Second * 30
    selDb.SetConnMaxLifetime(c.DataBase.ConnMaxLifetime.AsDuration())

    err = DBAutoMigrate(db)
    if err != nil {
        return nil, err
    }

    return db, nil
}

// DBAutoMigrate .
func DBAutoMigrate(db *gorm.DB) error {
    err := db.AutoMigrate(
        &biz.Game{},
        &biz.GameTags{},
        &biz.GameExistTags{},
        &biz.GameExistImgs{},
    )
    if err != nil {
        return err
    }

    return nil
}
