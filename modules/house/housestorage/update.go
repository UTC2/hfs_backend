package housestorage

import (
  "context"
  "hfs_backend/common"
  "hfs_backend/modules/house/housemodel"
)

func (s *sqlStore) UpdateData(ctx context.Context, id int, data *housemodel.HouseUpdate) error {
	db := s.db
	if error := db.Where("id = ?", id).Updates(data).Error; error != nil {
		return common.ErrDB(error)
	}
	return nil
}
