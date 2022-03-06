package housebiz

import (
	"context"
  "hfs_backend/modules/house/housemodel"
)

type CreateHouseStore interface {
	Create(ctx context.Context, data *housemodel.HouseCreate) error
}
type createHouseBiz struct {
	store CreateHouseStore
}

func NewCreateHouseBiz(store CreateHouseStore) *createHouseBiz {
	return &createHouseBiz{store: store}
}

func (biz *createHouseBiz) CreateHouse(ctx context.Context, data *housemodel.HouseCreate) error {

	//if err := data.Validate(); err != nil {
	//	return err
	//}
	err := biz.store.Create(ctx, data)
	return err
}
