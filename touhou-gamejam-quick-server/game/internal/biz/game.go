package biz

import (
    "context"
    "github.com/go-kratos/kratos/v2/log"
    "gorm.io/gorm"
)

// Game "游戏"模型
type Game struct {
    gorm.Model

    Name        string `gorm:"not null; type:varchar(32)"`
    AuthorID    uint   `gorm:"index; not null"`             // 作者id
    Description string `gorm:"not null; type:varchar(512)"` // 简介
    DownloadID  string `gorm:"not null; type:binary(16)"`   // 下载文件uuid
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

    GameID uint   `gorm:"index; not null"`
    ImgID  string `gorm:"not null; type:binary(16)"` // 介绍图片文件uuid
}

type GameRepo interface {
    // GetGame 通过id获取游戏页面数据
    GetGame(ctx context.Context, id uint) (*Game, error)
    // GameTags 通过id获取游戏标签
    GameTags(ctx context.Context, id uint) ([]string, error)
    // GameImgs 通过id获取游戏图片url
    GameImgs(ctx context.Context, id uint) ([]string, error)
}

type GameUseCase struct {
    repo GameRepo
    log  *log.Helper
}

func NewGameUseCase(repo GameRepo, logger log.Logger) *GameUseCase {
    return &GameUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *GameUseCase) GetGame(ctx context.Context, id uint) (*Game, error) {
    return uc.repo.GetGame(ctx, id)
}

func (uc *GameUseCase) GameTags(ctx context.Context, id uint) ([]string, error) {
    return uc.repo.GameTags(ctx, id)
}

func (uc *GameUseCase) GameImgs(ctx context.Context, id uint) ([]string, error) {
    return uc.repo.GameImgs(ctx, id)
}
