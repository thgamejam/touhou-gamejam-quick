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

func (r *downloadRepo) GetDownloadURL(ctx context.Context, key, location string) (string, error) {
    model := &biz.Download{}
    r.data.DataBase.Where("key = ?", key).First(model)

    return model.URL, nil
}
