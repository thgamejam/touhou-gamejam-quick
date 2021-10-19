package data

import (
    "context"
    "game/internal/biz"
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

func (r *gameRepo) GetGame(ctx context.Context, id uint) (*biz.Game, error) {
    model := &biz.Game{}
    err := r.data.DataBase.First(&model, id).Error
    if err != nil {
        return nil, err
    }
    return model, nil
}

func (r *gameRepo) GameTags(ctx context.Context, id uint) ([]string, error) {
    // 数据库获取game所有的标签id
    var gameExistTagsModels []biz.GameExistTags
    err := r.data.DataBase.Where("game_id = ?", id).Find(&gameExistTagsModels).Error
    if err != nil {
        return nil, err
    }
    tagIdArray  := make([]uint, 0, len(gameExistTagsModels))
    for _, gameModel := range gameExistTagsModels {
        tagIdArray = append(tagIdArray, gameModel.TagsID)
    }

    // 数据库获取标签的对应名字
    var tagsModels []biz.GameTags
    err = r.data.DataBase.Find(&tagsModels, []int{1, 2, 3}).Error
    if err != nil {
        return nil, err
    }
    tagNameArray := make([]string, 0, len(tagsModels))
    for _, tagModel := range tagsModels {
        tagNameArray = append(tagNameArray, tagModel.Name)
    }

    return tagNameArray, nil
}

func (r *gameRepo) GameImgs(ctx context.Context, id uint) ([]string, error) {
    // 数据库获取game所有的介绍图片资源uuid
    var gameExistImgsModels []biz.GameExistImgs
    err := r.data.DataBase.Where("game_id = ?", id).Find(&gameExistImgsModels).Error
    if err != nil {
        return nil, err
    }
    imgIdArray := make([]string, 0, len(gameExistImgsModels))
    for _, imgModel := range gameExistImgsModels {
        imgIdArray = append(imgIdArray, imgModel.ImgID)
    }

    // TODO img uuid => img url ❌

    return imgIdArray, nil
}
