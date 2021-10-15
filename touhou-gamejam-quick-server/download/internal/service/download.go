package service

import (
    "context"
    "download/internal/biz"
    "github.com/go-kratos/kratos/v2/log"

    pb "download/api/download/v1"
)

type DownloadService struct {
    pb.UnimplementedDownloadServer

    uc  *biz.DownloadUseCase
    log *log.Helper
}

func NewDownloadService(uc *biz.DownloadUseCase, logger log.Logger) *DownloadService {
    return &DownloadService{
        uc:  uc,
        log: log.NewHelper(logger),
    }
}

func (s *DownloadService) GetDownloadURL(ctx context.Context, req *pb.GetDownloadRequest) (*pb.GetDownloadReply, error) {

    return &pb.GetDownloadReply{}, nil
}
