package productbiz

import (
	"context"
	"errors"
	"hfs_backend/modules/product/productmodel"
)

type DeleteProductStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*productmodel.Product, error)
	SoftDeleteData(ctx context.Context,
		id int,
	) error
}
type deleteProductBiz struct {
	store DeleteProductStore
}

func NewDeleteProductBiz(store DeleteProductStore) *deleteProductBiz {
	return &deleteProductBiz{store: store}
}
func (biz *deleteProductBiz) DeleteProduct(ctx context.Context,
	id int,
) error {
	oldData, error := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if error != nil {
		return error
	}
	if oldData.Status == 0 {
		return errors.New("data deleted")
	}
	if err := biz.store.SoftDeleteData(ctx, id); err != nil {
		return err
	}
	return nil
}
