package entity

import (
	"context"
	"time"

	"github.com/linzhengen/gin-admin-with-frontend-sample/internal/app/schema"
	"github.com/linzhengen/gin-admin-with-frontend-sample/pkg/gormplus"
)

// GetBuyerOrderDB 仕入詳細
func GetBuyerOrderDB(ctx context.Context, defDB *gormplus.DB) *gormplus.DB {
	return getDBWithModel(ctx, defDB, BuyerOrder{})
}

// SchemaBuyerOrder 仕入詳細
type SchemaBuyerOrder schema.BuyerOrder

// ToBuyerOrder 转换为仕入詳細实体
func (a SchemaBuyerOrder) ToBuyerOrder() *BuyerOrder {
	item := &BuyerOrder{
		RecordID: &a.RecordID,
		Creator:  &a.Creator,
	}
	return item
}

// BuyerOrder 仕入詳細实体
type BuyerOrder struct {
	Model
	RecordID             *string    `gorm:"column:record_id;size:36;index;"`                // 记录ID
	Creator              *string    `gorm:"column:creator;size:36;index;"`                  // 创建者
	OrderID              *string    `gorm:"column:order_id;size:38;index;"`                 // 订单号
	BuyerID              *string    `gorm:"column:buyer_id;size:128;"`                      // 买家主账号id
	SellerID             *string    `gorm:"column:seller_id;size:128;"`                     // 卖家主账号id
	SellerLoginId        *string    `gorm:"column:seller_login_id;size:128;"`               // 卖家LoginId，旺旺Id
	Status               *string    `gorm:"column:status;size:32;"`                         // 交易状态
	SumProductPayment    *float64   `gorm:"column:sum_product_payment;type:decimal(10,2);"` // 产品总金额
	TotalAmount          *float64   `gorm:"column:total_amount;type:decimal(10,2);"`        // 应付款总金额
	ShippingFee          *float64   `gorm:"column:shipping_fee;type:decimal(10,2);"`        // 运费
	EntryDiscount        *float64   `gorm:"column:entry_discount;type:decimal(10,2);"`      // 订单明细涨价或降价的金额
	CreateTime           *time.Time `gorm:"column:create_time;"`                            // 创建时间
	ModifyTime           *time.Time `gorm:"column:modify_time;"`                            // 修改时间
	PayTime              *time.Time `gorm:"column:pay_time;"`                               // 首次付款时间
	BuyerAddress         *string    `gorm:"column:buyer_address;size:256;"`                 // 买家收货地址
	LogisticsCompanyName *string    `gorm:"column:logistics_company_name;size:256;"`        // 物流公司名称
	LogisticsBillNo      *string    `gorm:"column:logistics_code;size:36;"`                 // 物流公司运单号
}

func (a BuyerOrder) String() string {
	return toString(a)
}

// TableName 表名
func (a BuyerOrder) TableName() string {
	return a.Model.TableName("buyer_orders")
}

// ToSchemaBuyerOrder 转换为仕入詳細对象
func (a BuyerOrder) ToSchemaBuyerOrder() *schema.BuyerOrder {
	item := &schema.BuyerOrder{
		RecordID: *a.RecordID,
		Creator:  *a.Creator,
	}
	return item
}

// BuyerOrders 仕入詳細列表
type BuyerOrders []*BuyerOrder

// ToSchemaBuyerOrders 转换为仕入詳細对象列表
func (a BuyerOrders) ToSchemaBuyerOrders() []*schema.BuyerOrder {
	list := make([]*schema.BuyerOrder, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaBuyerOrder()
	}
	return list
}

// BuyerOrderProductItem 订单货品关联实体
type BuyerOrderProductItem struct {
	Model
	OrderID       string `gorm:"column:order_id;size:38;index;"`        // 订单号
	ProductItemID string `gorm:"column:product_item_id;size:36;index;"` // 货物号
}

// TableName 表名
func (a BuyerOrderProductItem) TableName() string {
	return a.Model.TableName("buyer_order_product_items")
}
