package housestorage

import (
  "context"
  "hfs_backend/common"
  "hfs_backend/modules/house/housemodel"
)
func (s *sqlStore) SoftDeleteData(ctx context.Context,
	id int,
) error {
	db := s.db
	if error := db.Table(housemodel.House{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": 0,
		}).Error; error != nil {
		return common.ErrDB(error)
	}
	return nil
}
