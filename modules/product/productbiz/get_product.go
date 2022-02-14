package productbiz

import (
	"context"
	"hfs_backend/modules/product/productmodel"
)

type GetProductStore interface {
	FindDataByCondition(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*productmodel.Product, error)
}
type getProductBiz struct {
	store GetProductStore
}

func NewGetProductBiz(store GetProductStore) *getProductBiz {
	return &getProductBiz{store: store}
}
func (biz *getProductBiz) GetProduct(ctx context.Context, id int) (*productmodel.Product, error) {
	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	return data, err
}
