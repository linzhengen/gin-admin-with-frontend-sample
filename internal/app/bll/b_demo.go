package bll

import (
	"context"

	"github.com/LyricTian/gin-admin/internal/app/schema"
)

// IDemo demo业务逻辑接口
type IDemo interface {
	// 查询数据
	Query(ctx context.Context, params schema.DemoQueryParam, opts ...schema.DemoQueryOptions) (*schema.DemoQueryResult, error)
	// 查询指定数据
	Get(ctx context.Context, recordID string, opts ...schema.DemoQueryOptions) (*schema.Demo, error)
	// 创建数据
	Create(ctx context.Context, item schema.Demo) (*schema.Demo, error)
	// 更新数据
	Update(ctx context.Context, recordID string, item schema.Demo) (*schema.Demo, error)
	// 删除数据
	Delete(ctx context.Context, recordID string) error
	// 更新状态
	UpdateStatus(ctx context.Context, recordID string, status int) error
}
