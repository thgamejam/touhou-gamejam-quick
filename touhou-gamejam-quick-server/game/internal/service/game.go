package service

import (
    "context"
    "game/internal/biz"
    "game/internal/util/uuid"
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
    body := req.GetBody()
    gameModel, err := s.uc.CreateGame(
        ctx,
        body.GetName(),
        uint(body.GetAuthorId()),
        body.GetDescription(),
        uuid.New(),
    )
    if err != nil {
        return nil, pb.ErrorGameNotFound("game creation failed: %v", body.GetName())
    }

    return &pb.CreateGameReply{
        Id: uint32(gameModel.ID),
        DownloadId: gameModel.DownloadID.String(),
    }, nil
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

    imgs, err := s.uc.GameImgs(ctx, id)

    return &pb.GetGameReply{
        Name:        gameModel.Name,
        AuthorId:    uint32(gameModel.AuthorID),
        Description: gameModel.Description,
        DownloadId:  gameModel.DownloadID.String(),
        Tags:        tags,
        Imgs:        imgs,
    }, nil
}
func (s *GameService) ListGame(ctx context.Context, req *pb.ListGameRequest) (*pb.ListGameReply, error) {
    return &pb.ListGameReply{}, nil
}
