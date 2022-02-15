package productbiz

import (
	"context"
	"hfs_backend/common"
	"hfs_backend/modules/product/productmodel"
)

type GetProductStore interface {
	FindDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*productmodel.Product, error)
}
type getProductBiz struct {
	store GetProductStore
}

func NewGetProductBiz(store GetProductStore) *getProductBiz {
	return &getProductBiz{store: store}
}
func (biz *getProductBiz) GetProduct(ctx context.Context, id int) (*productmodel.Product, error) {
	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(productmodel.EntityName, err)
		}
		return nil, common.ErrCannotGetEntity(productmodel.EntityName, err)
	}
	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(productmodel.EntityName, nil)
	}
	return data, err
}
