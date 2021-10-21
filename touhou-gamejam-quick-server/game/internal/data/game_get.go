package data

import (
    "context"
    "game/internal/biz"
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/service/s3"
)

const imgKeyHead = "game_imgs/"

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
    tagIdArray := make([]uint, 0, len(gameExistTagsModels))
    for _, gameModel := range gameExistTagsModels {
        tagIdArray = append(tagIdArray, gameModel.TagsID)
    }

    // 数据库获取标签的对应名字
    var tagsModels []biz.GameTags
    err = r.data.DataBase.Find(&tagsModels, tagIdArray).Error
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
        imgIdArray = append(imgIdArray, imgModel.ImgID.String())
    }

    // 获取url
    psClient := s3.NewPresignClient(r.data.ObjectStorage.Client)

    imgURLArray := make([]string, 0, len(imgIdArray))
    for _, imgId := range imgIdArray {
        getInput := &s3.GetObjectInput{
            Bucket:                     r.data.ObjectStorage.bucket,
            Key:                        aws.String(imgKeyHead + imgId), // TODO 这里用了常量 imgKeyHead ，需要修改为配置文件
            //ResponseContentDisposition: aws.String("attachment;filename=" + imgId + ".jpg"),
        }
        req, err := psClient.PresignGetObject(context.TODO(), getInput,
            func(options *s3.PresignOptions) {
                options.Expires = r.data.ObjectStorage.smallFileExpirationTime
            },
        )
        if err != nil {
            continue
        }
        imgURLArray = append(imgURLArray, req.URL)
    }

    return imgURLArray, nil
}

func (r *gameRepo) GameImg(ctx context.Context, id uint) (string, error) {
    // 获取第一张图片作为预览图
    gameExistImgModel := &biz.GameExistImgs{}
    err := r.data.DataBase.Where("game_id = ?", id).First(&gameExistImgModel).Error
    if err != nil {
        return "", err
    }

    // 获取url
    psClient := s3.NewPresignClient(r.data.ObjectStorage.Client)
    getInput := &s3.GetObjectInput{
        Bucket: r.data.ObjectStorage.bucket,
        Key:    aws.String(imgKeyHead + gameExistImgModel.ImgID.String()), // TODO 这里用了常量 imgKeyHead ，需要修改为配置文件
    }
    req, err := psClient.PresignGetObject(context.TODO(), getInput,
        func(options *s3.PresignOptions) {
            options.Expires = r.data.ObjectStorage.smallFileExpirationTime
        },
    )
    if err != nil {
        return "", err
    }

    return req.URL, nil
}

func (r *gameRepo) GetGames(ctx context.Context) ([]biz.Game, error) {
    var models []biz.Game

    err := r.data.DataBase.Find(&models).Error
    if err != nil {
        return nil, err
    }

    return models, nil
}
