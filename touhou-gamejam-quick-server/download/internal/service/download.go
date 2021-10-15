package service

import (
	"context"

	pb "download/api/download/v1"
)

type DownloadService struct {
	pb.UnimplementedDownloadServer
}

func NewDownloadService() *DownloadService {
	return &DownloadService{}
}

func (s *DownloadService) GetDownloadURL(ctx context.Context, req *pb.GetDownloadRequest) (*pb.GetDownloadReply, error) {


	return &pb.GetDownloadReply{

	}, nil
}
