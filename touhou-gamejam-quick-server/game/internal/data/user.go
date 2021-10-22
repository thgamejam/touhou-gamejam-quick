package data

import (
	"context"
	"game/internal/biz"
	"game/internal/util/uuid"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

const userKeyHead = "user_img/"

func NewUseRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u *userRepo) GetUser(ctx context.Context, id uint) (*biz.User, error) {
	model := &biz.User{}
	err := u.data.DataBase.First(&model, id).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (u *userRepo) GetImg(ctx context.Context, id uint) (string, error) {
	model, err := u.GetUser(ctx, id)
	if err != nil {
		return "", nil
	}
	psClient := s3.NewPresignClient(u.data.ObjectStorage.Client)
	getInput := &s3.GetObjectInput{
		Bucket: u.data.ObjectStorage.bucket,
		Key:    aws.String(userKeyHead + model.ImgID.String()), // TODO
	}
	req, err := psClient.PresignGetObject(ctx, getInput,
		func(options *s3.PresignOptions) {
			options.Expires = u.data.ObjectStorage.smallFileExpirationTime
		},
	)
	if err != nil {
		return "", err
	}
	return req.URL, nil
}

func (u *userRepo) CreatUser(ctx context.Context, name string, des string) (uint, error) {
	model := &biz.User{
		Name:        name,
		Description: des,
		ImgID:       uuid.New(),
	}
	err := u.data.DataBase.Create(&model).Error
	if err != nil {
		return 0, err
	}
	return model.ID, nil
}
