package housestorage

import (
  "context"
  "hfs_backend/common"
  "hfs_backend/modules/house/housemodel"
)

func (s *sqlStore) Create(ctx context.Context, data *housemodel.HouseCreate) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
