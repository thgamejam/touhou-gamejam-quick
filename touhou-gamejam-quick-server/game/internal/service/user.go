package service

import (
	"context"
	"game/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	
	pb "game/api/user/v1"
)

type UserService struct {
	pb.UnimplementedUserServer
	
	uc  *biz.UserUseCase
	log *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	body := req.GetBody()
	id, err := s.uc.CreatUser(
		ctx,
		body.GetName(),
		body.GetDescription(),
	)
	if err != nil {
		return nil, pb.ErrorUserFound("")
	}
	return &pb.CreateUserReply{
		Id: uint32(id),
	}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	id := uint(req.GetId())
	model, err := s.uc.GetUser(ctx, id)
	if err != nil {
		return nil, pb.ErrorUserFound("User Not Fond: %v", id)
	}
	url, err := s.uc.GetImg(ctx, id)
	return &pb.GetUserReply{
		Name: model.Name,
		Img:  url,
	}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}
