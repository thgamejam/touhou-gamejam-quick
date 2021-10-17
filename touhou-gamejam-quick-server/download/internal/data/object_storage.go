package data

import (
    "context"
    "download/internal/conf"
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/credentials"
    "github.com/aws/aws-sdk-go-v2/service/s3"
)

type ObjectStorage struct {
    Client *s3.Client             // 对象存储服务
    Input  *s3.ListObjectsV2Input // 服务调用的输入参数，目前只传入 bucket

    bucket *string
}

// NewObjectStorage 初始化对象存储
func NewObjectStorage(c *conf.Data) (*ObjectStorage, error) {

    // TODO 绑定 Log 未完成

    const defaultRegion = "us-east-1"

    staticResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
        return aws.Endpoint{
            PartitionID:       "aws",
            URL:               c.ObjectStorage.Url, // minio url
            SigningRegion:     defaultRegion,
            HostnameImmutable: true,
        }, nil
    })

    cfg, err := config.LoadDefaultConfig(
        context.TODO(),
        func(lo *config.LoadOptions) error {
            lo.Region = defaultRegion
            lo.Credentials = credentials.NewStaticCredentialsProvider(
                c.ObjectStorage.AccessKey, // Access Key
                c.ObjectStorage.SecretKey, // Secret Key
                "",
            )
            lo.EndpointResolver = staticResolver
            return nil
        },
    )
    if err != nil {
        return nil, err
    }

    client := s3.NewFromConfig(cfg)

    // 选择 bucket
    input := &s3.ListObjectsV2Input{
        Bucket: &c.ObjectStorage.Bucket,
    }

    return &ObjectStorage{
        Client: client,
        Input:  input,
        bucket: &c.ObjectStorage.Bucket,
    }, nil
}
