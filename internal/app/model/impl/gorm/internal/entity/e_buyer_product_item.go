package entity

import (
	"context"

	"github.com/linzhengen/gin-admin-with-frontend-sample/internal/app/schema"
	"github.com/linzhengen/gin-admin-with-frontend-sample/pkg/gormplus"
)

// GetBuyerProductItemDB 仕入商品詳細
func GetBuyerProductItemDB(ctx context.Context, defDB *gormplus.DB) *gormplus.DB {
	return getDBWithModel(ctx, defDB, BuyerProductItem{})
}

// SchemaBuyerProductItem 仕入商品詳細
type SchemaBuyerProductItem schema.BuyerProductItem

// ToBuyerProductItem 转换为仕入商品詳細实体
func (a SchemaBuyerProductItem) ToBuyerProductItem() *BuyerProductItem {
	item := &BuyerProductItem{
		RecordID: &a.RecordID,
		Creator:  &a.Creator,
	}
	return item
}

// BuyerProductItem 仕入商品詳細实体
type BuyerProductItem struct {
	Model
	RecordID    *string `gorm:"column:record_id;size:36;index;"` // 记录ID
	Creator     *string `gorm:"column:creator;size:36;index;"`   // 创建者
	ProductCode *string
	Name        *string
	Price       *int64
	Quantity    *int64
	Unit        *string
}

func (a BuyerProductItem) String() string {
	return toString(a)
}

// TableName 表名
func (a BuyerProductItem) TableName() string {
	return a.Model.TableName("buyer_product_items")
}

// ToSchemaBuyerProductItem 转换为仕入商品詳細对象
func (a BuyerProductItem) ToSchemaBuyerProductItem() *schema.BuyerProductItem {
	item := &schema.BuyerProductItem{
		RecordID: *a.RecordID,
		Creator:  *a.Creator,
	}
	return item
}

// BuyerProductItems 仕入商品詳細列表
type BuyerProductItems []*BuyerProductItem

// ToSchemaBuyerProductItems 转换为仕入商品詳細对象列表
func (a BuyerProductItems) ToSchemaBuyerProductItems() []*schema.BuyerProductItem {
	list := make([]*schema.BuyerProductItem, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaBuyerProductItem()
	}
	return list
}
