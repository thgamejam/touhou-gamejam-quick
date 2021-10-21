package biz

import (
	"context"
	"game/internal/util/uuid"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	Name        string `gorm:"not null; type:varchar(32)"`
	Description string `gorm:"type:varchar(64)"`
	ImgID  uuid.BinaryUUID `gorm:"not null; type:binary(16)"`
}

type UserRepo interface {
	GetUser(ctx context.Context, id uint) (*User, error)
	GetImg(ctx context.Context, id uint) (string, error)
	CreatUser(ctx context.Context,name string,des string)(uint,error)
}

type UserUseCase struct {
	UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{UserRepo: repo, log: log.NewHelper(logger)}
}
