package internal

import (
	"context"

	"github.com/linzhengen/gin-admin-with-frontend-sample/internal/app/errors"
	"github.com/linzhengen/gin-admin-with-frontend-sample/internal/app/model"
	"github.com/linzhengen/gin-admin-with-frontend-sample/internal/app/schema"
	"github.com/linzhengen/gin-admin-with-frontend-sample/pkg/util"
)

// NewBuyerOrder 创建仕入詳細
func NewBuyerOrder(mBuyerOrder model.IBuyerOrder) *BuyerOrder {
	return &BuyerOrder{
		BuyerOrderModel: mBuyerOrder,
	}
}

// BuyerOrder 仕入詳細业务逻辑
type BuyerOrder struct {
	BuyerOrderModel model.IBuyerOrder
}

// Query 查询数据
func (a *BuyerOrder) Query(ctx context.Context, params schema.BuyerOrderQueryParam, opts ...schema.BuyerOrderQueryOptions) (*schema.BuyerOrderQueryResult, error) {
	return a.BuyerOrderModel.Query(ctx, params, opts...)
}

// Get 查询指定数据
func (a *BuyerOrder) Get(ctx context.Context, recordID string, opts ...schema.BuyerOrderQueryOptions) (*schema.BuyerOrder, error) {
	item, err := a.BuyerOrderModel.Get(ctx, recordID, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *BuyerOrder) getUpdate(ctx context.Context, recordID string) (*schema.BuyerOrder, error) {
	return a.Get(ctx, recordID)
}

// Create 创建数据
func (a *BuyerOrder) Create(ctx context.Context, item schema.BuyerOrder) (*schema.BuyerOrder, error) {
	item.RecordID = util.MustUUID()
	err := a.BuyerOrderModel.Create(ctx, item)
	if err != nil {
		return nil, err
	}
	return a.getUpdate(ctx, item.RecordID)
}

// Update 更新数据
func (a *BuyerOrder) Update(ctx context.Context, recordID string, item schema.BuyerOrder) (*schema.BuyerOrder, error) {
	oldItem, err := a.BuyerOrderModel.Get(ctx, recordID)
	if err != nil {
		return nil, err
	} else if oldItem == nil {
		return nil, errors.ErrNotFound
	}

	err = a.BuyerOrderModel.Update(ctx, recordID, item)
	if err != nil {
		return nil, err
	}
	return a.getUpdate(ctx, recordID)
}

// Delete 删除数据
func (a *BuyerOrder) Delete(ctx context.Context, recordID string) error {
	oldItem, err := a.BuyerOrderModel.Get(ctx, recordID)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.BuyerOrderModel.Delete(ctx, recordID)
}
