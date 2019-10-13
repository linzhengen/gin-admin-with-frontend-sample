package ctl

import (
	"github.com/gin-gonic/gin"
	"github.com/linzhengen/gin-admin-with-frontend-sample/internal/app/bll"
	"github.com/linzhengen/gin-admin-with-frontend-sample/internal/app/errors"
	"github.com/linzhengen/gin-admin-with-frontend-sample/internal/app/ginplus"
	"github.com/linzhengen/gin-admin-with-frontend-sample/internal/app/schema"
)

// NewBuyerProductItem 创建仕入商品詳細控制器
func NewBuyerProductItem(bBuyerProductItem bll.IBuyerProductItem) *BuyerProductItem {
	return &BuyerProductItem{
		BuyerProductItemBll: bBuyerProductItem,
	}
}

// BuyerProductItem 仕入商品詳細
// @Name BuyerProductItem
// @Description 仕入商品詳細控制器
type BuyerProductItem struct {
	BuyerProductItemBll bll.IBuyerProductItem
}

// Query 查询数据
func (a *BuyerProductItem) Query(c *gin.Context) {
	switch c.Query("q") {
	case "page":
		a.QueryPage(c)
	default:
		ginplus.ResError(c, errors.ErrUnknownQuery)
	}
}

// QueryPage 查询分页数据
// @Summary 查询分页数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param current query int true "分页索引" 1
// @Param pageSize query int true "分页大小" 10
// @Success 200 []schema.BuyerProductItem "查询结果：{list:列表数据,pagination:{current:页索引,pageSize:页大小,total:总数量}}"
// @Failure 400 schema.HTTPError "{error:{code:0,message:未知的查询类型}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router GET /api/v1/buyer_product_items?q=page
func (a *BuyerProductItem) QueryPage(c *gin.Context) {
	var params schema.BuyerProductItemQueryParam

	result, err := a.BuyerProductItemBll.Query(ginplus.NewContext(c), params, schema.BuyerProductItemQueryOptions{
		PageParam: ginplus.GetPaginationParam(c),
	})
	if err != nil {
		ginplus.ResError(c, err)
		return
	}

	ginplus.ResPage(c, result.Data, result.PageResult)
}

// Get 查询指定数据
// @Summary 查询指定数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "记录ID"
// @Success 200 schema.BuyerProductItem
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 404 schema.HTTPError "{error:{code:0,message:资源不存在}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router GET /api/v1/buyer_product_items/{id}
func (a *BuyerProductItem) Get(c *gin.Context) {
	item, err := a.BuyerProductItemBll.Get(ginplus.NewContext(c), c.Param("id"))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, item)
}

// Create 创建数据
// @Summary 创建数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param body body schema.BuyerProductItem true
// @Success 200 schema.BuyerProductItem
// @Failure 400 schema.HTTPError "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router POST /api/v1/buyer_product_items
func (a *BuyerProductItem) Create(c *gin.Context) {
	var item schema.BuyerProductItem
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	item.Creator = ginplus.GetUserID(c)
	nitem, err := a.BuyerProductItemBll.Create(ginplus.NewContext(c), item)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, nitem)
}

// Update 更新数据
// @Summary 更新数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "记录ID"
// @Param body body schema.BuyerProductItem true
// @Success 200 schema.BuyerProductItem
// @Failure 400 schema.HTTPError "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router PUT /api/v1/buyer_product_items/{id}
func (a *BuyerProductItem) Update(c *gin.Context) {
	var item schema.BuyerProductItem
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	nitem, err := a.BuyerProductItemBll.Update(ginplus.NewContext(c), c.Param("id"), item)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, nitem)
}

// Delete 删除数据
// @Summary 删除数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "记录ID"
// @Success 200 schema.HTTPStatus "{status:OK}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router DELETE /api/v1/buyer_product_items/{id}
func (a *BuyerProductItem) Delete(c *gin.Context) {
	err := a.BuyerProductItemBll.Delete(ginplus.NewContext(c), c.Param("id"))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}
