package biz

import (
    "context"
    "github.com/go-kratos/kratos/v2/log"
    "gorm.io/gorm"
)

type Download struct {
    gorm.Model
    UUID   string `gorm:"unique; not null; type:char(16)"` // TODO UUID size not ✔
    Region string `gorm:"not null; type:char(16)"`
}

type DownloadRepo interface {
    GetGame(ctx context.Context, key, location string) (string, error)
}

type DownloadUseCase struct {
    repo DownloadRepo
    log  *log.Helper
}

func NewDownloadUseCase(repo DownloadRepo, logger log.Logger) *DownloadUseCase {
    return &DownloadUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *DownloadUseCase) GetURL(ctx context.Context, uuid, region string) (string, error) {
    return uc.repo.GetGame(ctx, uuid, region)
}
