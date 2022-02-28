package productbiz

import (
	"context"
	"errors"
	"hfs_backend/modules/product/productmodel"
)

type UpdateProductStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*productmodel.Product, error)
	UpdateData(
		ctx context.Context,
		id int,
		data *productmodel.ProductUpdate,
	) error
}
type ProductBiz struct {
	store UpdateProductStore
}

func NewUpdateProductBiz(store UpdateProductStore) *ProductBiz {
	return &ProductBiz{store: store}
}
func (biz *ProductBiz) UpdateProduct(ctx context.Context, id int, data *productmodel.ProductUpdate) error {
	oldData, error := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if error != nil {
		return error
	}
	if oldData.Status == 0 {
		return errors.New("data deleted")
	}
	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return err
	}
	return nil
}
