package data

import (
    "context"
    "game/internal/biz"
    "game/internal/util/uuid"
    "github.com/go-kratos/kratos/v2/log"
)

type gameRepo struct {
    data *Data
    log  *log.Helper
}

// NewGameRepo .
func NewGameRepo(data *Data, logger log.Logger) biz.GameRepo {
    return &gameRepo{
        data: data,
        log:  log.NewHelper(logger),
    }
}

func (r *gameRepo) CreateGame(ctx context.Context,
    name string, authorId uint, des string, dowId uuid.BinaryUUID) (*biz.Game, error) {
    
    model := &biz.Game{
        Name: name,
        AuthorID: authorId,
        Description: des,
        DownloadID: dowId,
    }

    err := r.data.DataBase.Create(&model).Error
    if err != nil {
        return nil, err
    }

    return model, nil
}
