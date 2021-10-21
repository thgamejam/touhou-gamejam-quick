package data

import (
    "context"
    "game/internal/util/uuid"
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/service/s3"
)

const zipKeyHead = "game/"

func (r *gameRepo) GetGameDownloadURL(ctx context.Context, uuid uuid.BinaryUUID, name string) (string, error) {
    zipName := uuid.String()

    // 获取url
    psClient := s3.NewPresignClient(r.data.ObjectStorage.Client)
    getInput := &s3.GetObjectInput{
        Bucket:                     r.data.ObjectStorage.bucket,
        Key:                        aws.String(zipKeyHead + zipName), // TODO 这里用了常量 zipKeyHead ，需要修改为配置文件
        ResponseContentType:        aws.String("application/octet-stream"),
        ResponseContentDisposition: aws.String("attachment;filename=" + name + ".zip"),
    }
    req, err := psClient.PresignGetObject(context.TODO(), getInput,
        func(options *s3.PresignOptions) {
            options.Expires = r.data.ObjectStorage.largeFileExpirationTime
        },
    )
    if err != nil {
        return "", err
    }

    return req.URL, err
}
