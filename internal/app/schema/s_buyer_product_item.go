package schema

// BuyerProductItem 仕入商品詳細
type BuyerProductItem struct {
	RecordID string `json:"record_id" swaggo:"false,记录ID"`
	Creator  string `json:"creator" swaggo:"false,创建者"`
}

// BuyerProductItemQueryParam 查询条件
type BuyerProductItemQueryParam struct {
}

// BuyerProductItemQueryOptions 查询可选参数项
type BuyerProductItemQueryOptions struct {
	PageParam *PaginationParam // 分页参数
}

// BuyerProductItemQueryResult 查询结果
type BuyerProductItemQueryResult struct {
	Data       BuyerProductItems
	PageResult *PaginationResult
}

// BuyerProductItems 仕入商品詳細列表
type BuyerProductItems []*BuyerProductItem
