package model

import (
	"context"

	"github.com/linzhengen/gin-admin-with-frontend-sample/internal/app/errors"
	"github.com/linzhengen/gin-admin-with-frontend-sample/internal/app/model/impl/gorm/internal/entity"
	"github.com/linzhengen/gin-admin-with-frontend-sample/internal/app/schema"
	"github.com/linzhengen/gin-admin-with-frontend-sample/pkg/gormplus"
)

// NewBuyerOrder 创建仕入詳細存储实例
func NewBuyerOrder(db *gormplus.DB) *BuyerOrder {
	return &BuyerOrder{db}
}

// BuyerOrder 仕入詳細存储
type BuyerOrder struct {
	db *gormplus.DB
}

func (a *BuyerOrder) getQueryOption(opts ...schema.BuyerOrderQueryOptions) schema.BuyerOrderQueryOptions {
	var opt schema.BuyerOrderQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *BuyerOrder) Query(ctx context.Context, params schema.BuyerOrderQueryParam, opts ...schema.BuyerOrderQueryOptions) (*schema.BuyerOrderQueryResult, error) {
	db := entity.GetBuyerOrderDB(ctx, a.db).DB

	db = db.Order("id DESC")

	opt := a.getQueryOption(opts...)
	var list entity.BuyerOrders
	pr, err := WrapPageQuery(db, opt.PageParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	qr := &schema.BuyerOrderQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaBuyerOrders(),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *BuyerOrder) Get(ctx context.Context, recordID string, opts ...schema.BuyerOrderQueryOptions) (*schema.BuyerOrder, error) {
	db := entity.GetBuyerOrderDB(ctx, a.db).Where("record_id=?", recordID)
	var item entity.BuyerOrder
	ok, err := a.db.FindOne(db, &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaBuyerOrder(), nil
}

// Create 创建数据
func (a *BuyerOrder) Create(ctx context.Context, item schema.BuyerOrder) error {
	BuyerOrder := entity.SchemaBuyerOrder(item).ToBuyerOrder()
	result := entity.GetBuyerOrderDB(ctx, a.db).Create(BuyerOrder)
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Update 更新数据
func (a *BuyerOrder) Update(ctx context.Context, recordID string, item schema.BuyerOrder) error {
	BuyerOrder := entity.SchemaBuyerOrder(item).ToBuyerOrder()
	result := entity.GetBuyerOrderDB(ctx, a.db).Where("record_id=?", recordID).Omit("record_id", "creator").Updates(BuyerOrder)
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Delete 删除数据
func (a *BuyerOrder) Delete(ctx context.Context, recordID string) error {
	result := entity.GetBuyerOrderDB(ctx, a.db).Where("record_id=?", recordID).Delete(entity.BuyerOrder{})
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
