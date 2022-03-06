package housebiz

import (
  "context"
  "errors"
  "hfs_backend/modules/house/housemodel"
)

type UpdateHouseStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*housemodel.House, error)
	UpdateData(
		ctx context.Context,
		id int,
		data *housemodel.HouseUpdate,
	) error
}
type ProductBiz struct {
	store UpdateHouseStore
}

func NewUpdateHouseBiz(store UpdateHouseStore) *ProductBiz {
	return &ProductBiz{store: store}
}
func (biz *ProductBiz) UpdateHouse(ctx context.Context, id int, data *housemodel.HouseUpdate) error {
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
