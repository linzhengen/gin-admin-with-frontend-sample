package model

import (
	"context"

	"github.com/linzhengen/gin-admin-with-frontend-sample/internal/app/errors"
	"github.com/linzhengen/gin-admin-with-frontend-sample/internal/app/model/impl/gorm/internal/entity"
	"github.com/linzhengen/gin-admin-with-frontend-sample/internal/app/schema"
	"github.com/linzhengen/gin-admin-with-frontend-sample/pkg/gormplus"
)

// NewBuyerProductItem 创建仕入商品詳細存储实例
func NewBuyerProductItem(db *gormplus.DB) *BuyerProductItem {
	return &BuyerProductItem{db}
}

// BuyerProductItem 仕入商品詳細存储
type BuyerProductItem struct {
	db *gormplus.DB
}

func (a *BuyerProductItem) getQueryOption(opts ...schema.BuyerProductItemQueryOptions) schema.BuyerProductItemQueryOptions {
	var opt schema.BuyerProductItemQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *BuyerProductItem) Query(ctx context.Context, params schema.BuyerProductItemQueryParam, opts ...schema.BuyerProductItemQueryOptions) (*schema.BuyerProductItemQueryResult, error) {
	db := entity.GetBuyerProductItemDB(ctx, a.db).DB

	db = db.Order("id DESC")

	opt := a.getQueryOption(opts...)
	var list entity.BuyerProductItems
	pr, err := WrapPageQuery(db, opt.PageParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	qr := &schema.BuyerProductItemQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaBuyerProductItems(),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *BuyerProductItem) Get(ctx context.Context, recordID string, opts ...schema.BuyerProductItemQueryOptions) (*schema.BuyerProductItem, error) {
	db := entity.GetBuyerProductItemDB(ctx, a.db).Where("record_id=?", recordID)
	var item entity.BuyerProductItem
	ok, err := a.db.FindOne(db, &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaBuyerProductItem(), nil
}

// Create 创建数据
func (a *BuyerProductItem) Create(ctx context.Context, item schema.BuyerProductItem) error {
	BuyerProductItem := entity.SchemaBuyerProductItem(item).ToBuyerProductItem()
	result := entity.GetBuyerProductItemDB(ctx, a.db).Create(BuyerProductItem)
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Update 更新数据
func (a *BuyerProductItem) Update(ctx context.Context, recordID string, item schema.BuyerProductItem) error {
	BuyerProductItem := entity.SchemaBuyerProductItem(item).ToBuyerProductItem()
	result := entity.GetBuyerProductItemDB(ctx, a.db).Where("record_id=?", recordID).Omit("record_id", "creator").Updates(BuyerProductItem)
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Delete 删除数据
func (a *BuyerProductItem) Delete(ctx context.Context, recordID string) error {
	result := entity.GetBuyerProductItemDB(ctx, a.db).Where("record_id=?", recordID).Delete(entity.BuyerProductItem{})
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
