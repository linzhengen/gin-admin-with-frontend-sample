package model

import (
	"context"

	"github.com/linzhengen/gin-admin-with-frontend-sample/internal/app/schema"
)

// IBuyerOrder 仕入詳細存储接口
type IBuyerOrder interface {
	// 查询数据
	Query(ctx context.Context, params schema.BuyerOrderQueryParam, opts ...schema.BuyerOrderQueryOptions) (*schema.BuyerOrderQueryResult, error)
	// 查询指定数据
	Get(ctx context.Context, recordID string, opts ...schema.BuyerOrderQueryOptions) (*schema.BuyerOrder, error)
	// 创建数据
	Create(ctx context.Context, item schema.BuyerOrder) error
	// 更新数据
	Update(ctx context.Context, recordID string, item schema.BuyerOrder) error
	// 删除数据
	Delete(ctx context.Context, recordID string) error
}
