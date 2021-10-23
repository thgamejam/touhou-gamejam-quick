package main

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "time"
)

// NewDataBase 初始化数据库
func NewDataBase() *gorm.DB {
    db, err := gorm.Open(
        mysql.Open("root:206242227@tcp(test.io:3306)/touhou_gamejam?charset=utf8mb4&parseTime=True&loc=Local"),
        &gorm.Config{
            //Logger: , // TODO 绑定 Log 未完成
        })
    if err != nil {
        return nil
    }

    selDb, err := db.DB()
    if err != nil {
        return nil
    }
    // 设置连接池
    // 空闲
    selDb.SetMaxIdleConns(50)
    // 打开
    selDb.SetMaxOpenConns(100)
    // 超时 time.Second * 30
    selDb.SetConnMaxLifetime(time.Second * 30)

    err = DBAutoMigrate(db)
    if err != nil {
        return nil
    }

    return db
}

// DBAutoMigrate .
func DBAutoMigrate(db *gorm.DB) error {
    err := db.AutoMigrate(
        &Game{},
        &GameTags{},
        &GameExistTags{},
        &GameExistImgs{},
    )
    if err != nil {
        return err
    }

    return nil
}
