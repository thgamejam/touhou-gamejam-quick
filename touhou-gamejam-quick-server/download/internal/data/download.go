package data

import (
    "context"
    "download/internal/biz"
    "github.com/go-kratos/kratos/v2/log"
)

type downloadRepo struct {
    data *Data
    log  *log.Helper
}

// NewDownloadRepo .
func NewDownloadRepo(data *Data, logger log.Logger) biz.DownloadRepo {
    return &downloadRepo{
        data: data,
        log:  log.NewHelper(logger),
    }
}

func (d *downloadRepo) GetGame(ctx context.Context, uuid, region string) (string, error) {
    // 从数据库中查找是否存在对应的uuid内容
    model := &biz.Download{}
    d.data.DataBase.Where("uuid = ?", uuid).First(model)

    return "", nil
}
