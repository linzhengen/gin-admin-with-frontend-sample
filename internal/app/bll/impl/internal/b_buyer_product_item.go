package internal

import (
	"context"

	"github.com/linzhengen/gin-admin-with-frontend-sample/internal/app/errors"
	"github.com/linzhengen/gin-admin-with-frontend-sample/internal/app/model"
	"github.com/linzhengen/gin-admin-with-frontend-sample/internal/app/schema"
	"github.com/linzhengen/gin-admin-with-frontend-sample/pkg/util"
)

// NewBuyerProductItem 创建仕入商品詳細
func NewBuyerProductItem(mBuyerProductItem model.IBuyerProductItem) *BuyerProductItem {
	return &BuyerProductItem{
		BuyerProductItemModel: mBuyerProductItem,
	}
}

// BuyerProductItem 仕入商品詳細业务逻辑
type BuyerProductItem struct {
	BuyerProductItemModel model.IBuyerProductItem
}

// Query 查询数据
func (a *BuyerProductItem) Query(ctx context.Context, params schema.BuyerProductItemQueryParam, opts ...schema.BuyerProductItemQueryOptions) (*schema.BuyerProductItemQueryResult, error) {
	return a.BuyerProductItemModel.Query(ctx, params, opts...)
}

// Get 查询指定数据
func (a *BuyerProductItem) Get(ctx context.Context, recordID string, opts ...schema.BuyerProductItemQueryOptions) (*schema.BuyerProductItem, error) {
	item, err := a.BuyerProductItemModel.Get(ctx, recordID, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *BuyerProductItem) getUpdate(ctx context.Context, recordID string) (*schema.BuyerProductItem, error) {
	return a.Get(ctx, recordID)
}

// Create 创建数据
func (a *BuyerProductItem) Create(ctx context.Context, item schema.BuyerProductItem) (*schema.BuyerProductItem, error) {
	item.RecordID = util.MustUUID()
	err := a.BuyerProductItemModel.Create(ctx, item)
	if err != nil {
		return nil, err
	}
	return a.getUpdate(ctx, item.RecordID)
}

// Update 更新数据
func (a *BuyerProductItem) Update(ctx context.Context, recordID string, item schema.BuyerProductItem) (*schema.BuyerProductItem, error) {
	oldItem, err := a.BuyerProductItemModel.Get(ctx, recordID)
	if err != nil {
		return nil, err
	} else if oldItem == nil {
		return nil, errors.ErrNotFound
	}

	err = a.BuyerProductItemModel.Update(ctx, recordID, item)
	if err != nil {
		return nil, err
	}
	return a.getUpdate(ctx, recordID)
}

// Delete 删除数据
func (a *BuyerProductItem) Delete(ctx context.Context, recordID string) error {
	oldItem, err := a.BuyerProductItemModel.Get(ctx, recordID)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.BuyerProductItemModel.Delete(ctx, recordID)
}
