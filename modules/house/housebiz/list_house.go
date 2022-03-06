package housebiz

import (
  "context"
  "hfs_backend/common"
  "hfs_backend/modules/house/housemodel"
)

type ListProductStore interface {
	ListDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		filter *housemodel.HouseFilter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]housemodel.House, error)
}
type listProductBiz struct {
	store ListProductStore
}

func NewListHouseBiz(store ListProductStore) *listProductBiz {
	return &listProductBiz{store: store}
}
func (biz *listProductBiz) ListHouse(ctx context.Context,
  filter *housemodel.HouseFilter,
  paging *common.Paging) ([]housemodel.House, error) {
	result, err := biz.store.ListDataByCondition(ctx, nil, filter, paging)
	return result, err
}
