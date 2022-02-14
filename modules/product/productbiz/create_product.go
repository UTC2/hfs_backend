package productbiz

import (
	"context"
	"hfs_backend/modules/product/productmodel"
)

type CreateProductStore interface {
	Create(ctx context.Context, data *productmodel.ProductCreate) error
}
type createProductBiz struct {
	store CreateProductStore
}

func NewCreateProductBiz(store CreateProductStore) *createProductBiz {
	return &createProductBiz{store: store}
}
func (biz *createProductBiz) CreateProduct(ctx context.Context, data *productmodel.ProductCreate) error {

	if err := data.Validate(); err != nil {
		return err
	}
	err := biz.store.Create(ctx, data)
	return err
}
