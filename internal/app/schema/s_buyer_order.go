package schema

// BuyerOrder 仕入詳細
type BuyerOrder struct {
	RecordID string `json:"record_id" swaggo:"false,记录ID"`
	Creator  string `json:"creator" swaggo:"false,创建者"`
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
