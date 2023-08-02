package biz

import (
	"context"
	v1 "kratos-layout/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter Greeter模型
type Greeter struct {
	Hello string
}

// GreeterRepo Greater数据存储定义
type GreeterRepo interface {
	// Save 保存
	Save(context.Context, *Greeter) (*Greeter, error)
	// Update 修改
	Update(context.Context, *Greeter) (*Greeter, error)
	// FindByID 通过主键id查找
	FindByID(context.Context, int64) (*Greeter, error)
	// ListByHello 条件查找
	ListByHello(context.Context, string) ([]*Greeter, error)
	// ListAll 查询所有
	ListAll(context.Context) ([]*Greeter, error)
}

// GreeterUsecase Greeter业务逻辑
type GreeterUsecase struct {
	repo GreeterRepo
	log  *log.Helper
}

// NewGreeterUsecase 创建Greeter业务逻辑处理
func NewGreeterUsecase(repo GreeterRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter 创建Greeter，并且返回创建的Greeter
func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
