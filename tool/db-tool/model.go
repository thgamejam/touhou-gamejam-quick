package main

import "gorm.io/gorm"

// Game "游戏"模型
type Game struct {
    gorm.Model

    Name        string     `gorm:"not null; type:varchar(32)"`
    AuthorID    uint       `gorm:"index; not null"`             // 作者id
    Description string     `gorm:"not null; type:varchar(512)"` // 简介
    DownloadID  BinaryUUID `gorm:"not null; type:binary(16)"`   // 下载文件uuid
}

// GameTags 游戏标签模型
// 存放游戏标签的属性
type GameTags struct {
    gorm.Model

    Name string `gorm:"not null; type:varchar(16)"`
}

// GameExistTags 游戏存在的标签
// 游戏和标签的对应表
type GameExistTags struct {
    gorm.Model

    GameID uint `gorm:"index; not null"`
    TagsID uint `gorm:"not null"`
}

// GameExistImgs 游戏存在的介绍图片
type GameExistImgs struct {
    gorm.Model

    GameID uint       `gorm:"index; not null"`
    ImgID  BinaryUUID `gorm:"not null; type:binary(16)"` // 介绍图片文件uuid
}
