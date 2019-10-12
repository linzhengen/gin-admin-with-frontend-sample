package schema

import "time"

// BuyerOrder 仕入詳細
type BuyerOrder struct {
	RecordID             string    `json:"record_id" swaggo:"false,记录ID"`
	Creator              string    `json:"creator" swaggo:"false,创建者"`
	OrderID              string    `json:"order_id" binding:"required" swaggo:"true,订单号"`
	BuyerID              string    `json:"buyer_id" binding:"required" swaggo:"true,买家主账号id"`
	SellerID             string    `json:"seller_id" binding:"required" swaggo:"true,卖家主账号id"`
	SellerLoginId        string    `json:"seller_login_id" binding:"required" swaggo:"true,卖家LoginId，旺旺Id"`
	Status               string    `json:"status" binding:"required" swaggo:"true,交易状态"`
	SumProductPayment    float64   `json:"sum_product_payment" binding:"required" swaggo:"true,产品总金额"`
	TotalAmount          float64   `json:"total_amount" binding:"required" swaggo:"true,应付款总金额"`
	ShippingFee          float64   `json:"total_amount" binding:"required" swaggo:"true,运费"`
	EntryDiscount        float64   `json:"entry_discount" binding:"required" swaggo:"true,订单明细涨价或降价的金额"`
	CreateTime           time.Time `json:"create_time" binding:"required" swaggo:"true,创建时间"`
	ModifyTime           time.Time `json:"modify_time" swaggo:"true,修改时间"`
	PayTime              time.Time `json:"pay_time" swaggo:"true,首次付款时间"`
	BuyerAddress         string    `json:"buyer_address" binding:"required" swaggo:"true,买家收货地址"`
	LogisticsCompanyName string    `json:"logistics_company_name" swaggo:"true,物流公司名称"`
	LogisticsBillNo      string    `json:"logistics_code" swaggo:"true,物流公司运单号"`
}

// BuyerOrderQueryParam 查询条件
type BuyerOrderQueryParam struct {
}

// BuyerOrderQueryOptions 查询可选参数项
type BuyerOrderQueryOptions struct {
	PageParam *PaginationParam // 分页参数
}

// BuyerOrderQueryResult 查询结果
type BuyerOrderQueryResult struct {
	Data       BuyerOrders
	PageResult *PaginationResult
}

// BuyerOrders 仕入詳細列表
type BuyerOrders []*BuyerOrder
