package housebiz

import (
	"context"
	"hfs_backend/common"
  "hfs_backend/modules/house/housemodel"
  "hfs_backend/modules/product/productmodel"
)

type GetHouseStore interface {
	FindDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*housemodel.House, error)
}
type getProductBiz struct {
	store GetHouseStore
}

func NewGetHouseBiz(store GetHouseStore) *getProductBiz {
	return &getProductBiz{store: store}
}
func (biz *getProductBiz) GetHouse(ctx context.Context, id int) (*housemodel.House, error) {
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
