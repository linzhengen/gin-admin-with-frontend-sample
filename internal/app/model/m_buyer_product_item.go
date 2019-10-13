package model

import (
	"context"

	"github.com/linzhengen/gin-admin-with-frontend-sample/internal/app/schema"
)

// IBuyerProductItem 仕入商品詳細存储接口
type IBuyerProductItem interface {
	// 查询数据
	Query(ctx context.Context, params schema.BuyerProductItemQueryParam, opts ...schema.BuyerProductItemQueryOptions) (*schema.BuyerProductItemQueryResult, error)
	// 查询指定数据
	Get(ctx context.Context, recordID string, opts ...schema.BuyerProductItemQueryOptions) (*schema.BuyerProductItem, error)
	// 创建数据
	Create(ctx context.Context, item schema.BuyerProductItem) error
	// 更新数据
	Update(ctx context.Context, recordID string, item schema.BuyerProductItem) error
	// 删除数据
	Delete(ctx context.Context, recordID string) error
}
