package biz

import (
    "context"
    "github.com/go-kratos/kratos/v2/log"
    "gorm.io/gorm"
)

type Download struct {
    gorm.Model
    KEY      string `gorm:"unique; not null; size:256"`
    Location string
    URL      string
}

type DownloadRepo interface {
    GetDownloadURL(ctx context.Context, key, location string) (string, error)
}

type DownloadUseCase struct {
    repo DownloadRepo
    log  *log.Helper
}

func NewDownloadUseCase(repo DownloadRepo, logger log.Logger) *DownloadUseCase {
    return &DownloadUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *DownloadUseCase) GetURL(ctx context.Context, key, location string) (string, error) {
    return uc.repo.GetDownloadURL(ctx, key, location)
}
