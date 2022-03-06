package housebiz

import (
	"context"
	"errors"
  "hfs_backend/modules/house/housemodel")

type DeleteHouseStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*housemodel.House, error)
	SoftDeleteData(ctx context.Context,
		id int,
	) error
}
type deleteHouseBiz struct {
	store DeleteHouseStore
}

func NewDeleteHouseBiz(store DeleteHouseStore) *deleteHouseBiz {
	return &deleteHouseBiz{store: store}
}
func (biz *deleteHouseBiz) DeleteHouse(ctx context.Context,
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
