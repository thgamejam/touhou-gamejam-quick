package service

import (
    "context"
    "game/internal/biz"
    "github.com/go-kratos/kratos/v2/log"

    pb "game/api/game/v1"
)

type GameService struct {
    pb.UnimplementedGameServer

    uc  *biz.GameUseCase
    log *log.Helper
}

func NewGameService(uc *biz.GameUseCase, logger log.Logger) *GameService {
    return &GameService{
        uc:  uc,
        log: log.NewHelper(logger),
    }
}

func (s *GameService) CreateGame(ctx context.Context, req *pb.CreateGameRequest) (*pb.CreateGameReply, error) {
    return &pb.CreateGameReply{}, nil
}
func (s *GameService) UpdateGame(ctx context.Context, req *pb.UpdateGameRequest) (*pb.UpdateGameReply, error) {
    return &pb.UpdateGameReply{}, nil
}
func (s *GameService) DeleteGame(ctx context.Context, req *pb.DeleteGameRequest) (*pb.DeleteGameReply, error) {
    return &pb.DeleteGameReply{}, nil
}
func (s *GameService) GetGame(ctx context.Context, req *pb.GetGameRequest) (*pb.GetGameReply, error) {
    id := uint(req.GetId())
    gameModel, err := s.uc.GetGame(ctx, id)
    if err != nil {
       //s.log.WithContext(ctx).Errorf("GetGame Error: %v. %v", id, err.Error())
       return nil, pb.ErrorGameNotFound("game not found: %v", id)
    }

    tags, err := s.uc.GameTags(ctx, id)
    if err != nil {
       //s.log.WithContext(ctx).Errorf("GetGame Error: %v. %v", id, err.Error())
       return nil, pb.ErrorGameNotFound("game tags not found: %v", id)
    }

    imgs, err := s.uc.GameImgs(ctx, id)
    if err != nil {
       //s.log.WithContext(ctx).Errorf("GetGame Error: %v. %v", id, err.Error())
       return nil, pb.ErrorGameNotFound("game imgs not found: %v", id)
    }

    return &pb.GetGameReply{
       Name: gameModel.Name,
       AuthorId: uint32(gameModel.AuthorID),
       Description: gameModel.Description,
       DownloadId: gameModel.DownloadID,
       Tags: tags,
       Imgs: imgs,
    }, nil
}
func (s *GameService) ListGame(ctx context.Context, req *pb.ListGameRequest) (*pb.ListGameReply, error) {
    return &pb.ListGameReply{}, nil
}
