package productbiz

import (
	"context"
	"hfs_backend/common"
	"hfs_backend/modules/product/productmodel"
)

type ListProductStore interface {
	ListDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		filter *productmodel.ProductFilter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]productmodel.Product, error)
}
type listProductBiz struct {
	store ListProductStore
}

func NewListProductBiz(store ListProductStore) *listProductBiz {
	return &listProductBiz{store: store}
}
func (biz *listProductBiz) ListRestaurant(ctx context.Context, filter *productmodel.ProductFilter, paging *common.Paging) ([]productmodel.Product, error) {
	result, err := biz.store.ListDataByCondition(ctx, nil, filter, paging)
	return result, err
}
