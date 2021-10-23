package main

import (
    "context"
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/credentials"
    "github.com/aws/aws-sdk-go-v2/service/s3"
    "time"
)

type ObjectStorage struct {
    Client *s3.Client // 对象存储服务

    bucket                  *string
    smallFileExpirationTime time.Duration //小文件到期时间
    largeFileExpirationTime time.Duration //大文件到期时间
}

// NewObjectStorage 初始化对象存储
func NewObjectStorage() *ObjectStorage {

    // TODO 绑定 Log 未完成

    const defaultRegion = "us-east-1"

    staticResolver := aws.EndpointResolverFunc(
        func(service, region string) (aws.Endpoint, error) {
            return aws.Endpoint{
                PartitionID:       "aws",
                URL:               "http://test.io:9000", // minio url
                SigningRegion:     defaultRegion,
                HostnameImmutable: true,
            }, nil
        },
    )

    cfg, err := config.LoadDefaultConfig(
        context.TODO(),
        func(lo *config.LoadOptions) error {
            lo.Region = defaultRegion
            lo.Credentials = credentials.NewStaticCredentialsProvider(
                "user", // Access Key
                "206242227", // Secret Key
                "",
            )
            lo.EndpointResolver = staticResolver
            return nil
        },
    )
    if err != nil {
        return nil
    }

    client := s3.NewFromConfig(cfg)

    oss := &ObjectStorage{
        Client:                  client,
        smallFileExpirationTime: time.Second * 300,
        largeFileExpirationTime: time.Second * 900,
    }

    oss.bucket = aws.String("test")

    return oss
}

// GetObjectURL 获取文件对象的预签名URL
// key: 对象路径
// expirationTime: 过期时间
func (o *ObjectStorage) GetObjectURL(ctx context.Context, key string, expirationTime time.Duration) (string, error) {
    input := &s3.GetObjectInput{
        Bucket: o.bucket,
        Key:    aws.String(key),
    }

    psClient := s3.NewPresignClient(o.Client)
    req, err := psClient.PresignGetObject(ctx, input,
        func(options *s3.PresignOptions) {
            options.Expires = expirationTime // 设置url过期时间
        },
    )
    if err != nil {
        return "", err
    }

    return req.URL, nil
}

func (o *ObjectStorage) GetSmallObjectURL(ctx context.Context, key string) (string, error) {
    return o.GetObjectURL(ctx, key, o.smallFileExpirationTime)
}

func (o *ObjectStorage) GetLargeObjectURL(ctx context.Context, key string) (string, error) {
    return o.GetObjectURL(ctx, key, o.largeFileExpirationTime)
}
